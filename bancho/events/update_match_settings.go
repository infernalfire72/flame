package events

import (
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func UpdateMatchSettings(p *objects.Player, bytes []byte) {
	if p.Match == nil || p.Match.Host != p.ID {
		return
	}

	m := p.Match
	err := m.ReadMatch(bytes)
	if err != nil {
		log.Error(err)
		return
	}

	matchUpdate := packets.Match(26, m)
	m.Write(matchUpdate)
	lobby.Write(matchUpdate)
	log.Info(p, "updated settings for", m)
}
