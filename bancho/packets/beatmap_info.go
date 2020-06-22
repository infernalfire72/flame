package packets

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/layouts"
)

// 2 + 4 + 4 + 4 + 1 + 4 + 2 + 32 = 53
func BeatmapInfo(beatmapInfo []*layouts.BeatmapInfo) Packet {
	s := io.NewStreamWithCapacity(53 * len(beatmapInfo))

	for i, info := range beatmapInfo {
		s.WriteInt16(int16(i))
		s.WriteInt32(int32(info.ID))
		s.WriteInt32(int32(info.SetID))
		s.WriteInt32(0) // ???
		s.WriteByte(info.Status)
		s.WriteInt32(151587081) // or *(*int32)(unsafe.Pointer(&info.Ranks))
		s.WriteString(info.Hash)
	}

	return s.Content
}
