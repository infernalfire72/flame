package clans

import (
	"database/sql"
	"sync"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
)

type Clan struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Tag         string `db:"tag"`
	Description string `db:"description"`
	Owner       int    `db:"owner"`

	Members []int
}

var (
	Mutex  sync.RWMutex
	Values map[int]*Clan
)

func Init() {
	Mutex.Lock()
	Values = make(map[int]*Clan)
	Mutex.Unlock()
}

func Get(id int) *Clan {
	Mutex.RLock()
	if v, ok := Values[id]; ok {
		Mutex.RUnlock()
		return v
	}

	Mutex.RUnlock()
	return FetchFromDb(id)
}

func FetchFromDb(id int) *Clan {
	c := &Clan{}

	err := config.Database.Get(c, "SELECT id, name, tag, description, owner FROM clans WHERE id = ?", id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
		}

		return nil
	}

	err = config.Database.Select(&c.Members, "SELECT id FROM users WHERE clan_id = ?", id)
	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		return nil
	}

	Mutex.Lock()
	Values[id] = c
	Mutex.Unlock()

	return c
}
