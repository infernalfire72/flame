package lobby

import (
	"sync"

	"github.com/infernalfire72/flame/objects"
)

var (
	Players	[]*objects.Player
	Mutex	sync.RWMutex
)

func AddPlayer(p *objects.Player) {
	Mutex.Lock()
	Players = append(Players, p)
	Mutex.Unlock()
	p.IsLobby = true
}

func RemovePlayer(p *objects.Player) {
	Mutex.Lock()
	for i, t := range Players {
		if t == p {
			Players[i] = Players[len(Players)-1]
			Players[len(Players)-1] = nil
			Players = Players[:len(Players)-1]
		}
	}
	Mutex.Unlock()
	p.IsLobby = false
}

func Write(data []byte) {
	Mutex.RLock()
	for _, t := range Players {
		t.Write(data)
	}
	Mutex.RUnlock()
}