package events

import (
	"database/sql"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"github.com/valyala/fasthttp"
	"github.com/google/uuid"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func invalidateLogin(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("cho-token", "no")
	ctx.Write(packets.LoginReply(-1))
}

func Login(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()

	lines := strings.Split(string(body), "\n")

	if len(lines) < 3 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	username := lines[0]
	password := lines[1]

	var (
		userID			int
		safeUsername	string
		dbPassword 		string
		privileges		constants.AkatsukiPrivileges
	)

	err := config.Database.QueryRow("SELECT id, username, username_safe, password_md5, privileges FROM users WHERE username = ? OR username_safe = ?;", username, username).Scan(&userID, &username, &safeUsername, &dbPassword, &privileges)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			invalidateLogin(ctx)
			return
		default:
			log.Error(err)
			ctx.SetStatusCode(http.StatusInternalServerError)
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)) != nil {
		invalidateLogin(ctx)
		return
	}

	if (privileges & constants.UserPendingVerification) != 0 {
		privileges &= ^constants.UserPendingVerification
		config.Database.Exec("UPDATE users SET privileges = ? WHERE id = ?", privileges, userID)
	}

	_ = lines[2]

	// check client data :d

	player := objects.NewPlayer(userID)
	player.Username = username
	player.SafeUsername = safeUsername
	player.Privileges = privileges
	player.IngamePrivileges = privileges.BanchoPrivileges()
	player.Token = uuid.Must(uuid.NewRandom()).String()

	ctx.Response.Header.Set("cho-token", player.Token)

	ctx.Write(packets.LoginReply(player.ID))
	ctx.Write(packets.ProtocolVersion(19))
	ctx.Write(packets.Alert("Welcome back!"))

	// Channels
	ctx.Write(packets.ChannelInfoEnd())

	objects.ChannelMutex.RLock()
	for _, c := range objects.Channels {
		if c.Autojoin && c.Join(player) {
			ctx.Write(packets.AutojoinChannel(c))
			ctx.Write(packets.JoinedChannel(c.Name))
		} else {
			ctx.Write(packets.AvailableChannel(c))
		}
	}
	objects.ChannelMutex.RUnlock()

	stats := packets.Stats(player)
	presence := packets.Presence(player)

	ctx.Write(presence)
	ctx.Write(stats)

	go func() {
		objects.PlayerMutex.RLock()
		for _, p := range objects.Players {
			p.Write(presence, stats)
			player.Write(packets.Presence(p), packets.Stats(p))
		}
		objects.PlayerMutex.RUnlock()

		objects.PlayerMutex.Lock()
		objects.Players[player.Token] = player
		objects.PlayerMutex.Unlock()

		var friends []int
		err := config.Database.Select(&friends, "SELECT user2 FROM users_relationships WHERE user1 = ?", player.ID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			return
		}

		player.Write(packets.FriendsList(friends))
	}()

	log.Infof("%s (%d) logged in.", player.Username, player.ID)
}