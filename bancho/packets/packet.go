package packets

import (
	"github.com/infernalfire72/flame/io"
)

type Packet []byte

type PacketType int16

func MakePacket(id PacketType, values ...interface{}) Packet {
	s := io.NewStreamWithCapacity(768)
	s.WriteInt16(int16(id))
	s.WriteByte(0)
	s.WriteInt32(0)

	for _, v := range values {
		s.WriteInterface(v)
	}

	s.Position = 3
	s.WriteInt32(int32(len(s.Content) - 7))

	return s.Content
}

func (p Packet) ChangeID(newID int16) {
	p[0] = byte(newID)
	p[1] = byte(newID >> 8)
}
