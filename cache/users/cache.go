package users

import (
	"database/sql"
	"sync"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/utils"
)

var (
	Values map[int]*layouts.User
	Mutex  sync.RWMutex
)

func Init() {
	Mutex.Lock()
	Values = make(map[int]*layouts.User)
	Mutex.Unlock()
}

func Get(id int) *layouts.User {
	Mutex.RLock()
	if v, ok := Values[id]; ok {
		Mutex.RUnlock()
		return v
	}
	Mutex.RUnlock()

	return FetchFromDb(id)
}

func FindUsername(username string) (*layouts.User, error) {
	Mutex.RLock()
	for _, u := range Values {
		if u.Username == username || u.SafeUsername == username {
			Mutex.RUnlock()
			return u, nil
		}
	}

	Mutex.RUnlock()
	return FetchFromDbUsername(username, username)
}

func FetchFromDbUsername(username, safe string) (*layouts.User, error) {
	u := &layouts.User{}
	err := config.Database.Get(u, "SELECT id, username, username_safe, password_md5, privileges, clan_id FROM users WHERE username = ? OR username_safe = ?", username, safe)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		log.Error(err)
		return nil, err
	}

	Mutex.Lock()
	Values[u.ID] = u
	Mutex.Unlock()

	return u, nil
}

func Update(id int) {
	Mutex.RLock()
	if u, ok := Values[id]; ok {
		err := config.Database.Get(u, "SELECT id, username, username_safe, privileges, clan_id FROM users WHERE id = ?", id)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
		}
	}
	Mutex.RUnlock()
}

func FetchFromDb(id int) *layouts.User {
	u := &layouts.User{} // & > *
	err := config.Database.Get(u, "SELECT users.id, username, username_safe, password_md5, privileges, clan_id FROM users WHERE id = ?", id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
		}
		return nil
	}

	var country string
	if err = config.Database.Get(&country, "SELECT country FROM users_stats WHERE id = ?", id); err == nil {
		u.Country = utils.CountryByte(country)
	}

	Mutex.Lock()
	Values[id] = u
	Mutex.Unlock()

	return u
}
