package beatmaps

import (
	"sync"
	"time"
)

var (
	Mutex  sync.RWMutex
	Values map[string]*Beatmap
)

func Init() {
	Mutex.Lock()
	Values = make(map[string]*Beatmap)
	Mutex.Unlock()
}

func Get(md5 string) *Beatmap {
	Mutex.RLock()
	if v, ok := Values[md5]; ok {
		Mutex.RUnlock()

		now := time.Now()
		if now.Sub(v.LastUpdate).Seconds() > 30 {
			v.FetchFromDb()
		}

		return v
	}
	Mutex.RUnlock()
	return FetchFromDb(md5)
}

func FetchFromDb(md5 string) *Beatmap {
	b := &Beatmap{
		Md5: md5,
	}

	b.FetchFromDb()
	Mutex.Lock()
	defer Mutex.Unlock()

	Values[md5] = b
	return b
}
