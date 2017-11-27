package user

import (
	"errors"
)

var users []*User

func Create(user *User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	users = append(users, user)
	return nil
}

func Exists(userToSearch *User) bool{
	exists := false
	for _,user := range users {
		if userToSearch == user {
			exists = true
		}
	}
	return exists
}