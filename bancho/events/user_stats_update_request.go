package events

import (
	"unsafe"

	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func UserStatsUpdateRequest(p *objects.Player, bytes []byte) {
	if len(bytes) < 4 {
		return
	}

	id := *(*int32)(unsafe.Pointer(&bytes[0]))
	if target := players.Find(int(id)); target != nil {
		target.SetStats(target.Gamemode, target.Relaxing)
		p.Write(packets.Stats(target))
	}
}
