package events

import (
	"encoding/binary"
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/layouts"

	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"
)

func RemoveFriend(p *objects.Player, bytes []byte) {
	if len(bytes) != 4 {
		return
	}

	id := binary.LittleEndian.Uint32(bytes)
	database.DB.Delete(&layouts.UserRelationship{
		User1: p.ID,
		User2: int(id),
	})
	log.Info(p, "removed", id, "from their friends list.")
}
