package events

import (
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func StartMatch(p *objects.Player /*, bytes []byte*/) {
	if p.Match == nil || p.Match.Host != p.ID {
		return
	}

	m := p.Match
	// m.ReadMatch(bytes) // We might not even need this
	m.Running = true

	for i := 0; i < len(m.Slots); i++ {
		if m.Slots[i].User != nil && m.Slots[i].Status.HasFlag(constants.SlotNotReady|constants.SlotReady) {
			m.Slots[i].Status = constants.SlotPlaying
		}
	}

	matchInfo := packets.Match(46, m)
	m.Write(matchInfo)
	lobby.Write(matchInfo)

	log.Info(m, "started playing.")
}
