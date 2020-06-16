package objects

import (
	"sync"

	"github.com/infernalfire72/flame/constants"
)

var ChannelMutex sync.RWMutex
var Channels map[string]*Channel

type Channel struct {
	Name	string
	Topic	string
	Players	[]*Player

	ReadPerms	constants.AkatsukiPrivileges
	WritePerms	constants.AkatsukiPrivileges
	Autojoin	bool
}

func (c *Channel) UserCount() int16 {
	return int16(len(c.Players))
}

func (c *Channel) Join(p *Player) bool {
	if !p.Privileges.Has(c.ReadPerms) {
		return false
	}
	
	c.Players = append(c.Players, p)
	return true
}