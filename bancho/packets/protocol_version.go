package packets

func ProtocolVersion(version int) Packet {
	return MakePacket(75, version)
}