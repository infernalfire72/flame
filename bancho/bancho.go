package bancho

import (
	"fmt"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/events"
	"github.com/infernalfire72/flame/bancho/packets"
)

func Start(conf *config.BanchoConfig) {
	objects.Players = make(map[string]*objects.Player)
	objects.Channels = make(map[string]*objects.Channel)
	objects.Matches = make(map[uint16]*objects.MultiplayerLobby)

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
		p := objects.Players[token]

		if p == nil {
			ctx.SetStatusCode(http.StatusUnauthorized)
			ctx.Write(packets.LoginReply(-5))
			return
		}

		s := io.StreamFrom(ctx.Request.Body())

		for s.Position + 6 < s.Length {
			id, err := s.ReadInt16()
			// TODO: do we need actual handling here?
			if err != nil {
				break
			}

			s.ReadByte()
			length, err := s.ReadInt32()
			if err != nil {
				break
			}

			data, err := s.ReadSegment(int(length))
			if err != nil {
				break
			}

			switch id {
			case 0:
				events.StatusUpdate(p, data)
			case 4:
				events.Ping(p)
			default:
				log.Infof("Unhandled Packet %d with length %d", id, length)
			}
		}

		p.Mutex.Lock()
		ctx.Write(p.Queue.Data())
		p.Mutex.Unlock()
	}
}