package user_test

import (
	"testing"
	"github.com/evilGopher/service/user"
	"github.com/evilGopher/service/tweet"
	"github.com/evilGopher/domain"
	"fmt"
)

func TestEditUserTweet(t *testing.T) {
	// Initialization
	var userService user.Service
	userService.Initialize()

	tweet.Initialize()
	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", "grupoesfera1234", &userService)

	tweetMessage,_ := domain.NewTweet(testUser, "hi Bro!")
	fmt.Println(tweetMessage.Id)
	userService.RegisterUser(testUser)
	userService.Tweet(testUser, tweetMessage)

	//operation
	tweetMessage.Text = "Hello bro!"
	fmt.Println(tweetMessage.Id)
	err := userService.EditTweet(testUser, tweetMessage)

	//validation
	if err != nil {
		t.Errorf("expected no error")
	}

}