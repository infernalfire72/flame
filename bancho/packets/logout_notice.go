package packets

func LogoutNotice(id, v int) Packet {
	return MakePacket(12, id, v)
}
