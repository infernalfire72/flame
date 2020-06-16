package objects

type PlayerCollection map[string]*Player

func (c PlayerCollection) FindPlayer(id int) *Player {
	for _, a := range c {
		if a.ID == id {
			return a
		}
	}
	return nil
}

func (c PlayerCollection) Broadcast(data []byte) {
	for _, a := range c {
		a.Write(data)
	}
}