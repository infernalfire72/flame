package packets

import (
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/objects"
)

func Match(id int16, m *objects.MultiplayerLobby) Packet {
	s := io.NewStreamWithCapacity(384)
	s.WriteInt16(id)
	s.WriteByte(0)
	s.WriteInt32(0)

	s.WriteInt16(int16(m.ID))
	s.WriteBoolean(m.Running)
	s.WriteByte(byte(m.Type))
	s.WriteInt32(m.Mods)

	s.WriteString(m.Name)
	s.WriteString(m.Password)
	s.WriteString(m.BeatmapName)
	s.WriteInt32(m.BeatmapID)
	s.WriteString(m.BeatmapHash)

	m.Mutex.RLock()
	for _, slot := range m.Slots {
		s.WriteByte(byte(slot.Status))
	}

	for _, slot := range m.Slots {
		s.WriteByte(byte(slot.Team))
	}

	for _, slot := range m.Slots {
		if slot.User != nil {
			s.WriteInt32(int32(slot.User.ID))
		}
	}

	m.Mutex.RUnlock()

	s.WriteInt32(int32(m.Host))
	s.WriteByte(m.Gamemode)
	s.WriteByte(byte(m.ScoringType))
	s.WriteByte(byte(m.TeamType))

	s.WriteBoolean(m.FreeMod)

	if m.FreeMod {
		m.Mutex.RLock()
		for _, slot := range m.Slots {
			s.WriteInt32(slot.Mods)
		}
		m.Mutex.RUnlock()
	}

	s.WriteInt32(m.ManiaSeed)

	s.Position = 3
	s.WriteInt32(int32(s.Length - 7))

	return s.Data()
}