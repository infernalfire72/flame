package packets

import "github.com/infernalfire72/flame/io"

func FriendsList(friends []int) Packet {
	p := io.NewStreamWithCapacity(2 + len(friends) * 4)

	p.WriteInt16(int16(len(friends)))

	for _, v := range friends {
		p.WriteInt32(int32(v))
	}

	return p.Content
}