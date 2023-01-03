package commands

import (
	"context"
	"github.com/corentings/kafejo-bot/app/commands/version"
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
)

func cmdVersion() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		versionCmd := version.Command{IHandler: cmdHandler.GetHandler()}
		return versionCmd.CmdVersion()
	}
}
