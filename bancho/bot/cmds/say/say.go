package icmd

import (
	"fmt"

	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

const (
	Syntax  = "say <words>"
	Aliases = []string{}
)

func Run(sender *objects.Player, args []string, target objects.Target) {
	res, err := sender.AwaitMessage(5000)
	if err != nil {
		target.Write(packets.IrcMessageArgs("Bot", "Command timed out", target.GetName(), 999))
		return
	}

	target.Write(packets.IrcMessageArgs("Bot", res, target.GetName(), 999))
}