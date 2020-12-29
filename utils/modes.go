package utils

import "github.com/infernalfire72/flame/constants"

var Modes = []string{"std", "taiko", "ctb", "mania",
	"std_relax", "taiko_relax", "ctb_relax"}

func ModeToDbSuffix(mode constants.Mode, relax bool) string {
	if relax && mode < 3 {
		return Modes[mode+4]
	} else if mode < 4 {
		return Modes[mode]
	} else {
		return Modes[0]
	}
}