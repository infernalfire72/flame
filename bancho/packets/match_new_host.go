package packets

func MatchNewHost(index int) Packet {
	return MakePacket(50, index)
}
