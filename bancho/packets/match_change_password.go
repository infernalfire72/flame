package packets

func MatchChangePassword(newPassword string) Packet {
	return MakePacket(91, newPassword)
}
