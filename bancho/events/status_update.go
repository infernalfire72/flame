package events

import (
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/bancho/packets"
)

func StatusUpdate(p *objects.Player, bytes []byte) {
	var status layouts.Status
	if err := layouts.ReadStatus(&status, bytes); err == nil {
		p.Status = status
		p.SetRelaxing((p.Mods & 128) != 0)
		objects.Players.Broadcast(packets.Stats(p))
	}
}