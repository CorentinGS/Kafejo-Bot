package commands

import (
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/corentings/kafejo-bot/services"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
)

var KarmaCommand = api.CreateCommandData{
	Name:        "karma",
	Description: "Karma main command",
	Options: []discord.CommandOption{
		&discord.SubcommandOption{
			OptionName:  "add",
			Description: "add karma",
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
			Description: "show karma",
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
	karmaCommand := services.GetServiceContainer().InjectKarmaCommandHandler()

	h.Router.Sub("karma", func(r *cmdroute.Router) {
		r.AddFunc("add", karmaCommand.AddKarmaCommand())
		r.AddFunc("show", karmaCommand.ShowKarmaCommand())
	})

}
