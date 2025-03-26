//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/corentings/kafejo-bot/app/commands/versioncommand"
	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/google/wire"
)

func InitializeVersion() versioncommand.Command {
	wire.Build(versioncommand.NewCommand, handler.GetHandler)
	return versioncommand.Command{}
}
