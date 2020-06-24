package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchFailed(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	if slot, index := m.FindPlayerSlot(p); slot != nil && slot.Status.HasFlag(constants.SlotPlaying) {
		m.WritePlaying(packets.MatchPlayerFailed(index))
	}
}
