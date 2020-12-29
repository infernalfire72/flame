package packets

import (
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/objects"
)

func Stats(p *objects.Player) Packet {
	var s *layouts.Stats

	if p.Relaxing {
		s = p.Stats[p.Gamemode]
	} else {
		s = p.Stats[p.Gamemode]
	}

	return MakePacket(11,
		p.ID,
		p.Action,
		p.InfoText,
		p.BeatmapHash,
		p.Mods,
		p.Gamemode,
		p.Beatmap,
		s.RankedScore,
		s.Accuracy,
		s.Playcount,
		s.TotalScore,
		s.Rank,
		s.Performance)
}
