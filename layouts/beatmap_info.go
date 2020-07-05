package layouts

import "github.com/infernalfire72/flame/io"

type BeatmapInfoRequest struct {
	Names []string
}

func ReadBeatmapInfoRequest(b *BeatmapInfoRequest, bytes []byte) (err error) {
	s := io.Stream{bytes, 0}

	count, err := s.ReadInt32()
	if err != nil {
		return
	}

	b.Names = make([]string, count)

	for i := int32(0); i < count; i++ {
		b.Names[i], err = s.ReadString()
		if err != nil {
			return
		}
	}

	return nil
}

type BeatmapInfo struct {
	ID     int
	SetID  int
	Status byte
	Ranks  [4]byte
	Hash   string
}
