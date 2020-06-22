package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func ChangeSlot(p *objects.Player, bytes []byte) {
	if p.Match == nil || len(bytes) < 2 {
		return
	}

	target := bytes[0]
	if target > 15 {
		return
	}

	m := p.Match
	slot := &m.Slots[target]
	if slot.User != nil || slot.Status != constants.SlotEmpty {
		return
	}

	if pSlot := m.FindPlayerSlot(p); pSlot != nil {
		pSlot.Clear()
	} else {
		log.Warn("Player", p, "had no slot?", m)
	}

	slot.Status = constants.SlotNotReady
	slot.User = p

	matchInfo := packets.Match(26, m)
	m.Write(matchInfo)
	lobby.Write(matchInfo)
}
