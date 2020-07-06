package packets

import "github.com/infernalfire72/flame/layouts"

func IrcMessage(m layouts.Message) Packet {
	return MakePacket(7, m.Username, m.Content, m.Target, m.UserID)
}

func IrcMessageArgs(username, content, target string, id int) Packet {
	return MakePacket(7, username, content, target, id)
}