package events

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func LeaveChannel(p *objects.Player, bytes []byte) {
	s := io.Stream{bytes, len(bytes), len(bytes), 0}

	target, err := s.ReadString()
	if err != nil || len(target) == 0 || target[0] != '#' {
		return
	}

	if c := channels.Get(target); c != nil {
		c.Leave(p)
		p.RemoveChannel(c)

		players.Broadcast(packets.AvailableChannel(c))
	} else {
		log.Warn(p, "tried to leave non-existent channel", target)
	}
}
