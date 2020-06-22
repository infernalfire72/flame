package constants

type SlotStatus byte

const (
	SlotEmpty SlotStatus = 1 << iota
	SlotLocked
	SlotNotReady
	SlotReady
	SlotMissingBeatmap
	SlotPlaying
	SlotQuit

	SlotOccupied = SlotNotReady | SlotReady | SlotMissingBeatmap | SlotPlaying
)

func (a SlotStatus) HasFlag(b SlotStatus) bool {
	return b == 0 || (a&b) != 0
}
