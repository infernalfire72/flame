package beatmaps

import (
	"fmt"
	"time"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/osuapi"
)

type Beatmap struct {
	Md5        string
	ID         int
	SetID      int
	Name       string
	Status     constants.BeatmapStatus
	Playcount  int
	Passcount  int
	LastUpdate time.Time
}

func (b *Beatmap) FetchFromDb() error {
	err := config.Database.QueryRow("SELECT beatmap_id, beatmapset_id, song_name, ranked, playcount, passcount FROM beatmaps WHERE beatmap_md5 = ?", b.Md5).Scan(
		&b.ID, &b.SetID, &b.Name, &b.Status, &b.Playcount, &b.Passcount,
	)
	if err != nil {
		return err
	}

	b.LastUpdate = time.Now()
	return nil
}

func (b *Beatmap) SetToDb() error {
	_, err := config.Database.Exec("REPLACE INTO beatmaps (beatmap_id, beatmapset_id, beatmap_md5, song_name, ranked, latest_update) VALUES (?, ?, ?, ?, ?, UNIX_TIMESTAMP())")
	return err
}

func (b *Beatmap) FetchFromApi(filename string) error {
	res, err := osuapi.GetBeatmaps(osuapi.GetBeatmapsOpts{
		BeatmapHash: b.Md5,
	})
	if err != nil {
		return err
	}

	if len(res) == 0 {
		file := osuapi.GetBeatmapContent(filename)
		if len(file) > 25 {
			b.Status = constants.StatusNeedUpdate
		} else {
			b.Status = constants.StatusNotSubmitted
		}
	} else {
		fmt.Println(res)
	}

	return nil
}

func (b *Beatmap) Online() string {
	return fmt.Sprintf("%d|false", b.Status)
}

func (b *Beatmap) OnlineRanked(scoresCount int) string {
	return fmt.Sprintf("%d|false|%d|%d|%d\n0\n%s\n10.0\n", b.Status, b.ID, b.SetID, scoresCount, b.Name)
}
