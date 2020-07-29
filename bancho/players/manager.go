package players

import (
	"sync"
	"time"

	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/objects"
)

var (
	Values map[string]*objects.Player
	Mutex  sync.RWMutex
)

func init() {
	Values = make(map[string]*objects.Player)
}

func Get(token string) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return Values[token]
}

func Find(id int) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()

	for _, a := range Values {
		if a.ID == id {
			return a
		}
	}
	return nil
}

func FindUsername(name string) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()

	for _, a := range Values {
		if a.Username == name {
			return a
		}
	}
	return nil
}

func FindSafeUsername(safeName string) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()

	for _, a := range Values {
		if a.SafeUsername == safeName {
			return a
		}
	}
	return nil
}

func Add(p *objects.Player) {
	Mutex.Lock()
	Values[p.Token] = p
	Mutex.Unlock()
}

func Remove(p *objects.Player) {
	Mutex.Lock()
	delete(Values, p.Token)
	Mutex.Unlock()
}

func ForEach(fn func(*objects.Player)) {
	Mutex.RLock()
	for _, p := range Values {
		fn(p)
	}
	Mutex.RUnlock()
}

func Broadcast(data ...[]byte) {
	ForEach(func(p *objects.Player) {
		if len(p.Token) != 0 {
			p.Write(data...)
		}
	})
}

func BroadcastDelayed(data []byte, delayBy int) {
	ForEach(func(p *objects.Player) {
		if len(p.Token) != 0 {
			p.WriteDelayed(data, delayBy)
		}
	})
}

func BroadcastExcept(data []byte, ignore map[int]bool) {
	ForEach(func(p *objects.Player) {
		if !ignore[p.ID] {
			p.Write(data)
		}
	})
}

func New(user *layouts.User) *objects.Player {
	return &objects.Player{
		User:  user,
		Queue: io.NewStreamWithCapacity(1024),
		Ping:  time.Now(),
	}
}
