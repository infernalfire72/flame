package users

import (
	"errors"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"gorm.io/gorm"
	"sync"
)

var values = map[int]*layouts.User{}
var mutex sync.RWMutex

func Get(id int) *layouts.User {
	mutex.RLock()

	v, ok := values[id]
	mutex.RUnlock()

	if ok {
		return v
	} else {
		return Fetch(id)
	}
}

func FindUsername(name string) *layouts.User {
	mutex.RLock()

	for _, u := range values {
		if u.Username == name {
			mutex.RUnlock()
			return u
		}
	}

	mutex.RUnlock()
	return FetchWithUsername(name)
}

func FetchWithUsername(name string) *layouts.User {
	user := &layouts.User{Username: name, SafeUsername: name}
	if err := database.DB.Where(user).First(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
		}
		return nil
	}

	mutex.Lock()
	values[user.ID] = user
	mutex.Unlock()

	return user
}

func Fetch(id int) *layouts.User {
	user := &layouts.User{}
	if err := database.DB.First(user, id).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
		}
		return nil
	}

	mutex.Lock()
	values[user.ID] = user
	mutex.Unlock()

	return user
}