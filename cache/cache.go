package cache

import (
	"github.com/infernalfire72/flame/cache/beatmaps"
	"github.com/infernalfire72/flame/cache/clans"
	"github.com/infernalfire72/flame/cache/leaderboards"
	"github.com/infernalfire72/flame/cache/users"
	"github.com/infernalfire72/flame/cache/users/stats"
)

func Init() {
	beatmaps.Init()
	leaderboards.Init()

	users.Init()
	stats.Init()
	clans.Init()
}
