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
		t.Errorf("User should be following the second user")
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

	if !userService.IsLoggedIn(testUser.Name) {
		t.Error("user is still logged in")
	}

	if err != nil {
		t.Errorf("User login failed: %s", err.Error())
	}

}

func TestService_LogOut(t *testing.T) {
	// Initialization

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)

	err := userService.RegisterUser(testUser)

	if err != nil {
		t.Errorf("no error was expected, but got: %s", err.Error())
	}

	userService.LogIn(testUser.Name, userPass)

	userService.LogOut(testUser.Name, userPass)

	if userService.IsLoggedIn(testUser.Name) {
		t.Error("user is still logged in")
	}
}