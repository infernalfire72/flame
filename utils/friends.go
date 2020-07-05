package utils

import (
	"github.com/infernalfire72/flame/config"
)

// TODO: make this better
func GetFriends(user1 int) []int {
	var friends []int
	config.Database.Select(&friends, "SELECT user2 FROM users_relationships WHERE user1 = ?", user1)
	return friends
}

func Has(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}

	return false
}
