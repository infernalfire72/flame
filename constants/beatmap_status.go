package constants

type BeatmapStatus int

const (
	StatusUnknown BeatmapStatus = iota - 2
	StatusNotSubmitted
	StatusPending
	StatusNeedUpdate
	StatusRanked
	StatusApproved
	StatusQualified
	StatusLoved
)
