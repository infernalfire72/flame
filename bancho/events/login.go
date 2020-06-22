package events

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

const (
	_ = -iota
	InvalidLoginData
	ClientOutdated
	UserBanned
	_
	ServerError
	SupporterOnly
	PasswordReset
	VerifyIdentity
)

func invalidateLogin(ctx *fasthttp.RequestCtx, code int) {
	ctx.Response.Header.Set("cho-token", "no")
	ctx.Write(packets.LoginReply(code))
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
		userID       int
		safeUsername string
		dbPassword   string
		privileges   constants.AkatsukiPrivileges
	)

	err := config.Database.QueryRow("SELECT id, username, username_safe, password_md5, privileges FROM users WHERE username = ? OR username_safe = ?;", username, username).Scan(&userID, &username, &safeUsername, &dbPassword, &privileges)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			invalidateLogin(ctx, InvalidLoginData)
			return
		default:
			log.Error(err)
			ctx.SetStatusCode(http.StatusInternalServerError)
			invalidateLogin(ctx, ServerError)
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)) != nil {
		invalidateLogin(ctx, InvalidLoginData)
		return
	}

	if (privileges & constants.UserPendingVerification) != 0 {
		privileges &= ^constants.UserPendingVerification
		privileges |= constants.UserNormal | constants.UserPublic
		config.Database.Exec("UPDATE users SET privileges = ? WHERE id = ?", privileges, userID)
	}

	// Client Data Structure:
	// [0] Version
	// [1] Timezone
	// [2] Display City Location
	// [3] Hash Set
	// [4] Non-Friend DMs blocked
	clientData := strings.Split(lines[2], "|")
	if len(clientData) != 5 {

	}

	// Hash Set Structure:
	// [0] osu! exe md5
	// [1] MAC Adresses
	// [2] Hashed MAC Adresses
	// [3] Disk UUID
	// [4] Disk HWID
	hashSet := strings.Split(clientData[3], ":")
	if len(hashSet) != 6 {

	}

	var match int
	if hashSet[4] == "runningunderwine" {
		err = config.Database.Get(&match, "SELECT userid FROM hw_user WHERE userid <> ? AND activated AND unique_id = ?", userID, hashSet[3])
	} else {
		err = config.Database.Get(&match, "SELECT userid FROM hw_user WHERE userid <> ? AND activated AND disk_id = ? AND unique_id = ? AND mac = ?", userID, hashSet[4], hashSet[3], hashSet[2])
	}

	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		invalidateLogin(ctx, ServerError)
		return
	}

	if match != 0 {
		// Do stuff (restrict etc.)
		// Restrict(match)
		// Ban(userID)
		invalidateLogin(ctx, UserBanned)
		return
	}

	player := players.New(userID)
	player.Username = username
	player.SafeUsername = safeUsername
	player.Password = password
	player.Privileges = privileges
	player.IngamePrivileges = privileges.BanchoPrivileges()
	player.Token = uuid.Must(uuid.NewRandom()).String()

	ctx.Response.Header.Set("cho-token", player.Token)

	ctx.Write(packets.LoginReply(player.ID))
	ctx.Write(packets.ProtocolVersion(19))
	ctx.Write(packets.Alert("Welcome back!"))

	// Channels
	ctx.Write(packets.ChannelInfoEnd())

	channels.Mutex.RLock()
	for _, c := range channels.Values {
		if !player.Privileges.Has(c.ReadPerms) {
			continue
		} else if c.Autojoin && c.Join(player) {
			player.AddChannel(c)
			ctx.Write(packets.AutojoinChannel(c))
			ctx.Write(packets.JoinedChannel(c.Name))
		} else {
			ctx.Write(packets.AvailableChannel(c))
		}
	}
	channels.Mutex.RUnlock()

	player.SetStats(player.Gamemode, player.Relaxing)
	stats := packets.Stats(player)
	presence := packets.Presence(player)

	ctx.Write(presence)
	ctx.Write(stats)

	go func() {
		players.Mutex.RLock()
		for _, p := range players.Values {
			p.Write(presence, stats)
			player.Write(packets.Presence(p), packets.Stats(p))
		}
		players.Mutex.RUnlock()

		players.Add(player)

		var friends []int
		err = config.Database.Select(&friends, "SELECT user2 FROM users_relationships WHERE user1 = ?", player.ID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			return
		}
		pl := packets.FriendsList(friends)
		player.Write(pl)
	}()

	log.Infof("%s (%d) logged in.", player.Username, player.ID)
}
