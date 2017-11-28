package tweet

import (
	"time"
	"github.com/evilGopher/domain"
)

var tweets []*domain.Tweet

func GetAllAsString() []string {
	return toString(tweets)
}

func GetAll() []*domain.Tweet {
	return tweets
}

func AddTweet(tweet *domain.Tweet) {
	tweets = append(tweets, tweet)
}

func GetById(id uint64) *domain.Tweet {

	var tweetToFind *domain.Tweet

	for _, curTweet := range tweets {
			if curTweet.Id == id {
				tweetToFind = curTweet
			}
		}

	if tweetToFind != nil {
		return tweetToFind
	}
	return nil
}

func Remove(tweet string) {
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


func toString(tweets []*domain.Tweet) []string {
	var result []string
	for _,tweet := range tweets {
		result = append(result, tweet.User.Name + " said: " + tweet.Text  + " at: " + string(tweet.Date.Format(time.UnixDate)) + "\n")
	}
	return result
}