package tweet_test

import (
	"testing"
	"github.com/evilGopher/service/tweet"
	"github.com/evilGopher/service/user"
	"github.com/evilGopher/domain"
)

func TestPublishIsSaved(t *testing.T) {

	// Initialization
	var tweetMessage *domain.Tweet

	var userService user.Service


	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)
	text := "This is my first tweetMessage"

	tweetMessage,_ = domain.NewTweet(testUser, text)

	// Operation
	userService.RegisterUser(testUser)
	userService.LogIn(testUser.Name, userPass)
	err := userService.Tweet(testUser, tweetMessage)

	// Validation
	if err != nil {
		t.Errorf("Didn't expect any error, but got: %s", err.Error())
	}

	publishedTweet := tweet.GetAll()[0]

	if publishedTweet.User != testUser &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweetMessage is %s: %s \nbut is %s: %s",
			testUser, text, publishedTweet.User, publishedTweet.Text)
	}

}

func TestDuplicatedTweet(t *testing.T) {

	// Initialization
	var tweetMessage *domain.Tweet

	var userService user.Service


	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)
	text := "This is my first tweetMessage"

	tweetMessage,_ = domain.NewTweet(testUser, text)

	// Operation
	userService.RegisterUser(testUser)
	userService.LogIn(testUser.Name, userPass)
	err := userService.Tweet(testUser, tweetMessage)
	err2 := userService.Tweet(testUser, tweetMessage)

	// Validation
	if err != nil {
		t.Errorf("Didn't expect any error, but got: %s", err.Error())
	}

	if err2.Error() != "user can't have duplicated tweets" {
		t.Errorf("expected error: user can't have duplicated tweets, but got: %s", err2.Error())
	}

}

func TestPasswordHash(t *testing.T) {

	// Initialization
	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	userPassword := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPassword, &userService)

	if !userService.CheckHash(userPassword, testUser.Password) {
		t.Error("Password isn't beint hashed correctly")
	}

}


func TestPublishWithoutTextError(t *testing.T) {

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)

	userService.RegisterUser(testUser)
	userService.LogIn(testUser.Name, userPass)

	tweetMessage,_ := domain.NewTweet(testUser, "")

	err := userService.Tweet(testUser, tweetMessage)

	if err == nil {
		t.Error("error was expected")
	}

	if err.Error() != "text is required" {
		t.Errorf("expected error: text is required, but was %s", err.Error())
	}

}

func TestPublishWithoutLoggedInUser(t *testing.T) {

	// Initialization
	var tweetMessage *domain.Tweet
	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)


	userService.RegisterUser(testUser)

	text := "This is my first tweet"

	tweetMessage,_ = domain.NewTweet(testUser, text)

	// Operation
	err := userService.Tweet(testUser, tweetMessage)

	// Validation
	if err.Error() == "" {
		t.Errorf("expected error")
	}

	if err.Error() != "user must be logged in order to publish tweets" {
		t.Errorf("expected error: testUser must be logged in order to publish tweets, but was: %s", err.Error())
	}

}

func TestGetById(t *testing.T) {

	// Initialization

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	userPass := "grupoesfera1234"

	testUser := domain.NewUser("grupoesfera","grupoesfera@gmail.com", "ge", userPass, &userService)


	userService.RegisterUser(testUser)
	userService.LogIn(testUser.Name, userPass)

	messages := []string{"This is my first tweetMessage", "This is my second tweetMessage", "This is my third tweetMessage"}

	testTweets := make([]*domain.Tweet, 0)

	tweetsIds := make([]uint64, 0)

	for _, message := range messages {
		newTweet, id := domain.NewTweet(testUser, message)
		testTweets = append(testTweets, newTweet)
		tweetsIds = append(tweetsIds, id)
	}

	// Operation
	userService.RegisterUser(testUser)

	for i,curTweet := range testTweets {
		err := userService.Tweet(testUser, curTweet)
		// Validation
		if err != nil {
			t.Errorf("Didn't expect any error, but got: %s", err.Error())
		}

		publishedTweet := tweet.GetById(tweetsIds[i])

		if publishedTweet.User != curTweet.User ||
			publishedTweet.Text != curTweet.Text {
			t.Errorf("Expected tweetMessage is %s: %s \nbut is %s: %s",
				curTweet.User, curTweet.Text, publishedTweet.User, publishedTweet.Text)
		}
	}

}
