package events

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/lobby"
	"github.com/infernalfire72/flame/bancho/packets"
)

func MatchChangePassword(p *objects.Player, bytes []byte) {
	if p.Match == nil || p.Match.Host != p.ID || len(bytes) < 12 {
		return
	}

	s := &io.Stream{bytes, len(bytes), len(bytes), 8}
	_, err := s.ReadString()
	if err != nil {
		return
	}

	newPassword, err := s.ReadString()
	if err != nil {
		return
	}
	m := p.Match

	m.Password = newPassword
	p.Write(packets.MatchChangePassword(newPassword))

	matchInfo := packets.Match(26, m)
	m.Write(matchInfo)
	lobby.Write(matchInfo)
}
