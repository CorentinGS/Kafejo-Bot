package karmacommand

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func GetCommandData() api.CreateCommandData {
	return karmacommand
}

var karmacommand = api.CreateCommandData{
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
