package layouts

type HwidMatch struct {
	User       int    `gorm:"check:user <> @user"`
	MacAddress string
	UniqueID   string
	DiskID     string
	Activated  bool
}