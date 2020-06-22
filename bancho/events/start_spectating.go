package events

import (
	"unsafe"

	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
	"github.com/infernalfire72/flame/bancho/players"
)

func StartSpectating(p *objects.Player, bytes []byte) {
	if len(bytes) != 4 {
		return
	}

	if p.Spectating != nil {
		StopSpectating(p)
	}

	targetID := *(*int32)(unsafe.Pointer(&bytes[0]))
	if target := players.Find(int(targetID)); target != nil {
		p.Spectating = target

		fellowSpectator := packets.NewFellowSpectator(int32(p.ID))
		target.SpectatorMutex.RLock()
		channelInfo := packets.AvailableChannelArgs("#spectator", "Spectator Channel for " + target.Username, int16(len(target.Spectators) + 1))
		for _, t := range target.Spectators {
			t.Write(channelInfo, fellowSpectator)
			p.Write(packets.NewFellowSpectator(int32(t.ID)))
		}
		target.SpectatorMutex.RUnlock()
		target.AddSpectator(p)

		target.Write(channelInfo, packets.NewSpectator(int32(p.ID)))
		p.Write(packets.JoinedChannel("#spectator"), channelInfo)

		log.Info(p, "started spectating", target)
	}
}