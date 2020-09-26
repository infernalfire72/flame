package constants

type ScoreStatus int

const (
	ScoreFailed ScoreStatus = iota
	ScorePassed
	_
	ScoreBestPerformance
)