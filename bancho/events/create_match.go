package events

import (
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/matches"
)

func CreateMatch(p *objects.Player, bytes []byte) {
	id := matches.GetNextID()

	if id == 0 {
		return
	}

	m := &objects.MultiplayerLobby{
		ID:	id,
	}

	m.ReadMatch(bytes)
}