package events

import (
	"time"

	"github.com/infernalfire72/flame/objects"
)

func Ping(p *objects.Player) {
	p.Ping = time.Now()
}
