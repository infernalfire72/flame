package objects

type ModeData struct {
	TotalScore	int64	`json:"total_score"`
	RankedScore	int64	`json:"ranked_score"`
	Performance	int32	`json:"pp"`
	Playcount	int32	`json:"playcount"`
	Rank		int32	`json:"rank"`
	Accuracy	float32	`json:"accuracy"`
	MaxCombo	int16	`json:"max_combo"`
}