package events

import (
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func SpectateFrames(p *objects.Player, bytes []byte) {
	p.SpectatorMutex.RLock()

	packet := packets.SpectateFrames(bytes)
	for _, t := range p.Spectators {
		t.Write(packet)
	}
	
	p.SpectatorMutex.RUnlock()
}