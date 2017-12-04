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

	testUser := domain.NewUser("grupoesfera", "grupoesfera@gmail.com", "ge", userPass, &userService)

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

func TestEditUserTweet(t *testing.T) {
	// Initialization
	var userService user.Service
	userService.Initialize()
	tweet.Initialize()

	testUserPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", testUserPass, &userService)

	tweetMessage,_ := domain.NewTweet(testUser, "hi Bro!")
	userService.RegisterUser(testUser)
	userService.LogIn(testUser.Name, testUserPass)
	userService.Tweet(testUser, tweetMessage)

	//operation
	tweetMessage.Text = "Hello bro!"
	err := userService.EditTweet(testUser, tweetMessage)

	//validation
	if err != nil {
		t.Errorf("expected no error, but got: %s", err.Error())
	}

}

func TestUserTimeline(t *testing.T) {
	// Initialization
	var userService user.Service
	userService.Initialize()

	testUserPass := "grupoesfera1234"
	testUserFollowedPass := "caetano1234"

	testUserTweet := "hi Bro!"
	testUserFollowedTweet := "good morning!"

	tweet.Initialize()
	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", testUserPass, &userService)
	testUserFollowed := domain.NewUser("caetano","caetano@gmail.com", "cae", testUserFollowedPass, &userService)

	testUserTweetMessage,_ := domain.NewTweet(testUser, testUserTweet)
	testUserFollowedTweetMessage,_ := domain.NewTweet(testUser, testUserFollowedTweet)

	userService.RegisterUser(testUser)
	userService.RegisterUser(testUserFollowed)

	userService.LogIn(testUser.Name, testUserPass)
	userService.LogIn(testUserFollowed.Name, testUserFollowedPass)


	userService.Tweet(testUserFollowed, testUserFollowedTweetMessage)
	userService.Tweet(testUser, testUserTweetMessage)

	testUser.Follow(testUserFollowed)

	//operation
	err, timeline := testUser.Timeline()

	if len(timeline) == 0 {
		t.Errorf("expected values")
	}

	if timeline[0].Text != testUserTweet && timeline[1].Text != testUserFollowedTweet {
		t.Error("tweets aren't ordered")
	}

	//validation
	if err != nil {
		t.Errorf("expected no error")
	}

}