package packets

func MatchFinished() Packet {
	return MakePacket(58)
}
