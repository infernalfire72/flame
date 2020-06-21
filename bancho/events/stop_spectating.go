package events

import (
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/objects"

	"github.com/infernalfire72/flame/bancho/packets"
)

func StopSpectating(p *objects.Player) {
	if p.Spectating == nil {
		return
	}

	host := p.Spectating
	p.Spectating = nil

	host.RemoveSpectator(p)
	revoked := packets.ChannelRevoked("#spectator")
	p.Write(revoked)

	host.SpectatorMutex.RLock()

	if len(host.Spectators) == 0 {
		host.Write(revoked)
	} else {
		channelInfo := packets.AvailableChannelArgs("#spectator", "Spectator Channel for " + host.Username, int16(len(host.Spectators)))
		fellowSpectator := packets.FellowSpectatorLeft(int32(p.ID))

		host.Write(channelInfo)
		
		for _, t := range host.Spectators {
			t.Write(channelInfo, fellowSpectator)
		}
	}

	host.SpectatorMutex.RUnlock()

	host.Write(packets.SpectatorLeft(int32(p.ID)))
	log.Info(p.Username, "stopped spectating", host.Username)
}