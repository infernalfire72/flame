package bancho

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"
	"github.com/valyala/fasthttp"

	"github.com/infernalfire72/flame/bancho/events"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func Start(conf *config.BanchoConfig) {
	r := router.New()

	r.POST("/", banchoMain)

	port := fmt.Sprintf(":%d", conf.Port)

	log.Info("Started Bancho at", port)
	go fasthttp.ListenAndServe(port, r.Handler)
}

func banchoMain(ctx *fasthttp.RequestCtx) {
	token := string(ctx.Request.Header.Peek("osu-token"))

	if len(token) == 0 {
		events.Login(ctx)
	} else {
		p := players.Get(token)

		if p == nil {
			log.Warn("Token", token, "not found. Forcing them to relog.")
			ctx.Write(packets.Restart(100))
			ctx.SetConnectionClose()
			return
		}

		if delayed := p.GetDelayed(); delayed != nil {
			for _, packet := range delayed {
				ctx.Write(packet)
			}
		}

		s := io.StreamFrom(ctx.Request.Body())

		for s.Position+6 < len(s.Content) {
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
			case 32:
				events.JoinMatch(p, data)
			case 33:
				events.LeaveMatch(p)
			case 38:
				events.MatchChangeSlot(p, data)
			case 39:
				events.SlotReady(p)
			case 40:
				events.SlotKick(p, data)
			case 41:
				events.UpdateMatchSettings(p, data)
			case 44:
				events.StartMatch(p)
			case 47:
				events.MatchScoreFrame(p, data)
			case 49:
				events.MatchCompleted(p)
			case 51:
				events.MatchChangeMods(p, data)
			case 52:
				events.MatchLoadComplete(p)
			case 54:
				events.MatchMissingBeatmap(p)
			case 55:
				events.SlotNotReady(p)
			case 56:
				events.MatchFailed(p)
			case 59:
				events.MatchHasBeatmap(p)
			case 60:
				events.MatchSkip(p)
			case 63:
				events.JoinChannel(p, data)
			case 68:
				events.BeatmapInfoRequest(p, data)
			case 70:
				events.MatchChangeHost(p, data)
			case 73:
				events.AddFriend(p, data)
			case 74:
				events.RemoveFriend(p, data)
			case 77:
				events.MatchChangeTeam(p)
			case 78:
				events.LeaveChannel(p, data)
			case 79:
				events.UserStatsUpdateRequest(p, data)
			case 85:
				events.StatsRequest(p, data)
			case 90:
				events.MatchChangePassword(p, data)
			default:
				log.Infof("Unhandled Packet %d with length %d", id, length)
			}
		}

		p.Mutex.Lock()
		ctx.Write(p.Queue.Content)
		p.Queue.Content = p.Queue.Content[:0]
		p.Queue.Position = 0
		p.Mutex.Unlock()
	}
}
