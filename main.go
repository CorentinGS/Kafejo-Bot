package main

import (
	"context"
	"os"

	"github.com/corentings/kafejo-bot/app/commands"
	"github.com/corentings/kafejo-bot/app/common"
	"github.com/corentings/kafejo-bot/app/events"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() { loadVar() }

func loadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}
}

func main() {
	utils.CreateLogger()

	// Get the token from the .env file
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal().Msg("No token provided")
	}

	go common.CreateSenderWorker()

	h := commands.NewHandler(state.New("Bot " + token))
	events.RegisterHandlers(h)
	h.S.AddInteractionHandler(h)
	h.S.AddIntents(gateway.IntentGuildMessages)
	h.S.AddIntents(gateway.IntentGuildMessageReactions)
	h.S.AddIntents(gateway.IntentGuildMembers)
	h.S.AddHandler(func(event *gateway.ReadyEvent) {
		me, _ := h.S.Me()
		log.Info().Msgf("connected to the gateway as %s", me.Tag())
	})

	if err := cmdroute.OverwriteCommands(h.S, commands.GetCommands()); err != nil {
		log.Fatal().Msgf("cannot update commands: %s", err)
	}

	if err := h.S.Connect(context.TODO()); err != nil {
		log.Fatal().Msgf("cannot connect: %s", err)
	}

	// Create a channel to block the main thread
	<-make(chan struct{})
}
