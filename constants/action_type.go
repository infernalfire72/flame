package constants

type ActionType int

const (
	Idle ActionType = iota
	AFK
	Playing
	Editing
	Modding
	Multiplayer
	Watching
	Ranking
	Testing
	Submitting
	Paused
	Lobby
	Multiplaying
	Direct
)
