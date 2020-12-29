package events

import (
	"database/sql"
	"errors"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/layouts"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"
	"github.com/infernalfire72/flame/utils"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"

	"github.com/infernalfire72/flame/cache/users"
	"github.com/infernalfire72/flame/cache/users/stats"
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

	s := &utils.Stopwatch{}
	s.Start()

	u := users.FindUsername(username)
	if u == nil || !u.VerifyPassword(password) {
		invalidateLogin(ctx, InvalidLoginData)
		return
	}

	if players.FindUsername(u.Username) != nil {
		ctx.Write(packets.Alert("You are already logged in on another client."))
		invalidateLogin(ctx, ServerError)
		return
	}

	if (u.Privileges & constants.UserPendingVerification) != 0 {
		u.Privileges &= ^constants.UserPendingVerification
		u.Privileges |= constants.UserNormal | constants.UserPublic
		database.DB.Save(u)
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

	var err error

	if hashSet[4] == "runningunderwine" {
		err = database.DB.
			Where("userid <> ? AND activated AND unique_id = ?", u.ID, hashSet[3]).
			Find(&layouts.HwidMatch{}).
			Error
	} else {
		err = database.DB.
			Where("userid <> ? AND activated AND unique_id = ?", u.ID, hashSet[3]).
			Find(&layouts.HwidMatch{}).
			Error
	}

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			invalidateLogin(ctx, ServerError)
			return
		}
	} else {
		// Do stuff (restrict etc.)
		// Restrict(match)
		// Ban(userID)
		invalidateLogin(ctx, UserBanned)
		return
	}

	player := players.New(u)
	player.IngamePrivileges = u.Privileges.BanchoPrivileges()
	player.Token = uuid.Must(uuid.NewRandom()).String()
	player.LoginTime = time.Now()

	ctx.Response.Header.Set("cho-token", player.Token)

	ctx.Write(packets.LoginReply(player.ID))
	ctx.Write(packets.ProtocolVersion(19))
	ctx.Write(packets.UserPrivileges(int(player.IngamePrivileges)))
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

	if s := stats.Get(stats.Identifier{
		User:  u.ID,
		Mode:  0,
		Relax: false,
	}); s != nil {
		player.Stats[0] = s
	}

	stats := packets.Stats(player)
	presence := packets.Presence(player)

	ctx.Write(presence)
	ctx.Write(stats)

	go func() {
		players.ForEach(func(p *objects.Player) {
			if len(p.Token) != 0 {
				p.Write(presence, stats)
			}

			player.Write(packets.Presence(p), packets.Stats(p))
		})

		players.Add(player)

		friends := utils.GetFriends(u.ID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			return
		}
		pl := packets.FriendsList(friends)
		player.Write(pl)
	}()

	s.Stop()
	log.Infof("%s (%d) logged in. | Elapsed: %s", player.Username, player.ID, s.ElapsedReadable())
}
