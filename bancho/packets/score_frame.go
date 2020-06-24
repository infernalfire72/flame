package packets

func ScoreFrame(bytes []byte) Packet {
	return MakePacket(48, bytes)
}
