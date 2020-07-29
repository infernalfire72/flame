package packets

func Pong() Packet {
	return MakePacket(8)
}
