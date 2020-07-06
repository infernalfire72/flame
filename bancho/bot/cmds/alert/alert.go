package icmd

import (
	"fmt"
	"strings"

	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

const (
	Syntax  = "alert <words>"
	Aliases = []string{}
)

func Run(sender *objects.Player, args []string, target objects.Target) {
	phrase := strings.Join(args, " ")
	if len(phrase) == 0 {
		return
	}

	target.Write(packets.Alert(phrase))
}