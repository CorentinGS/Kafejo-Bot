package version

import (
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

type Command struct {
	interfaces.IHandler
}

func (c Command) CmdVersion() *api.InteractionResponseData {
	me, _ := c.GetState().Me()
	return &api.InteractionResponseData{Embeds: &[]discord.Embed{
		{
			Title:       "Version",
			Description: "The current version of the bot is: " + utils.VERSION,
			Author: &discord.EmbedAuthor{
				Name: me.Username,
				URL:  utils.GITHUB,
				Icon: me.AvatarURL(),
			},
		},
	}}
}
