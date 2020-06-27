package layouts

import (
	"fmt"

	"github.com/infernalfire72/flame/constants"
	"golang.org/x/crypto/bcrypt"

	"github.com/infernalfire72/flame/cache/clans"
)

type User struct {
	ID           int    `db:"id"`
	Username     string `db:"username"`
	SafeUsername string `db:"username_safe"`
	Password     string `db:"password_md5"`
	Country      byte
	Privileges   constants.AkatsukiPrivileges `db:"privileges"`
	ClanID       int                          `db:"clan_id"`
}

func (u *User) VerifyPassword(password string) bool {
	if len(u.Password) == 32 {
		return u.Password == password
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
			u.Password = password
			return true
		} else {
			return false
		}
	}
}

func (u *User) FullName() string {
	if u.ClanID == 0 {
		return u.Username
	}

	if c := clans.Get(u.ClanID); c != nil {
		return fmt.Sprintf("[%s] %s", c.Tag, u.Username)
	} else {
		return u.Username
	}
}
