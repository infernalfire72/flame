package bot

import (
	"fmt"
	"github.com/infernalfire72/flame/bancho/models"
	"github.com/infernalfire72/flame/cache/users/stats"
	"github.com/infernalfire72/flame/constants"

	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
	"github.com/infernalfire72/flame/cache/users"
)

var Player *objects.Player

const ID = 999

func Setup() {
	log.Info("Initializing Bot")

	Player = &objects.Player{
		User:  users.Get(ID),
	}

	for i := constants.ModeStandard; i <= constants.ModeMania; i++ {
		Player.Stats[i] = stats.Get(stats.Identifier{
			User:  ID,
			Mode:  i,
			Relax: false,
		})
	}

	for i := constants.ModeStandard; i < constants.ModeMania; i++ {
		Player.Stats[i] = stats.Get(stats.Identifier{
			User:  ID,
			Mode:  i,
			Relax: true,
		})
	}

	Player.IngamePrivileges = Player.Privileges.BanchoPrivileges()

	players.Add(Player)
}

func WriteMessagef(target objects.Target, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	target.Write(packets.IrcMessage(models.Message{
		Player.Username, msg, target.GetName(), int32(Player.ID),
	}))
}
