package bot

import (
	"fmt"
	"strings"

	"github.com/1dayin2/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Can not read bot config")
		return
	}

	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Failed to initial a discord bot")
		return
	}

	u, err := bot.User("@me")
	if err != nil {
		fmt.Println("Failed getting current User:", err)
		return
	}

	BotId = u.ID

	bot.AddHandler(messageHandler)

	err = bot.Open()
	if err != nil {
		fmt.Println("Failed opening connection to Discord:", err)
		return
	}

	fmt.Println("Bot is now connected!")
}

func messageHandler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.ID == BotId {
		return
	}

	config, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Can not read bot config")
		return
	}

	prefix := config.BotPrefix
	if strings.HasPrefix(e.Content, prefix) {
		args := strings.Fields(e.Content)[strings.Index(e.Content, prefix)]
		fmt.Println(args)
		cmd := args[len(prefix):]
		fmt.Println(cmd)
		switch cmd {
		case "ping":
			_, err := s.ChannelMessageSend(e.ChannelID, "Pong!")
			if err != nil {
				fmt.Println("Failed sending Pong response:", err)
			}

		default:
			_, err := s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Unknown command %q.", cmd))
			if err != nil {
				fmt.Println("Failed sending Unknown Command response:", err)
			}
		}
	}

}
