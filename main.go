package main

import (
	"a2-recreate/src"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Beginning init")

	err := godotenv.Load()

	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Println("Error loading .env file!")
	// }

	token := os.Getenv("TOKEN")

	if token == "" {
		fmt.Println("Could not find token --> Set TOKEN in .env file.")
		return
	}

	fmt.Println("Connecting to bot with token", token)

	dg, err := discordgo.New("Bot " + token)
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}
	fmt.Println("Connected successfully!")

	if err != nil {
		fmt.Println("Error creating Discord session via bot,", err)
		return
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	dg.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		fmt.Println("Bot is now online!")
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == "howmanyspacemonke" {
			fleets, err := src.GetAllFleets()

			if err != nil {
				fmt.Println("Error getting fleets %v", err)
				return
			}

			fmt.Println("Fleets grabbed!")

			parsedFleets := src.GenerateStationTable(fleets)

			fmt.Println(parsedFleets)

			codeBlock := fmt.Sprintf("```\n%s\n```", parsedFleets)

			// Send it to the channel
			s.ChannelMessageSend(m.ChannelID, codeBlock)

		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	dg.Close()
}
