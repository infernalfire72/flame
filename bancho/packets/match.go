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
	s.WriteByte(m.Type)
	s.WriteInt32(int32(m.Mods))

	s.WriteString(m.Name)
	s.WriteString(m.Password)
	s.WriteString(m.BeatmapName)
	s.WriteInt32(m.BeatmapID)
	s.WriteString(m.BeatmapHash)

	m.ForSlots(func(slot *objects.MultiplayerSlot) {
		s.WriteByte(byte(slot.Status))
	})

	m.ForSlots(func(slot *objects.MultiplayerSlot) {
		s.WriteByte(byte(slot.Team))
	})

	m.ForSlots(func(slot *objects.MultiplayerSlot) {
		if slot.User != nil {
			s.WriteInt32(int32(slot.User.ID))
		}
	})

	s.WriteInt32(int32(m.Host))
	s.WriteByte(m.Gamemode)
	s.WriteByte(byte(m.ScoringType))
	s.WriteByte(byte(m.TeamType))

	s.WriteBoolean(m.FreeMod)

	if m.FreeMod {
		m.ForSlots(func(slot *objects.MultiplayerSlot) {
			s.WriteInt32(int32(slot.Mods))
		})
	}

	s.WriteInt32(m.ManiaSeed)

	s.Position = 3
	s.WriteInt32(int32(s.Length - 7))

	return s.Data()
}
