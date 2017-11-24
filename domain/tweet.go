package domain

import "time"

type Tweet struct {

	User string
	Text string
	Date *time.Time
}

func NewTweet(user, text string) *Tweet {
	tiempo := time.Now()
	return &Tweet{user, text, &(tiempo)}
}
