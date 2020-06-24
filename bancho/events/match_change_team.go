package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchChangeTeam(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	if m.TeamType < constants.TeamVs {
		return
	}

	if slot, _ := m.FindPlayerSlot(p); slot != nil {
		if slot.Team == 1 {
			slot.Team = 2
		} else {
			slot.Team = 1
		}

		matchInfo := packets.Match(26, m)
		m.Write(matchInfo)
	}
}
