package versioncommand

import (
	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

func NewCommand(handler handler.IHandler) Command {
	return Command{
		handler,
	}
}

type Command struct {
	handler.IHandler
}

func (c Command) GetCommandName() string {
	return CommandName
}

func (c Command) RespondVersion() *api.InteractionResponseData {
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
