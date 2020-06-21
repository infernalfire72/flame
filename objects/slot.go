package objects

import "github.com/infernalfire72/flame/constants"

type MultiplayerSlot struct {
	Status	constants.SlotStatus
	Team	byte
	User	*Player
	Mods	int32

	Loaded, Skipped, Completed bool
}