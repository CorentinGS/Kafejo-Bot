package main

import (
	"context"
	"github.com/corentings/kafejo-bot/app/commands"
	"github.com/corentings/kafejo-bot/app/common"
	"github.com/corentings/kafejo-bot/app/events"
	"github.com/corentings/kafejo-bot/data/infrastructures"
	"github.com/corentings/kafejo-bot/models"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
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

	dbConfig := infrastructures.DBConfig()
	err := dbConfig.Connect()
	if err != nil {
		log.Panic().Err(err).Msg("Error connecting to database")
	}

	// Models to migrate
	var migrates []interface{}
	migrates = append(migrates, models.Karma{})

	// AutoMigrate models
	for i := 0; i < len(migrates); i++ {
		err = infrastructures.GetDBConn().AutoMigrate(&migrates[i])
		if err != nil {
			log.Panic().Err(err).Msg("Can't auto migrate models")
		}
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

	if err = cmdroute.OverwriteCommands(h.S, commands.GetCommands()); err != nil {
		log.Fatal().Msgf("cannot update commands: %s", err)
	}

	if err = h.S.Connect(context.TODO()); err != nil {
		log.Fatal().Msgf("cannot connect: %s", err)
	}

	// Create a channel to block the main thread
	<-make(chan struct{})
}
