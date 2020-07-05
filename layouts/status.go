package layouts

import "github.com/infernalfire72/flame/io"

type Status struct {
	Action      byte
	InfoText    string
	BeatmapHash string
	Mods        int32
	Gamemode    byte
	Beatmap     int32
}

func ReadStatus(s *Status, bytes []byte) error {
	var err error

	data := io.Stream{bytes, 0}
	s.Action, err = data.ReadByte()
	if err != nil {
		return err
	}

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

	s.Gamemode, err = data.ReadByte()
	if err != nil {
		return err
	}

	s.Beatmap, err = data.ReadInt32()

	return err
}
