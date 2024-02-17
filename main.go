package main

import "github.com/1dayin2/discord-bot/bot"

func main() {
	bot.Start()
	<-make(chan struct{})

}
