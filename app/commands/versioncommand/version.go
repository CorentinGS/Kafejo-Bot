package versioncommand

import (
	"github.com/diamondburned/arikawa/v3/api"
)

const (
	// CommandName is the name of the command
	CommandName = "version"
)

func GetCommandData() api.CreateCommandData {
	return versionCommand
}

var versionCommand = api.CreateCommandData{
	Name:        CommandName,
	Description: "Returns the version of the bot",
}
