package icmd

import (
	"strings"
	//"time"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

const (
	Syntax  = "fryme"
	Aliases = []string{}
)

func Run(sender *objects.Player, args []string, target objects.Target) {
	//time.Sleep(2 * time.Second)
	/*sender.Write(packets.MakePacket(23), packets.IrcMessageArgs("BanchoBot", "Helo", sender.Username, 3))
	sender.Write(packets.MakePacket(64, "Bot"))
	sender.Write(packets.IrcMessageArgs("Bot", "Helo 2", sender.Username, 999))

	sender.WriteDelayed(packets.MakePacket(66, "BanchoBot"), 1)
	sender.Write(packets.MakePacket(8))*/

	/*sender.Write(packets.MakePacket(23),
		packets.IrcMessageArgs("BanchoBot", "Helo", sender.Username, 3),
		packets.MakePacket(66, "#osu"))

	sender.WriteDelayed(packets.MakePacket(64, "#osu"), 3)
	sender.WriteDelayed(packets.MakePacket(66, "BanchoBot"), 1)
	sender.WriteDelayed(packets.MakePacket(89), 1)*/

	//sender.Write(packets.MakePacket(86, 10))

	sender.Write(packets.MakePacket(12, sender.ID, 0))

	/*if len(args) != 0 {
		name := strings.Join(args, " ")

		sender.Username = name

		sender.Write(packets.Presence(sender))
	}*/
}
