package stats

import (
	"fmt"
	"sync"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/utils"
)

var (
	Mutex  sync.RWMutex
	Values map[int]*layouts.Stats
)

func init() {
	Mutex.Lock()
	Values = make(map[int]*layouts.Stats)
	Mutex.Unlock()
}

func Get(id int) *layouts.Stats {
	Mutex.RLock()
	if v, ok := Values[id]; ok {
		Mutex.RUnlock()
		return v
	}

	Mutex.RUnlock()
	return FetchFromDb(id)
}

func FetchFromDb(id int) *layouts.Stats {
	s := &layouts.Stats{}

	for i, table := range [2]string{"users_stats", "rx_stats"} {
		for j := byte(0); j < byte(4-i); j++ {
			var m *layouts.ModeData

			if i == 1 {
				m = &s.Relax[j]
			} else {
				m = &s.Vanilla[j]
			}

			query := fmt.Sprintf("SELECT total_score_%[1]s AS total_score, ranked_score_%[1]s AS ranked_score, pp_%[1]s AS pp, playcount_%[1]s AS playcount, avg_accuracy_%[1]s/100 AS accuracy FROM %[2]s WHERE id = ?", utils.DbMode(j), table)
			err := config.Database.Get(m, query, id)
			if err != nil {
				log.Error(err)
				return nil
			}

			err = config.Database.Get(&m.Rank, fmt.Sprintf("SELECT COUNT(id) + 1 FROM %s WHERE pp_%s > ?", table, utils.DbMode(j)), m.Performance)
			if err != nil {
				log.Error(err)
				return nil
			}
		}
	}

	Mutex.Lock()
	Values[id] = s
	Mutex.Unlock()
	return s
}

func FetchOneFromDb(id int, mode byte, relax bool) *layouts.ModeData {
	stats := Get(id)
	if stats == nil {
		return nil
	}

	var (
		m     *layouts.ModeData
		table string
	)
	if relax && mode < 3 {
		m = &stats.Relax[mode]
		table = "rx_stats"
	} else if mode < 4 {
		m = &stats.Vanilla[mode]
		table = "users_stats"
	} else {
		return nil
	}

	query := fmt.Sprintf("SELECT total_score_%[1]s AS total_score, ranked_score_%[1]s AS ranked_score, pp_%[1]s AS pp, playcount_%[1]s AS playcount, avg_accuracy_%[1]s/100 AS accuracy FROM %[2]s WHERE id = ?", utils.DbMode(mode), table)
	err := config.Database.Get(m, query, id)
	if err != nil {
		log.Error(err)
		return nil
	}

	err = config.Database.Get(&m.Rank, fmt.Sprintf("SELECT COUNT(id) + 1 FROM %s WHERE pp_%s > ?", table, utils.DbMode(mode)), m.Performance)
	if err != nil {
		log.Error(err)
		return nil
	}

	return m
}
