/*
 *
 * Author: Miltiadis Alexis <alexmiltiadis@gmai.com>
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v6"
	"github.com/miltalex/Mikasa/levi/messages"
)

// Variables used for command line parameters
var (
	Token     string
	ChannelID string
	GuildID   string
)

var buffer = make([][]byte, 0)

// Environment Variables
type environmentVariables struct {
	Token                string `env:BOT_TOKEN,envDefault:"MyToken"`
	DiscordrusWebHookURL string `env:"DISCORDRUS_WEBHOOK_URL"`
	Port                 string `env:"PORT" envDefault:"8081"`
	BotName              string `env:"BOT_NAME" envDefault:"armin"`
	BotKeyword           string `env:"BOT_KEYWORD" envDefault:"!ackerman"`
	InstanceName         string `env:"INSTANCE_NAME" envDefault:"Mikasa-0"`
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&GuildID, "g", "", "Guild in which voice channel exists")
	flag.StringVar(&ChannelID, "c", "", "Voice channel to connect to")
	flag.Parse()
}

func main() {
	envVars := &environmentVariables{}

	err := env.Parse(envVars)
	if err != nil {
		log.Fatalf("Error looking up environment variables: %s", err)
	}

	if envVars.Token == "" {
		if Token == "" {
			log.Fatal("A Token must be provided in order to coonect to discord")
		}
		envVars.Token = Token
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	messages.Init(dg)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Levi is now comming for you.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
