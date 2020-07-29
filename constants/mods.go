package constants

type Mod int32

const (
	ModEasy Mod = 1 << iota
	ModNoFail
	ModTouchDevice
	ModHidden
	ModHardRock
	ModSuddenDeath
	ModDoubletime
	ModRelax
	ModHalftime
	ModNightcore
	ModFlashlight
	ModAutoplay
	ModSpunOut
	ModAutopilot
	ModPerfect
	Mod4Keys
	Mod5Keys
	Mod6Keys
	Mod7Keys
	Mod8Keys
	ModFadeIn
	ModRandom
	ModCinema
	ModTargetPractice
	Mod9Keys
	ModCoop
	Mod1Key
	Mod2Keys
	Mod3Keys
	ModScoreV2

	ModsChangeSpeed = ModHalftime | ModDoubletime | ModNightcore
)

func (a Mod) Has(b Mod) bool {
	return b == 0 || (a & b) != 0
}