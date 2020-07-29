package bot

import (
	"fmt"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
	"github.com/infernalfire72/flame/cache/users"
	"github.com/infernalfire72/flame/cache/users/stats"
)

var Player *objects.Player

const ID = 999

func init() {
	go func() {
		for config.Database == nil {

		}

		log.Info("Initializing Bot")

		Player = &objects.Player{
			User:  users.Get(ID),
			Stats: stats.Get(ID),
		}
		Player.IngamePrivileges = Player.Privileges.BanchoPrivileges()

		players.Add(Player)
	}()
}

func WriteMessagef(target objects.Target, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	target.Write(packets.IrcMessage(layouts.Message{
		Player.Username, msg, target.GetName(), int32(Player.ID),
	}))
}
