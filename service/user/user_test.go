package user_test

import (
	"testing"
	"github.com/evilGopher/service/tweet"
	"github.com/evilGopher/domain"
	"github.com/evilGopher/service/user"
)

func TestFollowUser(t *testing.T) {
	// Initialization

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", "grupoesfera1234", &userService)
	testUserToFollow := domain.NewUser("federico","federico@gmail.com", "fede", "federico1234", &userService)

	userService.RegisterUser(testUser)
	userService.RegisterUser(testUserToFollow)

	testUser.Follow(testUserToFollow)

	if testUser.Following[testUserToFollow] == false {
		t.Errorf("User should be following the second used")
	}
	if testUserToFollow.Followers[testUser] == false {
		t.Errorf("The second user should be followed by the first user")
	}
}

func TestService_RegisterUser(t *testing.T) {

	// Initialization

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)

	err := userService.RegisterUser(testUser)

	if err != nil {
		t.Errorf("User registration failed: %s", err.Error())
	}

	err = userService.LogIn(testUser.Name, userPass)

	if err != nil {
		t.Errorf("User login failed: %s", err.Error())
	}

}