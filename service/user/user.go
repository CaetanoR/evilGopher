package user

import (
	"errors"
	"github.com/evilGopher/domain"
	"github.com/evilGopher/service/tweet"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func (s *Service) CheckHash(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func (s *Service) RegisterUser(user *domain.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	for _, curUser := range users {
		if user.Name == curUser.Name || user.Email == curUser.Email {
			return errors.New("user already exists")
		}
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