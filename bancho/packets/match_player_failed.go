package packets

func MatchPlayerFailed(index int) Packet {
	return MakePacket(57, index)
}
