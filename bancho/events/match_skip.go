package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchSkip(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	if slot, _ := m.FindPlayerSlot(p); slot != nil && slot.Status.HasFlag(constants.SlotPlaying) {
		slot.Skipped = true

		if m.CheckSkipped() {
			m.WritePlaying(packets.MatchSkip())
		}
	}
}
