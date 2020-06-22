package events

import (
	"github.com/infernalfire72/flame/objects"
)

func StatsUpdateRequest(p *objects.Player) {
	p.SetStats(p.Gamemode, p.Relaxing)
}
