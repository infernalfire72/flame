package events

import (
	"strings"

	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/bot/commands"
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

	p.AwaiterMutex.RLock()
	if p.MessageAwaiter != nil {
		p.MessageAwaiter <- m.Content
	}
	p.AwaiterMutex.RUnlock()

	if strings.HasPrefix(m.Content, commands.Prefix) {
		go commands.Execute(p, m.Content, target)
	}

	target.AddMessage(p, packets.IrcMessage(m))
	log.Chat(p.String(), m.Target, m.Content)
}
