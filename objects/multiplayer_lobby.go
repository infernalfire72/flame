package objects

import (
	"sync"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/log"
)

type MultiplayerLobby struct {
	ID			uint16		`json:"id"`
	Name		string
	Password	string		`json:"-"`

	BeatmapName	string
	BeatmapID	int32
	BeatmapHash	string

	Running		bool
	Type		constants.MatchType
	ScoringType	constants.MatchScoringType
	TeamType	constants.MatchTeamType

	Gamemode	byte
	Mods		int32
	FreeMod		bool
	ManiaSeed	int32

	Host		int
	Creator		int

	Mutex		sync.RWMutex
	Players		[]*Player

	Slots		[16]MultiplayerSlot
}

func (m *MultiplayerLobby) ReadMatch(bytes []byte) {
	s := &io.Stream{bytes, len(bytes), len(bytes), 2}
	var err error

	m.Running = s.ReadBoolean()
	m.Type = constants.MatchType(s.ReadByte())
	m.Mods, err = s.ReadInt32()
	if err != nil {
		log.Error(err)
		return
	}

	m.Name, err = s.ReadString()
	if err != nil {
		log.Error(err)
		return
	}

	m.Password, err = s.ReadString()
	if err != nil {
		log.Error(err)
		return
	}

	m.BeatmapName, err = s.ReadString()
	if err != nil {
		log.Error(err)
		return
	}

	m.BeatmapID, err = s.ReadInt32()
	if err != nil {
		log.Error(err)
		return
	}

	m.BeatmapHash, err = s.ReadString()
	if err != nil {
		log.Error(err)
		return
	}

	s.Position += 32

	for _, slot := range m.Slots {
		if slot.Status.HasFlag(constants.SlotOccupied) {
			s.Position += 4
		}
	}

	s.Position += 4

	m.Gamemode = byte(s.ReadByte())
	m.ScoringType = constants.MatchScoringType(s.ReadByte())
	m.TeamType = constants.MatchTeamType(s.ReadByte())

	freeMod := s.ReadBoolean()

	if freeMod != m.FreeMod {
		if freeMod {
			for _, slot := range m.Slots {
				if slot.User != nil {
					slot.Mods = m.Mods
				}
			}
		} else {
			for _, slot := range m.Slots {
				if slot.User != nil {
					slot.Mods = 0
				}
			}
		}

		m.Mods = 0
		m.FreeMod = freeMod
	}

	m.ManiaSeed, err = s.ReadInt32()
	if err != nil {
		log.Error(err)
		return
	}
}