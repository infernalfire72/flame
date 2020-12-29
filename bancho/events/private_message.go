package events

import (
	"github.com/infernalfire72/flame/bancho/models"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/bot"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func PrivateMessage(p *objects.Player, bytes []byte) {
	m := models.Message{}
	err := m.Unmarshal(bytes)
	if err != nil {
		log.Error(err)
		return
	}

	p.AwaiterMutex.RLock()
	if p.MessageAwaiter != nil {
		p.MessageAwaiter <- m.Content
	}
	p.AwaiterMutex.RUnlock()

	if bot.Player != nil && m.Target == bot.Player.Username {
		// go commands.Execute(p, m.Content, p)
		log.Chat(p.String(), bot.Player.String(), m.Content)
		return
	}

	if target := players.FindUsername(m.Target); target != nil {
		m.Username = p.Username
		target.Write(packets.IrcMessage(m))
		log.Chat(p.String(), target.String(), m.Content)
	}
}
