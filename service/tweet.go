package service

import (
	"errors"
	"github.com/evilGopher/domain"
	"time"
)

var tweets []*domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {

	if tweet.User.Name == "" {
		return errors.New("user is required")
	}

	if tweet.Text == "" {
		return errors.New("text is required")
	}

	tweets = append(tweets, tweet)
	return nil
}

func GetTweetsAsString() []string {
	return tweetToString(tweets)
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func RemoveTweet(tweet string) {
	tweets = removeIndex(tweets, tweet)
}

func removeIndex(s []*domain.Tweet, removeTweet string) []*domain.Tweet {

	for i,tweet := range tweets {

		if tweet.Text == removeTweet {
			return append(s[:i], s[i+1:]...)
		}
	}
	return tweets
}


func tweetToString (tweets []*domain.Tweet) []string {
	var result []string
	for _,tweet := range tweets {
		result = append(result, tweet.User.Name + " said: " + tweet.Text  + " at: " + string(tweet.Date.Format(time.UnixDate)) + "\n")
	}
	return result
}