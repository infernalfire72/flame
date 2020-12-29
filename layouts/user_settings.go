package layouts

type UserSettings struct {
	User                    int    `gorm:"primaryKey"`
	AlternativeUsername     string
	UseAlternativeUsername  bool
	ShowAlternativeUsername bool // This is purely for Ingame purposes!
	PreferRelax             bool
}