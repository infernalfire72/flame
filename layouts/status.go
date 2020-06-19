package layouts

import "github.com/infernalfire72/flame/io"

type Status struct {
	Action		byte
	InfoText	string
	BeatmapHash	string
	Mods		int32
	Gamemode	byte
	Beatmap		int32
}

func ReadStatus(s *Status, bytes []byte) error {
	var err error

	data := io.Stream{bytes, len(bytes), len(bytes), 0}
	s.Action = byte(data.ReadByte())
	s.InfoText, err = data.ReadString()
	if err != nil {
		return err
	}

	s.BeatmapHash, err = data.ReadString()
	if err != nil {
		return err
	}

	s.Mods, err = data.ReadInt32()
	if err != nil {
		return err
	}
	
	s.Gamemode = byte(data.ReadByte())

	s.Beatmap, err = data.ReadInt32()

	return err
}