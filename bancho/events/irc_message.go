package events

import (
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
)

func IrcMessage(p *objects.Player, bytes []byte) {
	var m layouts.Message

	err := layouts.ReadMessage(bytes, &m)
	if err != nil {
		log.Error(err)
		return
	}

	target := channels.Get(m.Target)
	if target == nil {
		log.Info(p.Username, "tried to write to non-existent channel", m.Target)
		return
	}

	target.AddMessage(p, packets.IrcMessage(m))
	log.Chat(m.Username, m.Target, m.Content)
}