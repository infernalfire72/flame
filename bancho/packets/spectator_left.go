package packets

func SpectatorLeft(id int32) Packet {
	return MakePacket(14, id)
}

func FellowSpectatorLeft(id int32) Packet {
	return MakePacket(43, id)
}
