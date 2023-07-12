package commands

import (
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/corentings/kafejo-bot/internal"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
)

var KarmaCommand = api.CreateCommandData{
	Name:        "karmaCmd",
	Description: "Karma main command",
	Options: []discord.CommandOption{
		&discord.SubcommandOption{
			OptionName:  "add",
			Description: "add karmaCmd",
			Required:    false,
			Options: []discord.CommandOptionValue{
				&discord.UserOption{
					OptionName:  "user-option",
					Description: "the user",
					Required:    true,
				},
			},
		},
		&discord.SubcommandOption{
			OptionName:  "show",
			Description: "show karmaCmd",
			Required:    false,
			Options: []discord.CommandOptionValue{
				&discord.UserOption{
					OptionName:  "user-option",
					Description: "the user",
					Required:    false,
				},
			},
		},
	},
}

func registerKarma(h *cmdHandler.HandlerModel) {
	karmaCommand := internal.GetServiceContainer().GetKarma()

	h.Router.Sub("karmaCmd", func(r *cmdroute.Router) {
		r.AddFunc("add", karmaCommand.AddKarmaCommand())
		r.AddFunc("show", karmaCommand.ShowKarmaCommand())
	})

}
