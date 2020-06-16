package packets

import "github.com/infernalfire72/flame/objects"

func AutojoinChannel(c *objects.Channel) Packet {
	return MakePacket(67, c.Name, c.Topic, c.UserCount())
}