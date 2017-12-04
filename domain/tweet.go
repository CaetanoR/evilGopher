package domain

import (
	"time"
	"sync/atomic"
)

var tweetId uint64

type Tweet struct {

	Id uint64
	User *User
	Text string
	Date time.Time

}

func NewTweet(user *User, text string) (*Tweet, uint64) {
	tiempo := time.Now()
	atomic.AddUint64(&tweetId, 1)
	tweet := &Tweet{tweetId,user, text, tiempo}
	return tweet, tweet.Id
}
