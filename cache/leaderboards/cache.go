package leaderboards

import (
	"errors"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"gorm.io/gorm"
	"sync"
)

var values = map[Identifier]*Leaderboard{}
var mutex sync.RWMutex

func Get(identifier Identifier) *Leaderboard {
	mutex.RLock()
	v, ok := values[identifier]
	mutex.RUnlock()

	if ok {
		return v
	} else {
		return Fetch(identifier)
	}
}

func Fetch(identifier Identifier) *Leaderboard {
	scores := make([]*layouts.Score, 0)

	var tableName string
	if identifier.Relax {
		tableName = "scores_relax"
	} else {
		tableName = "scores"
	}

	err := database.DB.
		Table(tableName).
		Where(&layouts.Score{
			Beatmap: identifier.Beatmap,
			Mode: identifier.Mode,
		}).
		Find(&scores).
		Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
		}
		return nil
	}

	lb := &Leaderboard{
		Identifier: identifier,
		Scores: scores,
	}

	mutex.Lock()
	values[identifier] = lb
	mutex.Unlock()

	return lb
}