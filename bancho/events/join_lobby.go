package events

import (
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func JoinLobby(p *objects.Player) {
	if c := channels.Get("#lobby"); c != nil && c.Join(p) {
		p.AddChannel(c)
		lobby.AddPlayer(p)
		p.Write(packets.JoinedChannel("#lobby"))
		players.Broadcast(packets.AvailableChannel(c))
	}
}
