package user

import (
	"errors"
	"github.com/evilGopher/domain"
	"github.com/evilGopher/service/tweet"
	"golang.org/x/crypto/bcrypt"
)

var registeredUsers []*domain.User
var loggedUsers []*domain.User

type Service struct {
}

func (s *Service) Initialize() {
	registeredUsers = []*domain.User{}
	loggedUsers = []*domain.User{}
}

func (s *Service) Users() ([]*domain.User) {
	return registeredUsers
}

func (s *Service) HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func (s *Service) CheckHash(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func (s *Service) RegisterUser(u *domain.User) error {

	if u.Name == "" {
		return errors.New("name is required")
	}

	for _, curUser := range registeredUsers {
		if u.Name == curUser.Name || u.Email == curUser.Email {
			return errors.New("user already exists")
		}
	}

	registeredUsers = append(registeredUsers, u)
	return nil
}

func (s *Service) LogIn(userName string, password string) error {
	userToLogIn := s.Exists(userName, registeredUsers)
	if userToLogIn == nil {
		return errors.New("user doesn't exist")
	}
	if !s.CheckHash(password, userToLogIn.Password) {
		return errors.New("invalid password")
	}
	loggedUsers = append(loggedUsers, userToLogIn)
	return nil
}

func (s *Service) LogOut(userName string, password string) error {

	err, userRemoval := s.removeUser(loggedUsers, userName, password)

	if err != nil {
		return errors.New(err.Error())
	}
	if userRemoval != nil {
		loggedUsers = userRemoval
	}
	return nil
}

func (s *Service) IsLoggedIn(userName string) bool {

	isLoggedIn := false

	for _, curUser := range loggedUsers {
		if curUser.Name == userName {
			isLoggedIn = true
		}
	}

	return isLoggedIn
}

func (s *Service) removeUser(list []*domain.User, userName, password string) (error, []*domain.User, ) {

	for i, curUser := range loggedUsers {
		if curUser.Name == userName {
			if !s.CheckHash(password, curUser.Password) {
				return errors.New("invalid password"), nil
			}
			return nil, append(loggedUsers[:i], loggedUsers[i+1:]...)
		}
	}
	return nil, nil
}

func (s *Service) Exists(userToSearch string, userList []*domain.User) *domain.User {
	var user *domain.User = nil
	for _,curUser := range userList {
		if userToSearch == curUser.Name {
			user = curUser
		}
	}
	return user
}

func (s *Service) Tweet(u *domain.User, t *domain.Tweet) error {

	if s.Exists(u.Name, loggedUsers) == nil {
		return errors.New("user must be logged in order to publish tweets")
	}

	err := u.PublishTweet(t)
	if err == nil {
		tweet.AddTweet(t)
	}
	return err
}