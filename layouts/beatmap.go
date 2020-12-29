package layouts

import (
	"fmt"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/osuapi"
)

type Beatmap struct {
	ID      int    `gorm:"primaryKey"`
	SetID   int
	Hash    string `gorm:"unique"`

	Artist  string
	Title   string
	Mapper  string
	Version string

	RankedStatus constants.BeatmapStatus
	StatusFrozen bool
}

func (b *Beatmap) Online(n int) string {
	return fmt.Sprintf("%d|false|%d|%d|%d\n0\n%s\n10.0\n", b.RankedStatus, b.ID, b.SetID, n, b.SongName())
}

func (b *Beatmap) SongName() string {
	return b.Artist + " - " + b.Title + " [" + b.Version + "]"
}

func BeatmapFromApiModel(b *Beatmap, beatmap *osuapi.Beatmap) {
	b.ID = beatmap.BeatmapID
	b.SetID = beatmap.BeatmapSetID
	b.Hash = beatmap.FileMD5
	b.Artist = beatmap.Artist
	b.Title = beatmap.Title
	b.Mapper = beatmap.Creator
	b.Version = beatmap.DiffName

	switch s := beatmap.Approved; s {
	case osuapi.StatusApproved:
		b.RankedStatus = constants.StatusRanked
	case osuapi.StatusRanked, osuapi.StatusLoved, osuapi.StatusQualified:
		b.RankedStatus = constants.BeatmapStatus(s + 1)
	default:
		b.RankedStatus = constants.StatusPending
	}

	b.StatusFrozen = true
}