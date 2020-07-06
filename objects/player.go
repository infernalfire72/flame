package objects

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/infernalfire72/flame/constants"
	"github.com/infernalfire72/flame/io"
	"github.com/infernalfire72/flame/layouts"
)

type Player struct {
	*layouts.User
	IngamePrivileges constants.BanchoPrivileges
	Token            string

	Timezone  byte
	Longitude float32
	Latitude  float32

	*layouts.Stats

	layouts.Status
	Relaxing bool

	Channels     []*Channel
	ChannelMutex sync.RWMutex

	Spectators     []*Player
	SpectatorMutex sync.RWMutex
	Spectating     *Player

	IsLobby bool
	Match   *MultiplayerLobby

	Ping      time.Time
	LoginTime time.Time
	Queue     *io.Stream
	Mutex     sync.Mutex

	AwaiterMutex   sync.RWMutex
	MessageAwaiter chan string
}

func (p *Player) AwaitMessage(timeout time.Duration) (string, error) {
	p.AwaiterMutex.RLock()
	if p.MessageAwaiter != nil {
		p.AwaiterMutex.RUnlock()
		return "", errors.New("cmderr: already waiting")
	}

	p.AwaiterMutex.RUnlock()

	p.AwaiterMutex.Lock()
	p.MessageAwaiter = make(chan string)
	p.AwaiterMutex.Unlock()

	defer func(){
		p.AwaiterMutex.Lock()
		p.MessageAwaiter = nil
		p.AwaiterMutex.Unlock()
	}()

	select {
	case res := <-p.MessageAwaiter:
		return res, nil
	case <-time.After(timeout * time.Millisecond):
		return "", errors.New("cmderr: timeout")
	}
}

func (p *Player) SetRelaxing(relaxing bool) {
	if relaxing == p.Relaxing {
		return
	}

	if relaxing && p.Gamemode >= 3 {
		p.Gamemode = 0
	}

	p.Relaxing = relaxing
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
	host.SpectatorMutex.Lock()
	for i, t := range p.Spectators {
		if t == p {
			p.Spectators[i] = p.Spectators[len(p.Spectators)-1]
			p.Spectators[len(p.Spectators)-1] = nil
			p.Spectators = p.Spectators[:len(p.Spectators)-1]
		}
	}
	host.SpectatorMutex.Unlock()
}

func (p *Player) Write(data ...[]byte) {
	p.Mutex.Lock()
	for _, segment := range data {
		p.Queue.WriteByteSlice(segment)
	}
	p.Mutex.Unlock()
}

func (p Player) String() string {
	return fmt.Sprintf("%s (%d)", p.Username, p.ID)
}
