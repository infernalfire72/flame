package beatmaps

import (
	"fmt"
	"time"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
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

func (b *Beatmap) String() string {
	return fmt.Sprintf("%d|false", b.Status)
}

func (b *Beatmap) StringC(scoresCount int) string {
	return fmt.Sprintf("%d|false|%d|%d|%d\n0\n%s\n10.0\n", b.Status, b.ID, b.SetID, scoresCount, b.Name)
}
