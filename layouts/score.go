package layouts

import (
	"github.com/infernalfire72/flame/constants"
	"strconv"
	"strings"
	"time"
)

type Score struct {
	ID      int `gorm:"primaryKey"`
	Hash    string
	Beatmap string
	User    int
	Mode    constants.Mode

	TotalScore   int
	Combo        int16
	PerfectCombo bool
	Passed       bool `gorm:"-"`
	Rank         string

	Mods     constants.Mod
	Accuracy `gorm:"embedded"`

	Time            time.Time
	ClientTimestamp int
	Status          int
	Performance     float32 `gorm:"column:pp"`
	Flags           int
}

func (out *Score) Read(parts []string) (err error) {
	out.Beatmap = parts[0]
	out.Hash = parts[2]

	n300, err := strconv.ParseInt(parts[3], 10, 16)
	if err != nil {
		return
	}

	n100, err := strconv.ParseInt(parts[4], 10, 16)
	if err != nil {
		return
	}

	n50, err := strconv.ParseInt(parts[5], 10, 16)
	if err != nil {
		return
	}

	nGeki, err := strconv.ParseInt(parts[6], 10, 16)
	if err != nil {
		return
	}

	nKatu, err := strconv.ParseInt(parts[7], 10, 16)
	if err != nil {
		return
	}

	nMiss, err := strconv.ParseInt(parts[8], 10, 16)
	if err != nil {
		return
	}

	out.Accuracy = Accuracy{int16(n300), int16(n100), int16(n50), int16(nGeki), int16(nKatu), int16(nMiss)}

	out.TotalScore, err = strconv.Atoi(parts[9])
	if err != nil {
		return
	}

	combo, err := strconv.ParseInt(parts[10], 10, 16)
	if err != nil {
		return
	}
	out.Combo = int16(combo)

	out.PerfectCombo, err = strconv.ParseBool(parts[11])
	if err != nil {
		return
	}

	out.Rank = parts[12]

	mods, err := strconv.Atoi(parts[13])
	if err != nil {
		return
	}
	out.Mods = constants.Mod(mods)

	out.Passed, err = strconv.ParseBool(parts[14])
	if err != nil {
		return
	}

	mode, err := strconv.Atoi(parts[15])
	if err != nil {
		return
	}

	out.Mode = constants.Mode(mode)

	out.ClientTimestamp, err = strconv.Atoi(parts[16])
	if err != nil {
		return
	}

	if len(parts[17]) < 8 {
		return
	}

	out.Flags = strings.Count(parts[17], " ")

	return
}