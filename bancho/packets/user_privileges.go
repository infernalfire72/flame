package packets

func UserPrivileges(privileges int) Packet {
	return MakePacket(71, privileges)
}
