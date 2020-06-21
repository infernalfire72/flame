package objects

import (
	"sync"

	"github.com/infernalfire72/flame/constants"
)

type Channel struct {
	Name	string
	Topic	string
	Players	[]*Player
	Mutex	sync.RWMutex

	ReadPerms	constants.AkatsukiPrivileges
	WritePerms	constants.AkatsukiPrivileges
	Autojoin	bool
}

func (c *Channel) UserCount() int16 {
	c.Mutex.RLock() // Avoid Data Races
	defer c.Mutex.RUnlock()
	return int16(len(c.Players))
}

func (c *Channel) Join(p *Player) bool {
	if !p.Privileges.Has(c.ReadPerms) {
		return false
	}

	c.Mutex.Lock()
	c.Players = append(c.Players, p)
	c.Mutex.Unlock()
	p.AddChannel(c)
	return true
}

func (c *Channel) Leave(p *Player) {
	c.Mutex.Lock()
	for i, t := range c.Players {
		if t == p {
			c.Players[i] = c.Players[len(c.Players)-1]
			c.Players[len(c.Players)-1] = nil
			c.Players = c.Players[:len(c.Players)-1]
			break
		}
	}
	c.Mutex.Unlock()
}

func (c *Channel) AddMessage(sender *Player, message []byte) {
	c.Mutex.RLock()
	for _, receiver := range c.Players {
		if receiver.ID != sender.ID {
			receiver.Write(message)
		}
	}
	c.Mutex.RUnlock()
}