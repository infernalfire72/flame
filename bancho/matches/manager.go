package matches

import (
	"math"
	"sync"

	"github.com/infernalfire72/flame/objects"
)

var (
	Values	map[uint16]*objects.MultiplayerLobby
	Mutex	sync.RWMutex
)

func Get(id uint16) *objects.MultiplayerLobby {
	Mutex.RLock()
	defer Mutex.RUnlock()
	return Values[id]
}

func GetNextID() uint16 {
	Mutex.RLock()
	defer Mutex.RUnlock()
	for i := uint16(1); i < math.MaxUint16; i++ {
		if Values[i] == nil {
			return i
		}
	}

	return 0
}