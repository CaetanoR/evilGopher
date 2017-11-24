package service_test

import (
	"testing"
	"github.com/evilGopher/service"
	"github.com/evilGopher/domain"
)

func TestPublishTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
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
