package layouts

type UserRelationship struct {
	User1 int `gorm:"primaryKey"`
	User2 int `gorm:"primaryKey"`
}