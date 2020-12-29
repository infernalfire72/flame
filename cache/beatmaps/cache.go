package beatmaps

import (
	"errors"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/osuapi"
	"gorm.io/gorm"
	"sync"
)

var values = map[string]*layouts.Beatmap{}
var mutex sync.RWMutex

func Get(hash string) *layouts.Beatmap {
	mutex.RLock()
	v, ok := values[hash]
	mutex.RUnlock()

	if ok {
		return v
	} else {
		return Fetch(hash)
	}
}

func Fetch(hash string) *layouts.Beatmap {
	b := &layouts.Beatmap{Hash: hash}
	if err := database.DB.Where(b).First(b).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
		}

		res, err := osuapi.GetBeatmaps(osuapi.GetBeatmapsOpts{
			BeatmapHash: hash,
		})

		if err != nil {
			log.Error(err)
			return nil
		}

		if len(res) == 0 {
			return nil
		}

		set, err := osuapi.GetBeatmaps(osuapi.GetBeatmapsOpts{
			BeatmapSetID: res[0].BeatmapSetID,
		})


		for _, m := range set {
			bm := &layouts.Beatmap{}
			layouts.BeatmapFromApiModel(bm, &m)

			mutex.Lock()
			if _, ok := values[bm.Hash]; !ok {
				values[bm.Hash] = bm
			}
			mutex.Unlock()

			database.DB.Save(bm)
			if m.FileMD5 == hash {
				b = bm
			}
		}
	}

	return b
}