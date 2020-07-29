package icmd

import (
	"fmt"

	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/bot"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

const (
	Syntax  = "ak <user>"
	Aliases = []string{}
)

func Run(sender *objects.Player, args []string, target objects.Target) {
	if len(args) == 0 { // jg vs jne who will win
		return
	}

	if p := players.FindSafeUsername(args[0]); p != nil {
		p.Write(packets.MakePacket(36), packets.MakePacket(37))
		bot.WriteMessagef(target, "%v got eliminated.", p)
	}
}