package layouts

import (
	"fmt"
	"github.com/infernalfire72/flame/cache/clans"
	"github.com/infernalfire72/flame/constants"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"unique;not null"`
	SafeUsername string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Country      string `gorm:"type:CHAR(2);default:'XX';not null"`
	Privileges   constants.AkatsukiPrivileges `gorm:"default:1048576"`
	ClanID       int
}

func (u *User) VerifyPassword(password string) bool {
	if len(u.Password) == 32 {
		return u.Password == password
	} else if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
		u.Password = password
		return true
	} else {
		return false
	}
}

func (u *User) Clan() *clans.Clan {
	if u.ClanID == 0 {
		return nil
	}

	return clans.Get(u.ClanID)
}

func (u *User) FullName() string {
	if c := u.Clan(); c != nil {
		return fmt.Sprintf("[%s] %s", c.Tag, u.Username)
	} else {
		return u.Username
	}
}