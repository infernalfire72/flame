package utils

import (
	"github.com/infernalfire72/flame/config/database"
	"github.com/infernalfire72/flame/layouts"
)

// TODO: make this better
func GetFriends(user1 int) []layouts.UserRelationship {
	var friends []layouts.UserRelationship
	database.DB.Where(&layouts.UserRelationship{
		User1: user1,
	}).Find(&friends)
	return friends
}

func Has(a []layouts.UserRelationship, b int) bool {
	for _, v := range a {
		if v.User2 == b {
			return true
		}
	}

	return false
}
