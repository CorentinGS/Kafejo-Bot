package commands

import (
	"context"
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

var (
	commands = []api.CreateCommandData{{Name: "ping", Description: "Ping!"},
		{Name: "get-version", Description: "Returns the version of the bot"},
		KarmaCommand}
)

func GetCommands() []api.CreateCommandData {
	return commands
}

func NewHandler(s *state.State) *cmdHandler.HandlerModel {
	h := cmdHandler.CreateHandler(s)
	h.AddFunc("ping", pingCommandHandler())
	h.AddFunc("get-version", cmdVersion())
	registerKarma(h)
	return h
}

type Command struct {
	interfaces.IHandler
}

func pingCommandHandler() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return &api.InteractionResponseData{Content: option.NewNullableString("Pong!")}
	}
}
