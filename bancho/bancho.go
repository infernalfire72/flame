package bancho

import (
	"fmt"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/events"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func Start(conf *config.BanchoConfig) {
	channels.Init()

	r := router.New()

	r.POST("/", banchoMain)

	port := fmt.Sprintf(":%d", conf.Port)

	log.Info("Started Bancho at", port)
	fasthttp.ListenAndServe(port, r.Handler)
}

func banchoMain(ctx *fasthttp.RequestCtx) {
	token := string(ctx.Request.Header.Peek("osu-token"))

	if len(token) == 0 {
		events.Login(ctx)
	} else {
		p := players.Get(token)

		if p == nil {
			ctx.SetStatusCode(http.StatusUnauthorized)
			ctx.Write(packets.LoginReply(-5))
			return
		}

		s := io.StreamFrom(ctx.Request.Body())

		for s.Position + 6 < s.Length {
			id, _ := s.ReadInt16()
			// TODO: do we need actual handling here?

			s.ReadByte()
			length, _ := s.ReadInt32()

			data, err := s.ReadSegment(int(length))
			if err != nil {
				break
			}

			switch id {
			case 0:
				events.StatusUpdate(p, data)
			case 1:
				events.IrcMessage(p, data)
			case 2:
				events.Logout(p)
			case 3:
				events.StatsUpdateRequest(p)
			case 4:
				events.Ping(p)
			case 16:
				events.StartSpectating(p, data)
			case 17:
				events.StopSpectating(p)
			case 18:
				events.SpectateFrames(p, data)
			case 25:
				events.PrivateMessage(p, data)
			case 29:
				events.LeaveLobby(p)
			case 30:
				events.JoinLobby(p)
			case 31:
				events.CreateMatch(p, data)
			default:
				log.Infof("Unhandled Packet %d with length %d", id, length)
			}
		}

		p.Mutex.Lock()
		ctx.Write(p.Queue.Data())
		p.Queue.Length = 0
		p.Queue.Position = 0
		p.Mutex.Unlock()
	}
}