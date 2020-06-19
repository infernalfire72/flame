package objects

import (
	"sync"
)

type PlayerCollection struct {
	Players	map[string]*Player
	Mutex	sync.RWMutex
}

