package channels

import (
	"sync"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"
)

var (
	Values map[string]*objects.Channel
	Mutex  sync.RWMutex
)

func init() {
	Values = make(map[string]*objects.Channel)
	Values["#osu"] = New("#osu", "Main Channel", 0, 0, true)
	Values["#announce"] = New("#announce", "Announcements Channel", 0, 0, true)
	Values["#lobby"] = New("#lobby", "Multiplayer Discussion", 0, 0, false)
	Values["#staff"] = New("#staff", "AKATSUKI GAMER", constants.AdminManagePrivileges, 0, true)
}

func Get(name string) *objects.Channel {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return Values[name]
}

func Add(c *objects.Channel) {
	Mutex.Lock()
	Values[c.Name] = c
	Mutex.Unlock()
}

func New(name, topic string, readPerms, writePerms constants.AkatsukiPrivileges, autojoin bool) *objects.Channel {
	return &objects.Channel{
		Name:       name,
		Topic:      topic,
		Players:    make([]*objects.Player, 0),
		ReadPerms:  readPerms,
		WritePerms: writePerms,
		Autojoin:   autojoin,
	}
}
