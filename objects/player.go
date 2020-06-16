package objects

import (
	"sync"
	"time"

	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/layouts"
)

var PlayerMutex sync.RWMutex
var Players PlayerCollection

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

	VanillaStats	[4]ModeData
	RelaxStats		[3]ModeData

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

func NewPlayer(id int) *Player {
	return &Player {
		ID:		id,
		Queue:	io.NewStreamWithCapacity(1024),
		Ping:	time.Now(),
	}
}

func (p *Player) SetRelaxing(relaxing bool) {

}

func (p *Player) Write(data ...[]byte) {
	p.Mutex.Lock()
	for _, segment := range data {
		p.Queue.WriteByteSlice(segment)
	}
	p.Mutex.Unlock()
}