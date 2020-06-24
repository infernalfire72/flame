package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchChangeHost(p *objects.Player, bytes []byte) {
	if p.Match == nil || p.Match.Host != p.ID || len(bytes) < 4 {
		return
	}

	index := bytes[0]
	if index > 15 {
		return
	}

	m := p.Match
	if slot := &m.Slots[index]; slot.User != nil && slot.Status.HasFlag(constants.SlotOccupied) {
		m.Host = slot.User.ID

		matchInfo := packets.Match(26, m)
		m.Write(packets.MatchNewHost(int(index)), matchInfo)
		lobby.Write(matchInfo)
	}
}
