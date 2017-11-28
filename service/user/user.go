package user

import (
	"errors"
	"github.com/evilGopher/domain"
	"github.com/evilGopher/service/tweet"
)

var users []*domain.User

type Service struct {
}

func (s *Service) Initialize() {
	users = []*domain.User{}
}

func (s *Service) Users() ([]*domain.User) {
	return users
}

func (s *Service) RegisterUser(user *domain.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	users = append(users, user)
	return nil
}

func (s *Service) Exists(userToSearch *domain.User) bool {
	exists := false
	for _,user := range users {
		if userToSearch == user {
			exists = true
		}
	}
	return exists
}

func (s *Service) Tweet(u *domain.User, t *domain.Tweet) error {
	err := u.PublishTweet(t)
	if err == nil {
		tweet.AddTweet(t)
	}
	return err
}