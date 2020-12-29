package models

import "github.com/infernalfire72/flame/io"

type Message struct {
	Username string
	Content  string
	Target   string
	Sender   int32
}

func (m *Message) Unmarshal(data []byte) error {
	var err error

	stream := io.Stream{data, 0}
	m.Username, err = stream.ReadString()
	if err != nil {
		return err
	}

	m.Content, err = stream.ReadString()
	if err != nil {
		return err
	}

	m.Target, err = stream.ReadString()
	if err != nil {
		return err
	}

	m.Sender, err = stream.ReadInt32()

	return err
}