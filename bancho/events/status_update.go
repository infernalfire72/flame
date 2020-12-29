package events

import (
	"github.com/infernalfire72/flame/bancho/models"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func StatusUpdate(p *objects.Player, bytes []byte) {
	status := &models.Status{}
	if err := status.Unmarshal(bytes); err == nil {
		p.Status = *status
		p.SetRelaxing((p.Mods & 128) != 0)
		players.Broadcast(packets.Stats(p))
	}
}
