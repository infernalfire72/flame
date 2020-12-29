package models

import "github.com/infernalfire72/flame/io"

type Status struct {
	Action      byte   `json:"action"`
	InfoText    string `json:"info_text"`
	BeatmapHash string `json:"beatmap_hash"`
	Mods        int32  `json:"mods"`
	Gamemode    byte   `json:"mode"`
	Beatmap     int32  `json:"beatmap_id"`
}

func (s *Status) Unmarshal(data []byte) error {
	var err error

	stream := io.Stream{data, 0}
	s.Action, err = stream.ReadByte()
	if err != nil {
		return err
	}

	s.InfoText, err = stream.ReadString()
	if err != nil {
		return err
	}

	s.BeatmapHash, err = stream.ReadString()
	if err != nil {
		return err
	}

	s.Mods, err = stream.ReadInt32()
	if err != nil {
		return err
	}

	s.Gamemode, err = stream.ReadByte()
	if err != nil {
		return err
	}

	s.Beatmap, err = stream.ReadInt32()

	return err
}