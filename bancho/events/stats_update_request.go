package events

import (
	"github.com/infernalfire72/flame/cache/users/stats"
	"github.com/infernalfire72/flame/objects"
)

func StatsUpdateRequest(p *objects.Player) {
	stats.FetchOneFromDb(p.ID, p.Gamemode, p.Relaxing)
}
