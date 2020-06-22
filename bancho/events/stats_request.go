package events

import (
	"unsafe"

	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func StatsRequest(p *objects.Player, bytes []byte) {
	if len(bytes) < 6 {
		return
	}

	for i := 2; i < len(bytes); i += 4 {
		id := *(*int32)(unsafe.Pointer(&bytes[i]))
		if target := players.Find(int(id)); target != nil {
			p.Write(packets.Stats(target))
		}
	}
}
