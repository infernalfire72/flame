package packets

func ChannelRevoked(name string) Packet {
	return MakePacket(66, name)
}
