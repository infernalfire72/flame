package packets

import "github.com/infernalfire72/flame/io"

func FriendsList(friends []int) Packet {
	p := io.NewStreamWithCapacity(9 + len(friends)*4)

	p.WriteInt16(72)
	p.WriteByte(0)
	p.WriteInt32(int32(2 + len(friends)*4))

	p.WriteInt16(int16(len(friends)))

	for _, v := range friends {
		p.WriteInt32(int32(v))
	}

	return p.Content
}
