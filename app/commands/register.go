package commands

import (
	"context"

	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/corentings/kafejo-bot/internal"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
)

func registerVersion(h *handler.CommandHandler) {
	versionCommand := internal.GetServiceContainer().GetVersion()

	h.AddFunc(versionCommand.GetCommandName(), interactionToCommandHandlerFunc(versionCommand.RespondVersion()))
}

func interactionToCommandHandlerFunc(interaction *api.InteractionResponseData) cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return interaction
	}
}
