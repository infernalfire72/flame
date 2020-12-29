package clans

type Clan struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Tag         string `gorm:"unique"`
	Description string
	Owner       int
}