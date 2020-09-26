package leaderboards

import (
	"github.com/infernalfire72/flame/utils"
	"sort"
	"sync"

	"github.com/infernalfire72/flame/cache/beatmaps"
	"github.com/infernalfire72/flame/cache/users"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
)

type Identifier struct {
	Md5   string
	Mode  constants.Mode
	Relax bool
}

type Scores []*layouts.Score

func (s Scores) GetPersonalBest(id int) (*layouts.Score, int) {
	for i, v := range s {
		if u := users.Get(v.UserID); u != nil && u.ID == id {
			return v, i
		}
	}
	return nil, -1
}

type Leaderboard struct {
	BeatmapMd5 string
	Scores     Scores
	Mode       constants.Mode
	Relax      bool
	Mutex      sync.RWMutex
}

func (l *Leaderboard) Map() *beatmaps.Beatmap {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	return beatmaps.Get(l.BeatmapMd5)
}

func (l *Leaderboard) Sort() {
	if m := l.Map(); m != nil && m.Status >= constants.StatusRanked {
		l.Mutex.Lock()
		sort.Slice(l.Scores, func(i, j int) bool {
			if !l.Relax || m.Status == constants.StatusLoved {
				return l.Scores[i].TotalScore > l.Scores[j].TotalScore
			} else {
				return l.Scores[i].Performance > l.Scores[j].Performance
			}
		})
		l.Mutex.Unlock()
	}
}

func (l *Leaderboard) Country(country byte) (scores Scores) {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()
	for _, v := range l.Scores {
		if u := users.Get(v.UserID); u != nil && u.Country == country {
			scores = append(scores, v)
		}
	}

	return
}

func (l *Leaderboard) Mods(mods constants.Mod) (scores Scores) {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()

	for _, v := range l.Scores {
		if v.Mods == mods {
			scores = append(scores, v)
		}
	}

	return
}

func (l *Leaderboard) Friends(id int) (scores Scores) {
	friends := utils.GetFriends(id)

	l.Mutex.RLock()
	defer l.Mutex.RUnlock()

	for _, v := range l.Scores {
		if utils.Has(friends, v.UserID) || v.UserID == id {
			scores = append(scores, v)
		}
	}

	return
}

func (l *Leaderboard) Count() int {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()

	return len(l.Scores)
}

func (l *Leaderboard) AddScore(score *layouts.Score) {
	l.RemoveUser(score.UserID)

	l.Mutex.Lock()
	l.Scores = append(l.Scores, score)
	l.Mutex.Unlock()

	l.Sort()
}

func (l *Leaderboard) FindUserScore(id int) (*layouts.Score, int) {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()

	for i, score := range l.Scores {
		if score.UserID == id {
			return score, i
		}
	}

	return nil, -1
}

// Method used for wiping a users scores
func (l *Leaderboard) RemoveUser(id int) {
	if _, i := l.FindUserScore(id); i != -1 {
		l.Mutex.Lock()
		copy(l.Scores[i:], l.Scores[i+1:])
		l.Scores[len(l.Scores)-1] = nil // or the zero value of T
		l.Scores = l.Scores[:len(l.Scores)-1]
		l.Mutex.Unlock()
	}
}

// Fetches all scores from the database
func (l *Leaderboard) FetchFromDb() {
	Scores := make(Scores, 0)

	if m := l.Map(); m != nil && m.Status >= constants.StatusRanked {
		table := "scores"
		if l.Relax {
			table = "scores_relax"
		}

		/*tableSort := "score"
		if l.Relax && m.Status != constants.StatusLoved {
			tableSort = "pp"
		}*/

		err := config.Database.Select(&Scores, `SELECT id, userid, score, pp, max_combo, full_combo,
			mods, 300_count, 100_count, 50_count, katus_count, gekis_count, misses_count, time
			FROM `+table+`
			WHERE beatmap_md5 = ? AND completed = 3 AND play_mode = ?`, l.BeatmapMd5, l.Mode)
		if err != nil {
			log.Error(err)
			return
		}
	}

	l.Mutex.Lock()
	l.Scores = Scores
	l.Mutex.Unlock()
}