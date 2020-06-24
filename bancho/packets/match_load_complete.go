package packets

func MatchLoadComplete() Packet {
	return MakePacket(53)
}
