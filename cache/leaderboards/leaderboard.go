package leaderboards

import (
	"sort"
	"sync"

	"github.com/infernalfire72/flame/cache/beatmaps"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/utils"
)

type Identifier struct {
	Md5   string
	Mode  byte
	Relax bool
}

type Scores []*Score

func (s *Scores) GetPersonalBest(id int) (*Score, int) {
	for i, v := range *s {
		if u := v.User(); u != nil && u.ID == id {
			return v, i
		}
	}
	return nil, -1
}

type Leaderboard struct {
	BeatmapMd5 string
	Scores     []*Score
	Mode       byte
	Relax      bool
	Mutex      sync.RWMutex
}

func (l *Leaderboard) Map() *beatmaps.Beatmap {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	return beatmaps.Get(l.BeatmapMd5)
}

func (l *Leaderboard) Count() int {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	return len(l.Scores)
}

func (l *Leaderboard) AddScore(s *Score) {
	l.RemoveUser(s.UserID)

	l.Mutex.Lock()
	l.Scores = append(l.Scores, s)
	l.Mutex.Unlock()
	l.Sort()
}

func (l *Leaderboard) AddFromLayout(s *layouts.Score) {
	l.AddScore(&Score{
		ID:          s.ID,
		UserID:      s.UserID,
		Score:       s.TotalScore,
		Performance: s.Performance,
		Combo:       s.Combo,
		FullCombo:   s.FullCombo,
	})
}

func (l *Leaderboard) Sort() {
	if m := l.Map(); m != nil && m.Status >= constants.StatusRanked {
		l.Mutex.Lock()
		sort.Slice(l.Scores, func(i, j int) bool {
			if !l.Relax || m.Status == constants.StatusLoved {
				return l.Scores[i].Score > l.Scores[j].Score
			} else {
				return l.Scores[i].Performance > l.Scores[j].Performance
			}
		})
		l.Mutex.Unlock()
	}
}

func (l *Leaderboard) Country(country byte) Scores {
	scores := make(Scores, 0)

	l.Mutex.RLock()
	for _, v := range l.Scores {
		if u := v.User(); u != nil && u.Country == country {
			scores = append(scores, v)
		}
	}
	l.Mutex.RUnlock()

	return scores
}

func (l *Leaderboard) Mods(mods int) Scores {
	scores := make(Scores, 0)

	l.Mutex.RLock()
	for _, v := range l.Scores {
		if v.Mods == mods {
			scores = append(scores, v)
		}
	}
	l.Mutex.RUnlock()

	return scores
}

func (l *Leaderboard) Friends(id int) Scores {
	friends := utils.GetFriends(id)
	scores := make(Scores, 0)

	l.Mutex.RLock()
	for _, v := range l.Scores {
		if utils.Has(friends, v.UserID) || v.UserID == id {
			scores = append(scores, v)
		}
	}
	l.Mutex.RUnlock()

	return scores
}

func (l *Leaderboard) FindUserScore(id int) (*Score, int) {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()
	for i, score := range l.Scores {
		if score.UserID == id {
			return score, i
		}
	}
	return nil, -1
}

func (l *Leaderboard) RemoveUser(id int) {
	if s, i := l.FindUserScore(id); s != nil {
		l.RemoveScoreIndex(i)
	}
}

func (l *Leaderboard) RemoveScoreIndex(i int) {
	l.Mutex.Lock()
	copy(l.Scores[i:], l.Scores[i+1:])
	l.Scores[len(l.Scores)-1] = nil // or the zero value of T
	l.Scores = l.Scores[:len(l.Scores)-1]
	l.Mutex.Unlock()
}

func (l *Leaderboard) FetchFromDb() {
	Scores := make([]*Score, 0)

	if m := l.Map(); m != nil && m.Status >= constants.StatusRanked {
		table := "scores"
		if l.Relax {
			table = "scores_relax"
		}

		tableSort := "score"
		if l.Relax && m.Status != constants.StatusLoved {
			tableSort = "pp"
		}

		rows, err := config.Database.Query("SELECT "+table+".id, userid, score, pp, max_combo, full_combo, mods, 300_count, 100_count, 50_count, katus_count, gekis_count, misses_count, time FROM "+table+" LEFT JOIN users ON users.id = userid LEFT JOIN clans ON clans.id = users.clan_id WHERE beatmap_md5 = ? AND completed = 3 AND play_mode = ? AND users.privileges & 1 ORDER BY "+tableSort+" DESC", l.BeatmapMd5, l.Mode)
		if err != nil {
			log.Error(err)
		}
		defer rows.Close()

		for rows.Next() {
			s := &Score{}
			err = rows.Scan(&s.ID, &s.UserID, &s.Score, &s.Performance, &s.Combo, &s.FullCombo, &s.Mods, &s.N300, &s.N100, &s.N50, &s.NKatu, &s.NGeki, &s.NMiss, &s.Timestamp)
			if err != nil {
				log.Error(err)
			}
			Scores = append(Scores, s)
		}
	}

	l.Mutex.Lock()
	l.Scores = Scores
	l.Mutex.Unlock()
}
