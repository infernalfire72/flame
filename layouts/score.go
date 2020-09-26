package layouts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/infernalfire72/flame/constants"
)

type Score struct {
	ID          int    `db:"id"`
	BeatmapHash string `db:"beatmap_md5"`

	UserID      int    `db:"userid"`
	ScoreHash   string

	N300  int        `db:"300_count"`
	N100  int        `db:"100_count"`
	N50   int        `db:"50_count"`
	NGeki int        `db:"gekis_count"`
	NKatu int        `db:"katus_count"`
	NMiss int        `db:"misses_count"`
	Accuracy float32 `db:"accuracy"`

	TotalScore int    `db:"score"`
	Combo      int    `db:"max_combo"`
	FullCombo  bool   `db:"full_combo"`
	Rank       string `db:"rank"`

	Mods   constants.Mod `db:"mods"`
	Passed bool
	Mode   constants.Mode `db:"play_mode"`

	Timestamp int `db:"time"`
	Version   int

	Flags int

	Relax       bool
	Performance float32               `db:"pp"`
	Status      constants.ScoreStatus `db:"completed"`
}

func (s *Score) AddToDatabase(db *sqlx.DB) (err error) {
	table := "scores"
	if s.Relax {
		table = "scores_relax"
	}

	_, err = db.NamedExec("INSERT INTO "+table+" VALUES (NULL, :beatmap_md5, :userid , :score, :max_combo, :full_combo, :mods, :300_count, :100_count, :50_count, :katus_count, :gekis_count, :misses_count, :time, :play_mode, :completed, :accuracy, :pp)", s)
	return
}

func (s *Score) Online(displayScore bool, displayName string, pos int) string {
	lbScore := int(s.Performance)
	if displayScore {
		lbScore = s.TotalScore
	}

	fc := "False"
	if s.FullCombo {
		fc = "True"
	}

	return fmt.Sprintf("%d|%s|%d|%d|%d|%d|%d|%d|%d|%d|%s|%d|%d|%d|%d|1\n", s.ID, displayName, lbScore, s.Combo, s.N50, s.N100, s.N300, s.NMiss, s.NKatu, s.NGeki, fc, s.Mods, s.UserID, pos, s.Timestamp)
}

func ReadScoreLayout(parts []string, out *Score) (err error) {
	out.BeatmapHash = parts[0]
	out.ScoreHash = parts[2]

	out.N300, err = strconv.Atoi(parts[3])
	if err != nil {
		return
	}

	out.N100, err = strconv.Atoi(parts[4])
	if err != nil {
		return
	}

	out.N50, err = strconv.Atoi(parts[5])
	if err != nil {
		return
	}

	out.NGeki, err = strconv.Atoi(parts[6])
	if err != nil {
		return
	}

	out.NKatu, err = strconv.Atoi(parts[7])
	if err != nil {
		return
	}

	out.NMiss, err = strconv.Atoi(parts[8])
	if err != nil {
		return
	}

	out.TotalScore, err = strconv.Atoi(parts[9])
	if err != nil {
		return
	}

	out.Combo, err = strconv.Atoi(parts[10])
	if err != nil {
		return
	}

	out.FullCombo, err = strconv.ParseBool(parts[11])
	if err != nil {
		return
	}

	out.Rank = parts[12]

	mods, err := strconv.Atoi(parts[13])
	if err != nil {
		return
	}
	out.Mods = constants.Mod(mods)

	out.Relax = out.Mods.Has(constants.ModRelax)

	out.Passed, err = strconv.ParseBool(parts[14])
	if err != nil {
		return
	}

	mode, err := strconv.Atoi(parts[15])
	if err != nil {
		return
	}

	out.Mode = constants.Mode(mode)

	out.Timestamp, err = strconv.Atoi(parts[16])
	if err != nil {
		return
	}

	if len(parts[17]) < 8 {
		return
	}

	out.Flags = strings.Count(parts[17], " ")

	out.Version, err = strconv.Atoi(parts[17][:8])
	if err != nil {
		return
	}

	return
}
