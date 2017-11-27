package tweet

import (
	"time"
	"github.com/evilGopher/service/user"
)

type Tweet struct {

	User *user.User
	Text string
	Date *time.Time

}

func NewTweet(user *user.User, text string) *Tweet {
	tiempo := time.Now()
	return &Tweet{user, text, &(tiempo)}
}