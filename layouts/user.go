package layouts

import "github.com/infernalfire72/flame/constants"

type User struct {
	ID           int
	Username     string
	SafeUsername string
	Password     string
	Country      byte
	Privileges   constants.AkatsukiPrivileges
}

/*func (u *User) VerifyPassword(password string) bool {
	if len(u.Password) == 32 {
		return u.Password == password
	} else {
		var dbPassword string
		config.Database.Get(&dbPassword, "SELECT password_md5 FROM users WHERE id = ?", u.ID)

		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password)); err == nil {
			u.Password = password
			return true
		} else {
			return false
		}
	}
}*/
