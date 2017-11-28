package domain

import (
	"errors"
)

type User struct {

	Name string
	Following map[*User]bool
	Followers map[*User]bool
	Tweets []*Tweet
	service UserService
}

func (u *User) Follow(user *User) {
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

	if !u.service.Exists(tweetToPublish.User) {
		return errors.New("user must be registered in order to publish tweets")
	}

	u.Tweets = append(u.Tweets, tweetToPublish)
	return nil
}

func NewUser(name string, service UserService) *User {
	return &User{name, nil, nil, nil, service}
}

type UserService interface {
	Users() ([]*User)
	RegisterUser(u *User) error
	Exists(u *User) bool
	Tweet(u* User, t *Tweet) error
}

