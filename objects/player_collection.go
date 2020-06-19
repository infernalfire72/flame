package objects

import (
	"sync"
)

type PlayerCollection struct {
	Players	map[string]*Player
	Mutex	sync.RWMutex
}

func (c *PlayerCollection) Get(token string) *Player {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.Players[token]
}

func (c *PlayerCollection) FindPlayer(id int) *Player {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	for _, a := range c.Players {
		if a.ID == id {
			return a
		}
	}
	return nil
}

func (c *PlayerCollection) FindPlayerName(name string) *Player {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	for _, a := range c.Players {
		if a.Username == name {
			return a
		}
	}
	return nil
}

func (c *PlayerCollection) FindPlayerNameSafe(safeName string) *Player {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	for _, a := range c.Players {
		if a.SafeUsername == safeName {
			return a
		}
	}
	return nil
}

func (c *PlayerCollection) AddPlayer(p *Player) {
	c.Mutex.Lock()
	c.Players[p.Token] = p
	c.Mutex.Unlock()
}

func (c *PlayerCollection) RemovePlayer(p *Player) {
	c.Mutex.Lock()
	c.Players[p.Token] = nil
	c.Mutex.Unlock()
}

func (c *PlayerCollection) Broadcast(data []byte) {
	c.Mutex.RLock()

	for _, a := range c.Players {
		a.Write(data)
	}

	c.Mutex.RUnlock()
}

func (c *PlayerCollection) BroadcastExcept(data []byte, ignore map[int]bool) {
	c.Mutex.RLock()

	for _, a := range c.Players {
		if !ignore[a.ID] {
			a.Write(data)
		}
	}

	c.Mutex.RUnlock()
}