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

	testUser := domain.NewUser("grupoesfera", &userService)
	text := "This is my first tweetMessage"

	tweetMessage,_ = domain.NewTweet(testUser, text)

	// Operation
	userService.RegisterUser(testUser)
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

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestPublishWithoutTextError(t *testing.T) {

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	testUser := domain.NewUser("caetano", &userService)

	tweetMessage,_ := domain.NewTweet(testUser, "")

	err := userService.Tweet(testUser, tweetMessage)

	if err == nil {
		t.Error("error was expected")
	}

	if err.Error() != "text is required" {
		t.Errorf("expected error: text is required, but was %s", err.Error())
	}

}

func TestPublishWithoutRegisteredUser(t *testing.T) {

	// Initialization
	var tweetMessage *domain.Tweet
	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	testUser := domain.NewUser("grupoesfera", &userService)
	text := "This is my first tweet"

	tweetMessage,_ = domain.NewTweet(testUser, text)

	// Operation
	err := userService.Tweet(testUser, tweetMessage)

	// Validation
	if err.Error() == "" {
		t.Errorf("expected error")
	}

	if err.Error() != "user must be registered in order to publish tweets" {
		t.Errorf("expected error: testUser must be registered in order to publish tweets, but was: %s", err.Error())
	}

}

func TestGetById(t *testing.T) {

	// Initialization

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	testUser := domain.NewUser("grupoesfera", &userService)

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

func TestFollowUser(t *testing.T) {
	// Initialization

	var userService user.Service

	userService.Initialize()
	tweet.Initialize()

	testUser := domain.NewUser("grupoesfera", &userService)
	testUserToFollow := domain.NewUser("grupoesfera", &userService)

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
