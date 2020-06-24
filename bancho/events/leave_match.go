package events

import (
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/matches"
	"github.com/infernalfire72/flame/bancho/packets"
)

func LeaveMatch(p *objects.Player) {
	if p.Match == nil {
		return
	}

	m := p.Match
	m.RemovePlayer(p)

	if m.UserCount() == 0 {
		lobby.Write(packets.Match(28, m))
		matches.Disband(m)
	} else {
		// Find a new Host
		if m.Host == p.ID {
			if slot, index := m.FindNextPlayer(); slot != nil {
				m.Host = slot.User.ID
				m.Write(packets.MatchNewHost(int(index)))
			}
		}

		matchInfo := packets.Match(26, m)
		channelInfo := packets.AvailableChannelArgs("#multiplayer", "Multiplayer Channel", m.UserCount())
		m.Write(matchInfo, channelInfo)
		lobby.Write(matchInfo)
	}

	p.Write(packets.ChannelRevoked("#multiplayer"))

	log.Info(p, "left", m)
}
