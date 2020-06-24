package events

import (
	"time"

	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func Logout(p *objects.Player) {
	if now := time.Now(); now.Sub(p.LoginTime).Seconds() < 1 {
		return
	}

	players.Remove(p)
	players.Broadcast(packets.LogoutNotice(p.ID, 0))

	for _, c := range p.Channels {
		c.Leave(p)
	}

	if p.Spectating != nil {
		StopSpectating(p)
	}

	if p.Match != nil {
		LeaveMatch(p)
	}

	if p.IsLobby {
		lobby.RemovePlayer(p)
	}

	log.Info(p.Username, "logged out.")
}
