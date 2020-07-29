package leaderboards

import (
	"sync"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
)

var (
	Mutex  sync.RWMutex
	Values map[Identifier]*Leaderboard
)

func init() {
	Mutex.Lock()
	Values = make(map[Identifier]*Leaderboard)
	Mutex.Unlock()
}

func Get(id Identifier) *Leaderboard {
	Mutex.RLock()
	if v, ok := Values[id]; ok {
		Mutex.RUnlock()
		return v
	}
	Mutex.RUnlock()
	return FetchFromDb(id)
}

func FetchFromDb(identifier Identifier) *Leaderboard {
	lb := &Leaderboard{
		BeatmapMd5: identifier.Md5,
		Mode:       identifier.Mode,
		Relax:      identifier.Relax,
	}

	lb.FetchFromDb()
	Mutex.Lock()
	defer Mutex.Unlock()

	Values[identifier] = lb
	return lb
}

func RemoveUser(id int) {
	Mutex.RLock()
	for _, a := range Values {
		a.RemoveUser(id)
	}
	Mutex.RUnlock()
}

func RemoveUserWithIdentifier(id int, rx bool) {
	Mutex.RLock()
	for _, a := range Values {
		if a.Relax == rx {
			a.RemoveUser(id)
		}
	}
	Mutex.RUnlock()
}

func AddUser(id int) {
	for i, table := range [...]string{"scores", "scores_relax"} {
		var relax bool
		if i == 1 {
			relax = true
		}

		rows, err := config.Database.Query("SELECT "+table+".id, userid, score, pp, COALESCE(CONCAT('[', tag, '] ', username), username) AS username, max_combo, full_combo, mods, 300_count, 100_count, 50_count, katus_count, gekis_count, misses_count, time, play_mode, beatmap_md5 FROM "+table+" LEFT JOIN users ON users.id = userid LEFT JOIN clans ON clans.id = users.clan_id WHERE userid = ? AND completed = 3", id)
		if err != nil {
			log.Error(err)
		}
		defer rows.Close()

		for rows.Next() {
			s := &Score{}
			var (
				md5  string
				mode byte
			)
			err = rows.Scan(&s.ID, &s.UserID, &s.Score, &s.Performance, &s.Combo, &s.FullCombo, &s.Mods, &s.N300, &s.N100, &s.N50, &s.NKatu, &s.NGeki, &s.NMiss, &s.Timestamp, &mode, &md5)
			if err != nil {
				log.Error(err)
			}

			id := Identifier{md5, mode, relax}
			Mutex.RLock()
			if value, ok := Values[id]; ok {
				value.AddScore(s)
			}
			Mutex.RUnlock()
		}
	}
}
