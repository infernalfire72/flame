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
	ID					int			`json:"id"`
	Username			string
	SafeUsername		string
	Password			string
	Token				string
	IngamePrivileges	constants.BanchoPrivileges
	Privileges			constants.AkatsukiPrivileges

	Country				byte
	Timezone			byte
	Longitude			float32
	Latitude			float32

	VanillaStats	[4]layouts.ModeData
	RelaxStats		[3]layouts.ModeData

	layouts.Status
	Relaxing		bool

	Channels		[]*Channel
	Spectators		[]*Player
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

func (p *Player) Write(data ...[]byte) {
	p.Mutex.Lock()
	for _, segment := range data {
		p.Queue.WriteByteSlice(segment)
	}
	p.Mutex.Unlock()
}