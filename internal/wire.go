//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/corentings/kafejo-bot/app/commands/karmacommand"
	"github.com/corentings/kafejo-bot/app/commands/versioncommand"
	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/corentings/kafejo-bot/infrastructures"
	"github.com/corentings/kafejo-bot/internal/karma"
	"github.com/google/wire"
)

// InitializeUser initializes the user controller.
func InitializeKarma() karmacommand.Command {
	wire.Build(karmacommand.NewCommand, handler.GetHandler, karma.NewUseCase, karma.NewSQLRepository, infrastructures.GetDBConn)
	return karmacommand.Command{}
}

func InitializeVersion() versioncommand.Command {
	wire.Build(versioncommand.NewCommand, handler.GetHandler)
	return versioncommand.Command{}
}
