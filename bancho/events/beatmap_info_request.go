package events

import (
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"
)

func BeatmapInfoRequest(p *objects.Player, bytes []byte) {
	var info layouts.BeatmapInfoRequest

	if err := layouts.ReadBeatmapInfoRequest(&info, bytes); err == nil {
		log.Info("Got", len(info.Names), "Beatmaps in Info Request (what do i do with this)")
		/*for _, id := range info.Names {

		}*/
	}
}
