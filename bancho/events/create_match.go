package events

import (
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/matches"
	"github.com/infernalfire72/flame/bancho/packets"
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
	m.Host = p.ID
	m.Creator = p.ID
	log.Info(p.Username, "created a new MultiplayerLobby", m.Name)

	m.AddPlayer(p, m.Password)

	packet := packets.Match(26, m)

	lobby.Mutex.RLock()
	for _, t := range lobby.Players {
		t.Write(packet)
	}
	lobby.Mutex.RUnlock()

	p.Write(packet)
	packet[0] = 36
	p.Write(packet, packets.JoinedChannel("#multiplayer"), packets.AvailableChannelArgs("#multiplayer", "Multiplayer Channel", 1))
}