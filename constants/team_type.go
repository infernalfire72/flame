package constants

type MatchTeamType byte

const (
	TeamHeadToHead MatchTeamType = iota
	TeamTagCoop
	TeamVs
	TeamTagVs
)
