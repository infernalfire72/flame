package events

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/matches"
	"github.com/infernalfire72/flame/bancho/packets"
)

func JoinMatch(p *objects.Player, bytes []byte) {
	if p.Match != nil {
		LeaveMatch(p)
	}

	s := io.Stream{bytes, 0}

	id, err := s.ReadUint16()
	if err != nil {
		log.Error(err)
		return
	}

	m := matches.Get(id)
	if password, err := s.ReadString(); err != nil || m == nil || !m.AddPlayer(p, password) {
		p.Write(packets.MatchRevoked())
		return
	}
	lobby.RemovePlayer(p)

	matchInfo := packets.Match(36, m)
	channelInfo := packets.AvailableChannelArgs("#multiplayer", "Multiplayer Channel", m.UserCount())
	p.Write(matchInfo)

	matchInfo.ChangeID(26) // just change the packet id 3head
	m.Write(matchInfo, channelInfo)
	lobby.Write(matchInfo)

	p.Write(packets.JoinedChannel("#multiplayer"))
	log.Info(p, "joined", m)
}
