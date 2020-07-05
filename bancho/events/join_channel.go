package events

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func JoinChannel(p *objects.Player, bytes []byte) {
	s := io.Stream{bytes, 0}

	target, err := s.ReadString()
	if err != nil || len(target) == 0 || target[0] != '#' {
		return
	}

	if c := channels.Get(target); c != nil && c.Join(p) {
		p.AddChannel(c)
		p.Write(packets.JoinedChannel(c.Name))

		players.Broadcast(packets.AvailableChannel(c))
	} else {
		log.Warn(p, "failed joining", target)
	}
}
