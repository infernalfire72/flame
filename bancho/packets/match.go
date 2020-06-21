package packets

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/objects"
)

func Match(id int16, m *objects.MultiplayerLobby) Packet {
	s := io.NewStreamWithCapacity(384)
	s.WriteInt16(id)
	s.WriteByte(0)
	s.WriteInt32(0)

	

	s.Position = 3
	s.WriteInt32(int32(s.Length - 7))

	return s.Data()
}