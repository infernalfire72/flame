package events

import (
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func PrivateMessage(p *objects.Player, bytes []byte) {
	var m layouts.Message

	err := layouts.ReadMessage(bytes, &m)
	if err != nil {
		log.Error(err)
		return
	}

	if target := players.FindPlayerName(m.Target); target != nil {
		m.Username = p.Username
		target.Write(packets.IrcMessage(m))
	}
}