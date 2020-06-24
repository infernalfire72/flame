package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchMissingBeatmap(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	if slot, _ := m.FindPlayerSlot(p); slot != nil {
		slot.Status = constants.SlotMissingBeatmap

		matchInfo := packets.Match(26, m)
		m.Write(matchInfo)
		lobby.Write(matchInfo)
	}
}
