package tweet

import (
	"errors"
	"time"
	"github.com/evilGopher/service/user"
)

var tweets []*Tweet


func Publish(tweet *Tweet) error {



	if tweet.User.Name == "" {
		return errors.New("user is required")
	}

	if tweet.Text == "" {
		return errors.New("text is required")
	}

	if !user.Exists(tweet.User) {
		return errors.New("user must be registered in order to publish tweets")
	}

	tweets = append(tweets, tweet)
	return nil
}

func GetAllAsString() []string {
	return toString(tweets)
}

func GetAll() []*Tweet {
	return tweets
}

func Remove(tweet string) {
	tweets = removeIndex(tweets, tweet)
}

func removeIndex(s []*Tweet, removeTweet string) []*Tweet {

	for i,tweet := range tweets {

		if tweet.Text == removeTweet {
			return append(s[:i], s[i+1:]...)
		}
	}
	return tweets
}


func toString(tweets []*Tweet) []string {
	var result []string
	for _,tweet := range tweets {
		result = append(result, tweet.User.Name + " said: " + tweet.Text  + " at: " + string(tweet.Date.Format(time.UnixDate)) + "\n")
	}
	return result
}