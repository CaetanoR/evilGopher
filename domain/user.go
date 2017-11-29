package domain

import (
	"errors"
	"log"
)

type User struct {

	Name string
	Email string
	Nick string
	Password string
	Following map[*User]bool
	Followers map[*User]bool
	Tweets []*Tweet
	service UserService
}

func (u *User) Follow(user *User) {
	if u.Following == nil {
		u.Following = make(map[*User]bool)
	}

	if user.Followers == nil {
		user.Followers = make(map[*User]bool)
	}
	u.Following[user] = true
	user.Followers[u] = true
}

func (u *User) PublishTweet(tweetToPublish *Tweet) error {

	if tweetToPublish.User.Name == "" {
		return errors.New("user is required")
	}

	if tweetToPublish.Text == "" {
		return errors.New("text is required")
	}

	u.Tweets = append(u.Tweets, tweetToPublish)
	return nil
}

func NewUser(name string, email string, nick string, pass string, service UserService) *User {
	hashedPassword, err := service.HashPassword(pass)
	if err != nil {
		log.Fatalf("Error hashing password: %s", err.Error())
	}
	return &User{name, email, nick, hashedPassword, nil,nil, make([]*Tweet, 0), service}
}

type UserService interface {
	Users() ([]*User)
	HashPassword(pwd string) (string, error)
	CheckHash(pwd string, hash string) bool
	RegisterUser(u *User) error
	LogIn(userName string, password string) error
	Exists(u string, userList[]*User) *User
	Tweet(u* User, t *Tweet) error
}

