package objects

import (
	"fmt"
	"sync"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/io"
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

func (m *MultiplayerLobby) FindFreeSlot() *MultiplayerSlot {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	for i := 0; i < 16; i++ {
		if !m.Slots[i].Status.HasFlag(constants.SlotOccupied) {
			return &m.Slots[i]
		}
	}

	return nil
}

func (m *MultiplayerLobby) FindPlayerSlot(p *Player) *MultiplayerSlot {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	for i := 0; i < 16; i++ {
		if m.Slots[i].User == p {
			return &m.Slots[i]
		}
	}
	return nil
}

func (m *MultiplayerLobby) AddPlayer(p *Player, password string) bool {
	if password != m.Password {
		return false
	}

	m.Mutex.Lock()
	m.Players = append(m.Players, p)
	m.Mutex.Unlock()

	if slot := m.FindFreeSlot(); slot != nil {
		slot.Status = constants.SlotNotReady
		slot.User = p
		p.Match = m
		return true
	}

	return false
}

// TODO: this
func (m *MultiplayerLobby) RemovePlayer(p *Player) {
	if slot := m.FindPlayerSlot(p); slot != nil {
		p.Match = nil
		slot.Clear()
	}

	m.Mutex.Lock()
	for i, t := range m.Players {
		if t == p {
			m.Players[i] = m.Players[len(m.Players)-1]
			m.Players[len(m.Players)-1] = nil
			m.Players = m.Players[:len(m.Players)-1]
			break
		}
	}
	m.Mutex.Unlock()
}

func (m *MultiplayerLobby) UserCount() int16 {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()

	return int16(len(m.Players))
}

func (m *MultiplayerLobby) Write(data ...[]byte) {
	m.Mutex.RLock()

	for _, t := range m.Players {
		t.Write(data...)
	}

	m.Mutex.RUnlock()
}

func (m *MultiplayerLobby) ReadMatch(bytes []byte) error {
	s := &io.Stream{bytes, len(bytes), len(bytes), 2}
	var err error

	m.Running = s.ReadBoolean()
	m.Type = constants.MatchType(s.ReadByte())
	m.Mods, err = s.ReadInt32()
	if err != nil {
		return err
	}

	m.Name, err = s.ReadString()
	if err != nil {
		return err
	}

	m.Password, err = s.ReadString()
	if err != nil {
		return err
	}

	m.BeatmapName, err = s.ReadString()
	if err != nil {
		return err
	}

	m.BeatmapID, err = s.ReadInt32()
	if err != nil {
		return err
	}

	m.BeatmapHash, err = s.ReadString()
	if err != nil {
		return err
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
	return err
}

func (m MultiplayerLobby) String() string {
	return fmt.Sprintf("%d (%s)", m.ID, m.Name)
}