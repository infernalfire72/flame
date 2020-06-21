package channels

import (
	"sync"

	"github.com/infernalfire72/flame/objects"
	"github.com/infernalfire72/flame/constants"
)

var (
	Values	map[string]*objects.Channel
	Mutex	sync.RWMutex
)

func Init() {
	Mutex.Lock()
	Values = make(map[string]*objects.Channel)
	Values["#osu"] = New("#osu", "Main Channel", 0, 0, true)
	Values["#announce"] = New("#announce", "Announcements Channel", 0, 0, true)

	Mutex.Unlock()
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
	return &objects.Channel {
		Name:		name,
		Topic:		topic,
		Players:	make([]*objects.Player, 0),
		ReadPerms:	readPerms,
		WritePerms:	writePerms,
		Autojoin:	autojoin,
	}
}