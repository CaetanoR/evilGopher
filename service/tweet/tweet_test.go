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

	var service user.Service

	service.Initialize()

	testUser := domain.NewUser("grupoesfera", &service)
	text := "This is my first tweetMessage"

	tweetMessage = domain.NewTweet(testUser, text)

	// Operation
	service.RegisterUser(testUser)
	err := service.Tweet(testUser, tweetMessage)

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

	var service user.Service

	testUser := domain.NewUser("caetano", &service)

	tweetMessage := domain.NewTweet(testUser, "")

	err := service.Tweet(testUser, tweetMessage)

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
	var service user.Service

	testUser := domain.NewUser("grupoesfera", &service)
	text := "This is my first tweet"

	tweetMessage = domain.NewTweet(testUser, text)

	// Operation
	err := service.Tweet(testUser, tweetMessage)

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

	var service user.Service

	testUser := domain.NewUser("grupoesfera", &service)

	messages := []string{"This is my first tweetMessage", "This is my second tweetMessage", "This is my third tweetMessage"}

	var tweets []*domain.Tweet

	for _, message := range messages {
		tweets = append(tweets, domain.NewTweet(testUser, message))
	}

	// Operation
	service.RegisterUser(testUser)

	for i,curTweet := range tweets {
		err := service.Tweet(testUser, curTweet)
		// Validation
		if err != nil {
			t.Errorf("Didn't expect any error, but got: %s", err.Error())
		}

		publishedTweet := tweet.GetById(uint64(i+1))

		if publishedTweet.User != curTweet.User ||
			publishedTweet.Text != curTweet.Text {
			t.Errorf("Expected tweetMessage is %s: %s \nbut is %s: %s",
				curTweet.User, curTweet.Text, publishedTweet.User, publishedTweet.Text)
		}
	}

}