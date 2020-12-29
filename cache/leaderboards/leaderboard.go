package leaderboards

import (
	"github.com/infernalfire72/flame/cache/users"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/utils"
	"sync"
)

type Identifier struct {
	Beatmap string
	Relax   bool
	Mode    constants.Mode
}

type Leaderboard struct {
	Identifier
	Scores []*layouts.Score
	Mutex  sync.RWMutex
}

func (l *Leaderboard) ModsFilter(mods constants.Mod) (scores []*layouts.Score) {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()
	for _, v := range l.Scores {
		if mods == v.Mods {
			scores = append(scores, v)
		}
	}

	return
}

func (l *Leaderboard) FriendsFilter(id int) (scores []*layouts.Score) {
	friends := utils.GetFriends(id)

	l.Mutex.RLock()
	defer l.Mutex.RUnlock()

	for _, v := range l.Scores {
		if v.User == id || utils.Has(friends, v.User) {
			scores = append(scores, v)
		}
	}

	return
}

func (l *Leaderboard) CountryFilter(country string)  (scores []*layouts.Score) {
	l.Mutex.RLock()
	defer l.Mutex.RUnlock()

	for _, v := range l.Scores {
		if u := users.Get(v.User); u != nil && u.Country == country {
			scores = append(scores, v)
		}
	}

	return
}