package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchLoadComplete(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	if slot, _ := m.FindPlayerSlot(p); slot != nil && slot.Status.HasFlag(constants.SlotPlaying) {
		slot.Loaded = true

		if m.CheckLoaded() {
			m.WritePlaying(packets.MatchLoadComplete())
		}
	}
}
