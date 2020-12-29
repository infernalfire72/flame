package events

import (
	"github.com/infernalfire72/flame/objects"
)

func UserStatsUpdateRequest(p *objects.Player, bytes []byte) {
	/*if len(bytes) < 4 {
		return
	}

	id := *(*int32)(unsafe.Pointer(&bytes[0]))
	if target := players.Find(int(id)); target != nil {
		stats.FetchOneFromDb(target.ID, target.Gamemode, target.Relaxing)
		p.Write(packets.Stats(target))
	}*/
}
