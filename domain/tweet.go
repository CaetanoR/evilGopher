package domain

import "time"

type Tweet struct {

	User User
	Text string
	Date *time.Time
}

type User struct {

	Name string

}

func NewTweet(user User, text string) *Tweet {
	tiempo := time.Now()
	return &Tweet{user, text, &(tiempo)}
}
