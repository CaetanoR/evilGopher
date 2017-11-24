package service

import (
	"fmt"
	"github.com/evilGopher/domain"
	"time"
)

var tweets []*domain.Tweet

func PublishTweet(tweet *domain.Tweet) {
	tweets = append(tweets, tweet)
}

func GetTweetsAsString() []string {
	return tweetToString(tweets)
}

func GetTweets() []*domain.Tweet {
	return tweets
}

func RemoveTweet(tweet string) {
	defer func() {
		if recover() != nil {
			fmt.Println("array index out of bounds")
		}
	}()
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
		result = append(result, tweet.User + " said: " + tweet.Text  + " at: " + string(tweet.Date.Format(time.UnixDate)) + "\n")
	}
	return result
}