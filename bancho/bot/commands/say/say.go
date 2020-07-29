package icmd

import (
	"fmt"

	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/bot"
	"github.com/infernalfire72/flame/bancho/packets"
)

const (
	Syntax  = "say <words>"
	Aliases = []string{}
)

func Run(sender *objects.Player, args []string, target objects.Target) {
	res, err := sender.AwaitMessage(5000)
	if err != nil {
		bot.WriteMessagef(target, "Command timed out")
		return
	}

	bot.WriteMessagef(target, res)
}