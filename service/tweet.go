package service

var Tweet string

func PublishTweet(tweet string) string {
	Tweet = tweet
	return tweet
}