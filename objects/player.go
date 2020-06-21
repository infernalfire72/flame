package objects

import (
	"fmt"
	"sync"
	"time"

	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/config"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
	"github.com/infernalfire72/flame/log"
	"github.com/infernalfire72/flame/utils"
)

type Player struct {
	ID					int
	Username			string
	SafeUsername		string
	Password			string
	Country				byte
	Privileges			constants.AkatsukiPrivileges
	IngamePrivileges	constants.BanchoPrivileges
	Token				string

	Timezone			byte
	Longitude			float32
	Latitude			float32

	VanillaStats	[4]layouts.ModeData
	RelaxStats		[3]layouts.ModeData

	layouts.Status
	Relaxing		bool

	Channels		[]*Channel
	ChannelMutex	sync.RWMutex

	Spectators		[]*Player
	SpectatorMutex	sync.RWMutex
	Spectating		*Player

	Match			*MultiplayerLobby

	Ping			time.Time
	Queue			*io.Stream
	Mutex			sync.Mutex
}

func (p *Player) SetRelaxing(relaxing bool) {
	if relaxing == p.Relaxing {
		return
	}

	if relaxing && p.Gamemode >= 3 {
		p.Gamemode = 0
	}

	p.Relaxing = relaxing
	p.SetStats(p.Gamemode, relaxing)
}

func (p *Player) SetStats(mode byte, relax bool) {
	var (
		s		*layouts.ModeData
		table	string
	)

	if p.Relaxing {
		s = &p.RelaxStats[p.Gamemode]
		table = "rx_stats"
	} else {
		s = &p.VanillaStats[p.Gamemode]
		table = "users_stats"
	}

	query := fmt.Sprintf("SELECT total_score_%[1]s AS total_score, ranked_score_%[1]s AS ranked_score, pp_%[1]s AS pp, playcount_%[1]s AS playcount, avg_accuracy_%[1]s/100 AS accuracy FROM %[2]s WHERE id = ?", utils.DbMode(p.Gamemode), table)
	err := config.Database.Get(s, query, p.ID)
	if err != nil {
		log.Error(err)
	}
	
	err = config.Database.Get(&s.Rank, fmt.Sprintf("SELECT COUNT(id) + 1 FROM %s WHERE pp_%s > ?", table, utils.DbMode(p.Gamemode)), s.Performance)
	if err != nil {
		log.Error(err)
	}
}

func (p *Player) AddChannel(c *Channel) {
	p.ChannelMutex.Lock()
	p.Channels = append(p.Channels, c)
	p.ChannelMutex.Unlock()
}

func (p *Player) RemoveChannel(c *Channel) {
	p.ChannelMutex.Lock()
	for i, t := range p.Channels {
		if t == c {
			p.Channels[i] = p.Channels[len(p.Channels)-1]
			p.Channels[len(p.Channels)-1] = nil
			p.Channels = p.Channels[:len(p.Channels)-1]
			break
		}
	}
	p.ChannelMutex.Unlock()
}

func (host *Player) AddSpectator(p *Player) {
	host.SpectatorMutex.Lock()
	host.Spectators = append(host.Spectators, p)
	host.SpectatorMutex.Unlock()
}

func (host *Player) RemoveSpectator(p *Player) {

}

func (p *Player) Write(data ...[]byte) {
	p.Mutex.Lock()
	for _, segment := range data {
		p.Queue.WriteByteSlice(segment)
	}
	p.Mutex.Unlock()
}