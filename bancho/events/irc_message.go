package events

import (
	"github.com/infernalfire72/flame/bancho/models"

	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/channels"
	"github.com/infernalfire72/flame/bancho/packets"
)

func IrcMessage(p *objects.Player, bytes []byte) {
	m := models.Message{}
	err := m.Unmarshal(bytes)
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

	// TODO: Commands
	/*if strings.HasPrefix(m.Content, commands.Prefix) {
		go commands.Execute(p, m.Content, target)
	}*/

	target.AddMessage(p, packets.IrcMessage(m))
	log.Chat(p.String(), m.Target, m.Content)
}
