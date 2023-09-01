package main

import (
	"fmt"
	"log"
	"time"

	"github.com/techygrrrl/go-twitch-irc"
)

const (
	/// The username of the user that will be saying the message
	chatUsername = "__"

	/// The OAuth access token of the chatUsername user (`oauth:` prefix required)
	chatAccessToken = "oauth:__"

	/// The message to say on a timer
	chatMessage = "Hello, world! 🌎"

	/// The channel to connect to
	chatChannel = "__"

	/// The repeat interval, e.g. 60 means the message will send every 60 seconds
	chatIntervalSeconds = 180 // 3 minutes
)

func main() {
	fmt.Println("Initializing...")

	client := twitch.NewClient(chatUsername, chatAccessToken)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		//fmt.Println("OnPrivateMessage")
		//fmt.Println(message)
	})

	client.OnSelfJoinMessage(func(message twitch.UserJoinMessage) {
		fmt.Println("OnSelfJoinMessage")

		go func() {
			for {
				fmt.Printf("Sending: channel = %s , message = '%s'\n", chatChannel, chatMessage)

				client.Say(chatChannel, chatMessage)
				time.Sleep(chatIntervalSeconds * time.Second)
			}
		}()
	})

	client.OnConnect(func() {
		fmt.Println("OnConnect")
	})

	client.Join(chatChannel)

	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
