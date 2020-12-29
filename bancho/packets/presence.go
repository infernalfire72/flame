package packets

import (
	"github.com/infernalfire72/flame/objects"
	"github.com/infernalfire72/flame/utils"
)

func Presence(p *objects.Player) Packet {
	return MakePacket(83,
		p.ID,
		p.Username,
		p.Timezone,
		utils.CountryByte(p.Country),
		byte(p.IngamePrivileges&0x1f)|((p.Gamemode&0x70)<<5),
		p.Longitude,
		p.Latitude,
		-1,
	)
}
