package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchCompleted(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	if slot, _ := m.FindPlayerSlot(p); slot != nil && slot.Status.HasFlag(constants.SlotPlaying) {
		slot.Completed = true

		if m.CheckCompleted() {
			m.Running = false

			m.Mutex.RLock()

			for i := 0; i < len(m.Slots); i++ {
				if m.Slots[i].User != nil && m.Slots[i].Status.HasFlag(constants.SlotPlaying) {
					m.Slots[i].User.Write(packets.MatchFinished())
					m.Slots[i].Status = constants.SlotNotReady
				}
			}

			m.Mutex.RUnlock()

			matchInfo := packets.Match(26, m)
			m.Write(matchInfo)
			lobby.Write(matchInfo)
		}
	}
}
