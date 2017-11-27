package service_test

import (
	"testing"
	"github.com/evilGopher/service"
	"github.com/evilGopher/domain"
)

func TestPublishTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := domain.User{"grupoesfera"}
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	err := service.PublishTweet(tweet)

	// Validation
	if err != nil {
		t.Errorf("Didn't expect any error, but got: %s", err.Error())
	}

	publishedTweet := service.GetTweets()[0]

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestPublishTweetWithoutUserOrTextError(t *testing.T) {

	tweet := domain.NewTweet(domain.User{""}, "Some tweet")

	err := service.PublishTweet(tweet)

	if err == nil {
		t.Error("error was expected")
	}

	if err.Error() != "user is required" {
		t.Errorf("expected error: user is required, but was %s", err.Error())
	}

	tweet = domain.NewTweet(domain.User{"caetano"}, "")

	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("error was expected")
	}

	if err.Error() != "text is required" {
		t.Errorf("expected error: text is required, but was %s", err.Error())
	}

}
