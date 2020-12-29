package layouts

type Stats struct {
	User        int   `gorm:"primaryKey"`
	RankedScore int64
	TotalScore  int64

	Accuracy `gorm:"embedded"`
	Playtime       int
	Playcount      int
	ReplaysWatched int
	Performance    float32
	Rank           int
}

func (s *Stats) TotalHits() int {
	return int(s.N300) + int(s.N100) + int(s.N50)
}