package packets

func Alert(msg string) Packet {
	return MakePacket(24, msg)
}