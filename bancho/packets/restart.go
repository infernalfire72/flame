package packets

func Restart(ms int) Packet {
	return MakePacket(86, ms)
}
