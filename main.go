package main

import (
	"a2-recreate/src"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

	// Create rate limiter for bulk commands (1 minute cooldown)
	bulkRateLimiter := src.NewRateLimiter(time.Minute)

	dg.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		fmt.Println("Bot is now online!")
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		originalContent := m.Content
		isBulkMode := false

		// Check for "!" prefix
		if strings.HasPrefix(originalContent, "!") {
			isBulkMode = true
			originalContent = strings.TrimPrefix(originalContent, "!")
		}

		if src.IsSimilar(originalContent, "howmanyspacemonke", 0.75) {
			// Check rate limiting for bulk mode
			if isBulkMode {
				rateLimitKey := m.ChannelID // Use channel ID as rate limit key
				if !bulkRateLimiter.CheckAndUpdate(rateLimitKey) {
					remaining := bulkRateLimiter.GetRemainingCooldown(rateLimitKey)
					s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Bulk command is on cooldown. Please wait %v before using it again.", remaining.Round(time.Second)))
					return
				}
			}

			s.ChannelTyping(m.ChannelID)

			var fleets []src.Fleet
			var err error

			if isBulkMode {
				fleets, err = src.GetAllFleetsAllPages()
			} else {
				fleets, err = src.GetAllFleets()
			}

			if err != nil {
				fmt.Printf("Error getting fleets: %v\n", err)
				s.ChannelMessageSend(m.ChannelID, "Error fetching server data. Please try again later.")
				return
			}

			fmt.Println("Fleets grabbed!")

			if isBulkMode {
				// Generate multiple tables for bulk mode
				tables := src.GenerateStationTables(fleets, 16)
				if len(tables) == 0 {
					s.ChannelMessageSend(m.ChannelID, "No servers found.")
					return
				}

				// Send each table as a separate message
				for i, table := range tables {
					codeBlock := fmt.Sprintf("```\n%s\n```", table)
					if i == 0 {
						// Add header to first message
						header := fmt.Sprintf("**Showing all %d servers (page %d of %d):**\n", getTotalStationCount(fleets), i+1, len(tables))
						s.ChannelMessageSend(m.ChannelID, header+codeBlock)
					} else {
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**Page %d of %d:**\n%s", i+1, len(tables), codeBlock))
					}
				}
			} else {
				// Normal mode - single table, limit to 16 stations
				parsedFleets := src.GenerateStationTable(fleets)
				if parsedFleets == "" {
					s.ChannelMessageSend(m.ChannelID, "No servers found.")
					return
				}
				fmt.Println(parsedFleets)
				codeBlock := fmt.Sprintf("```\n%s\n```", parsedFleets)
				s.ChannelMessageSend(m.ChannelID, codeBlock)
			}
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	dg.Close()
}

// Helper function to count total stations across all fleets
func getTotalStationCount(fleets []src.Fleet) int {
	total := 0
	for _, fleet := range fleets {
		total += len(fleet.Stations)
	}
	return total
}
