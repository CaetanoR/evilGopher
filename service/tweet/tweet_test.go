package tweet_test

import (
	"testing"
	"github.com/evilGopher/service/tweet"
	"github.com/evilGopher/service/user"
)

func TestPublishIsSaved(t *testing.T) {

	// Initialization
	var tweetMessage *tweet.Tweet

	testUser := user.New("grupoesfera")
	text := "This is my first tweetMessage"

	tweetMessage = tweet.New(testUser, text)

	// Operation
	user.Create(testUser)
	err := tweet.Publish(tweetMessage)

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

func TestPublishWithoutUserOrTextError(t *testing.T) {

	tweetMessage := tweet.New(user.New(""), "Some tweet")

	err := tweet.Publish(tweetMessage)

	if err == nil {
		t.Error("error was expected")
	}

	if err.Error() != "user is required" {
		t.Errorf("expected error: user is required, but was %s", err.Error())
	}

	tweetMessage = tweet.New(user.New("caetano"), "")

	err = tweet.Publish(tweetMessage)

	if err == nil {
		t.Error("error was expected")
	}

	if err.Error() != "text is required" {
		t.Errorf("expected error: text is required, but was %s", err.Error())
	}

}

func TestPublishWithoutRegisteredUser(t *testing.T) {

	// Initialization
	var tweetMessage *tweet.Tweet

	user := user.New("grupoesfera")
	text := "This is my first tweet"

	tweetMessage = tweet.New(user, text)

	// Operation
	err := tweet.Publish(tweetMessage)

	// Validation
	if err.Error() == "" {
		t.Errorf("expected error")
	}

	if err.Error() != "user must be registered in order to publish tweets" {
		t.Errorf("expected error: user must be registered in order to publish tweets, but was: %s", err.Error())
	}

}
