package events

import (
	"encoding/binary"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchChangeMods(p *objects.Player, bytes []byte) {
	if p.Match == nil || len(bytes) != 4 {
		return
	}

	m := p.Match
	mods := constants.Mod(binary.LittleEndian.Uint32(bytes))
	relax := (mods & constants.ModRelax) != 0

	if m.FreeMod {
		if dt := mods & constants.ModsChangeSpeed; m.Host == p.ID {
			m.Mods = dt
		}

		mods &= ^constants.ModsChangeSpeed

		if slot, _ := m.FindPlayerSlot(p); slot != nil {
			slot.Mods = mods
		}
		p.SetRelaxing(relax)
	} else if m.Host == p.ID {
		m.Mods = mods

		m.ForSlots(func(s *objects.MultiplayerSlot) {
			if s.User != nil {
				s.User.SetRelaxing(relax)
			}
		})
	} else {
		return
	}

	matchInfo := packets.Match(26, m)
	m.Write(matchInfo)
	lobby.Write(matchInfo)
}
