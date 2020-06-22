package events

import (
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
	"github.com/infernalfire72/flame/bancho/lobby"
)

func LeaveLobby(p *objects.Player) {
	if c := channels.Get("#lobby"); c != nil {
		c.Leave(p)
		p.RemoveChannel(c)
		lobby.RemovePlayer(p)
		p.Write(packets.ChannelRevoked("#lobby"))
		players.Broadcast(packets.AvailableChannel(c))
	}
}