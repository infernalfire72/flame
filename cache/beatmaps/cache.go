package beatmaps

import (
	"sync"
	"time"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/log"
)

var (
	Mutex  sync.RWMutex
	Values map[string]*Beatmap
)

func init() {
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

	err := b.FetchFromDb()
	if err != nil {
		log.Error(err)
		return nil
	}

	Mutex.Lock()
	defer Mutex.Unlock()

	Values[md5] = b
	return b
}

func FetchFromApi(md5, filename string) *Beatmap {
	b := &Beatmap{
		Md5: md5,
	}

	err := b.FetchFromApi(filename)
	if err != nil {
		log.Error(err)
		return nil
	}

	if b.Status >= constants.StatusRanked {
		b.SetToDb()

		Mutex.Lock()
		Values[md5] = b
		Mutex.Unlock()
	}

	return b
}
