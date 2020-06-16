package packets

import "github.com/infernalfire72/flame/objects"

func AvailableChannel(c *objects.Channel) Packet {
	return MakePacket(65, c.Name, c.Topic, c.UserCount())
}