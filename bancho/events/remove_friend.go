package events

import (
	"encoding/binary"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"
)

func RemoveFriend(p *objects.Player, bytes []byte) {
	if len(bytes) != 4 {
		return
	}

	id := binary.LittleEndian.Uint32(bytes)
	config.Database.Exec("DELETE FROM users_relationships WHERE user1 = ? AND user2 = ?", p.ID, id)
	log.Info(p, "removed", id, "from their friends list.")
}
