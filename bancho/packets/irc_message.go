package packets

import "github.com/infernalfire72/flame/bancho/models"

func IrcMessage(m models.Message) Packet {
	return MakePacket(7, m.Username, m.Content, m.Target, m.Sender)
}

func IrcMessageArgs(username, content, target string, id int) Packet {
	return MakePacket(7, username, content, target, id)
}
