package symbols

import (
	"reflect"

	"github.com/containous/yaegi/interp"

	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/bot"
	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

var Bancho = interp.Exports{
	"github.com/infernalfire72/flame/objects": map[string]reflect.Value{
		"Player":  reflect.ValueOf((*objects.Player)(nil)),
		"Channel": reflect.ValueOf((*objects.Channel)(nil)),
		"Target":  reflect.ValueOf((*objects.Target)(nil)),
	},

	"github.com/infernalfire72/flame/layouts": map[string]reflect.Value{
		"Message": reflect.ValueOf((*layouts.Message)(nil)),
	},

	"github.com/infernalfire72/flame/bancho/players": map[string]reflect.Value{
		"FindSafeUsername": reflect.ValueOf(players.FindSafeUsername),
	},

	"github.com/infernalfire72/flame/bancho/packets": map[string]reflect.Value{
		"Packet":         reflect.ValueOf((*packets.Packet)(nil)),
		"MakePacket":     reflect.ValueOf(packets.MakePacket),
		"Alert":          reflect.ValueOf(packets.Alert),
		"IrcMessageArgs": reflect.ValueOf(packets.IrcMessageArgs),
		"Stats":          reflect.ValueOf(packets.Stats),
		"Presence":       reflect.ValueOf(packets.Presence),
	},

	"github.com/infernalfire72/flame/bancho/bot": map[string]reflect.Value{
		"WriteMessagef": reflect.ValueOf(bot.WriteMessagef),
	},
}