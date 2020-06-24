package events

import (
	"encoding/binary"

	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"
)

func AddFriend(p *objects.Player, bytes []byte) {
	if len(bytes) != 4 {
		return
	}

	id := binary.LittleEndian.Uint32(bytes)
	config.Database.Exec("REPLACE INTO users_relationships VALUES (?, ?)", p.ID, id)
	log.Info(p, "added", id, "to their friends list.")
}
