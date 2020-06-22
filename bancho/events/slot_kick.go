package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)


func SlotKick(p *objects.Player, bytes []byte) {
	if len(bytes) < 2 || p.Match == nil || p.Match.Host != p.ID {
		return
	}

	target := bytes[0]
	if target > 15 {
		return
	}

	m := p.Match
	slot := &m.Slots[target]
	if slot.User != nil {
		slot.Clear()
		slot.Status = constants.SlotLocked
	} else if slot.Status == constants.SlotEmpty {
		slot.Status = constants.SlotLocked
	} else {
		slot.Status = constants.SlotEmpty
	}

	matchInfo := packets.Match(26, m)
	m.Write(matchInfo)
	lobby.Write(matchInfo)
}