package main

import (
	"github.com/abiosoft/ishell"
	"github.com/evilGopher/service/user"
	"github.com/evilGopher/service/tweet"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("¿Who are you? ")
			currentUser := c.ReadLine()
			c.Print("Write your tweet: ")
			text := c.ReadLine()

			currentTweet := tweet.New(user.New(currentUser), text)
			err := tweet.Publish(currentTweet)

			if err != nil {
				c.Println(err.Error())
				return
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Println(tweet.GetAllAsString())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "removeTweet",
		Help: "Removes a tweet",

		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Tweet to remove: ")

			tweet.Remove(c.ReadLine())
			c.Println(tweet.GetAll())

			return
		},
	})

		shell.Run()

}
