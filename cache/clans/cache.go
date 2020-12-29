package clans

import (
	"errors"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/log"
	"gorm.io/gorm"
	"sync"
)

var values = map[int]*Clan{}
var mutex sync.RWMutex

func Get(id int) *Clan {
	mutex.RLock()
	v, ok := values[id]
	mutex.RUnlock()

	if ok {
		return v
	} else {
		return Fetch(id)
	}
}

func Fetch(id int) *Clan {
	clan := &Clan{}
	if err := database.DB.First(clan, id).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
		}
		return nil
	}

	mutex.Lock()
	values[id] = clan
	mutex.Unlock()

	return clan
}