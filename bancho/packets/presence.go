package packets

import (
	"github.com/infernalfire72/flame/objects"
)

func Presence(p *objects.Player) Packet {
	return MakePacket(83, 
		p.ID,
		p.Username,
		p.Timezone,
		p.Country,
		byte(p.IngamePrivileges & 0x1f) | ((p.Gamemode & 0x70) << 5),
		p.Longitude,
		p.Latitude,
		-1,
	)
}