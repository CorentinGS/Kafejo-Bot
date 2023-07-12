//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/corentings/kafejo-bot/app/commands/karmaCmd"
	"github.com/corentings/kafejo-bot/data/infrastructures"
	"github.com/corentings/kafejo-bot/internal/karma"
	"github.com/google/wire"
)

// InitializeUser initializes the user controller.
func InitializeKarma() karmaCmd.Command {
	wire.Build(karmaCmd.NewCommand, karma.NewUseCase, karma.NewSQLRepository, infrastructures.GetDBConn)
	return karmaCmd.Command{}
}
