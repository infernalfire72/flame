package packets

// Efficiency tip: pass a channel pointer and access its field instead (4/8 bytes allocated and passed vs. string struct allocated copied and passed 12/24 bytes)
// Tip ignored for use
func JoinedChannel(name string) Packet {
	return MakePacket(64, name)
}
