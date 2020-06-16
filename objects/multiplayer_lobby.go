package objects

import (
	"math"
)

var Matches map[uint16]*MultiplayerLobby

func GetNewMatchID() uint16 {
	for i := uint16(1); i < math.MaxUint16; i++ {
		if Matches[i] != nil {
			return i
		}
	}

	return 0
}

type MultiplayerLobby struct {
	ID		uint16		`json:"id"`
}