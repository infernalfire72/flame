package layouts

import "github.com/infernalfire72/flame/io"

type Message struct {
	Username string
	Content  string
	Target   string
	UserID   int32
}

func ReadMessage(bytes []byte, m *Message) (err error) {
	data := io.Stream{bytes, 0}

	m.Username, err = data.ReadString()
	if err != nil {
		return
	}

	m.Content, err = data.ReadString()
	if err != nil {
		return
	}

	m.Target, err = data.ReadString()
	if err != nil {
		return
	}

	m.UserID, err = data.ReadInt32()

	return
}
