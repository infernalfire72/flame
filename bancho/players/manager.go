package players

import (
	"sync"
	"time"

	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/objects"
)

var (
	Values	map[string]*objects.Player
	Mutex	sync.RWMutex
)

func Get(token string) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return Values[token]
}

func FindPlayer(id int) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()

	for _, a := range Values {
		if a.ID == id {
			return a
		}
	}
	return nil
}

func FindPlayerName(name string) *objects.Player {
	Mutex.RLock()
	defer Mutex.RUnlock()

	for _, a := range Values {
		if a.Username == name {
			return a
		}
	}
	return nil
}

func FindPlayerNameSafe(safeName string) *objects.Player {
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

func Broadcast(data []byte) {
	Mutex.RLock()

	for _, a := range Values {
		a.Write(data)
	}

	Mutex.RUnlock()
}

func BroadcastExcept(data []byte, ignore map[int]bool) {
	Mutex.RLock()

	for _, a := range Values {
		if !ignore[a.ID] {
			a.Write(data)
		}
	}

	Mutex.RUnlock()
}

func New(id int) *objects.Player {
	return &objects.Player {
		ID:		id,
		Queue:	io.NewStreamWithCapacity(1024),
		Ping:	time.Now(),
	}
}