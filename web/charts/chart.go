package charts

import "fmt"

type BaseChart struct {
	BeatmapID        int
	BeatmapSetID     int
	BeatmapPlaycount int
	BeatmapPasscount int
	ApprovedDate     time
}

type Chart struct {
	ID   string
	URL  string
	Name string

	Rank        [2]int
	MaxCombo    [2]int
	Accuracy    [2]float32
	RankedScore [2]int
	TotalScore  [2]int
	Performance [2]int
	ScoreID     int
}

// TODO: This is horrible
func (c Chart) String() string {
	return fmt.Sprintf("chartId:%s|chartUrl:%s|chartName:%s|%s|%s|%s|%s|%s|%s|onlineScoreId: %d", c.ID, c.URL, c.Name,
		beforeAfterString("rank", c.Rank[0], c.Rank[1]),
		beforeAfterString("maxCombo", c.MaxCombo[0], c.MaxCombo[1]),
		beforeAfterString("accuracy", c.Accuracy[0], c.Accuracy[1]),
		beforeAfterString("rankedScore", c.RankedScore[0], c.RankedScore[1]),
		beforeAfterString("totalScore", c.TotalScore[0], c.TotalScore[1]),
		beforeAfterString("pp", c.Performance[0], c.Performance[1]),
		c.ScoreID,
	)
}

func beforeAfterString(attr string, before, after interface{}) string {
	return fmt.Sprintf("%[1]sBefore:%v|%[1]sAfter:%v", attr, before, after)
}
