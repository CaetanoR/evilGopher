package main

import (
	"github.com/abiosoft/ishell"
	"github.com/evilGopher/service"
	"github.com/evilGopher/domain"
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
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			text := c.ReadLine()

			tweet := domain.NewTweet(domain.User{user}, text)
			err := service.PublishTweet(tweet)

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

			c.Println(service.GetTweetsAsString())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "removeTweet",
		Help: "Removes a tweet",

		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Tweet to remove: ")

			service.RemoveTweet(c.ReadLine())
			c.Println(service.GetTweets())

			return
		},
	})

		shell.Run()

}
