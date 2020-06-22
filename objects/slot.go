package objects

import "github.com/infernalfire72/flame/constants"

type MultiplayerSlot struct {
	Status	constants.SlotStatus
	Team	byte
	User	*Player
	Mods	int32

	Loaded, Skipped, Completed bool
}

func (s *MultiplayerSlot) Clear() {
	s.Status = constants.SlotEmpty
	s.User = nil
	s.Team = 0
	s.Mods = 0
}