package views

import (
	"fmt"
	"github.com/diamondburned/arikawa/v3/discord"
)

func Error(message, err string) discord.Embed {
	return discord.Embed{
		Title:       message,
		Description: err,
		Color:       0xFF0000,
	}
}

func Success(message, description string) discord.Embed {
	return discord.Embed{
		Title:       message,
		Description: description,
		Color:       0x00FF00,
	}
}

func Warning(message, description string) discord.Embed {
	return discord.Embed{
		Title:       message,
		Description: description,
		Color:       0xFFFF00,
	}
}

func Forbidden() discord.Embed {
	return Error("Forbidden", "You don't have permission to do that.")
}

func NewEmbeds(embed ...discord.Embed) *[]discord.Embed {
	return &embed
}

func Welcome(member *discord.Member) discord.Embed {
	return discord.Embed{
		Title:       "Welcome",
		Description: fmt.Sprintf("Welcome to the kafejo, %s!\nPlease take a seat and order a coffee ☕.", member.Mention()),
		Color:       0x00FF00,
	}
}
