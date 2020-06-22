package layouts

type ModeData struct {
	TotalScore  int64   `db:"total_score" json:"total_score"`
	RankedScore int64   `db:"ranked_score" json:"ranked_score"`
	Performance int32   `db:"pp" json:"pp"`
	Playcount   int32   `db:"playcount" json:"playcount"`
	Accuracy    float32 `db:"accuracy" json:"accuracy"`
	Rank        int32   `json:"rank"`
	MaxCombo    int16   `json:"max_combo"`
}
