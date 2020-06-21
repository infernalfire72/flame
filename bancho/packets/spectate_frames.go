package packets

import "github.com/infernalfire72/flame/io"

func SpectateFrames(bytes []byte) Packet {
	s := io.NewStreamWithCapacity(7 + len(bytes))
	s.WriteInt16(15)
	s.WriteByte(0)
	s.WriteInt32(int32(len(bytes)))
	s.WriteByteSlice(bytes)

	return s.Content
}