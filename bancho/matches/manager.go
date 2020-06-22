package matches

import (
	"math"
	"sync"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/objects"
)

var (
	Values map[uint16]*objects.MultiplayerLobby
	Mutex  sync.RWMutex
)

func New() *objects.MultiplayerLobby {
	m := &objects.MultiplayerLobby{
		ID: GetNextID(),
	}

	for i := 0; i < 16; i++ {
		m.Slots[i].Status = constants.SlotEmpty
	}

	return m
}

func Init() {
	Values = make(map[uint16]*objects.MultiplayerLobby)
}

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

func Disband(m *objects.MultiplayerLobby) {
	Mutex.Lock()
	delete(Values, m.ID)
	Mutex.Unlock()
}
