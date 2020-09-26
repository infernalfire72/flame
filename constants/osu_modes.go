package constants

type Mode byte

const (
	ModeStandard Mode = iota
	ModeTaiko
	ModeCatch
	ModeMania
)

func (m *Mode) Clamp() {
	if *m < ModeStandard || *m > ModeMania {
		*m = ModeStandard
	}
}