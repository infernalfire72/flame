package layouts

import (
	"strconv"
	"strings"

	"github.com/infernalfire72/flame/constants"
)

type Score struct {
	ID          int
	BeatmapHash	string
	Username	string
	UserID      int
	ScoreHash	string

	N300		int
	N100		int
	N50			int
	NGeki		int
	NKatu		int
	NMiss		int

	TotalScore	int
	Combo		int
	FullCombo	bool
	Rank		string

	Mods		constants.Mod
	Passed		bool
	Mode		int

	Timestamp	int
	Version		int

	Flags		int
	
	Relax       bool
	Performance float32
	Status      int
}

func ReadScoreLayout(parts []string, out *Score) (err error) {
	out.BeatmapHash = parts[0]
	out.Username	= parts[1]
	out.ScoreHash	= parts[2]

	out.N300, err	= strconv.Atoi(parts[3])
	if err != nil {
		return
	}

	out.N100, err	= strconv.Atoi(parts[4])
	if err != nil {
		return
	}

	out.N50, err	= strconv.Atoi(parts[5])
	if err != nil {
		return
	}

	out.NGeki, err	= strconv.Atoi(parts[6])
	if err != nil {
		return
	}

	out.NKatu, err	= strconv.Atoi(parts[7])
	if err != nil {
		return
	}

	out.NMiss, err	= strconv.Atoi(parts[8])
	if err != nil {
		return
	}

	out.TotalScore, err	= strconv.Atoi(parts[9])
	if err != nil {
		return
	}

	out.Combo, err	= strconv.Atoi(parts[10])
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

	out.Mode, err = strconv.Atoi(parts[15])
	if err != nil {
		return
	}

	out.Timestamp, err = strconv.Atoi(parts[16])
	if err != nil {
		return
	}

	out.Flags = strings.Count(parts[17], " ")

	out.Version, err = strconv.Atoi(parts[17][:8])
	if err != nil {
		return
	}

	return
}