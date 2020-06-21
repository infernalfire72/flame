package packets

func NewSpectator(id int32) Packet {
	return MakePacket(13, id)
}

func NewFellowSpectator(id int32) Packet {
	return MakePacket(42, id)
}