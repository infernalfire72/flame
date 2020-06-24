package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchScoreFrame(p *objects.Player, bytes []byte) {
	if p.Match == nil {
		return
	}

	m := p.Match

	if slot, index := m.FindPlayerSlot(p); slot != nil && slot.Status.HasFlag(constants.SlotPlaying) {
		bytes[4] = byte(index)

		m.WritePlaying(packets.ScoreFrame(bytes))
	}
}
