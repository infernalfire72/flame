package packets

func LoginReply(id int) Packet {
	return MakePacket(5, id)
}
