package service_test

import (
		"testing"
		"github.com/evilGopher/service"
		)


func TestHelloWorld(t *testing.T) {
	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is", tweet)
	}

}
