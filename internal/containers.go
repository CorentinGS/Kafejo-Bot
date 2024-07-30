package internal

import (
	"sync"

	"github.com/corentings/kafejo-bot/app/commands/karmacommand"
	"github.com/corentings/kafejo-bot/app/commands/versioncommand"
)

// ServiceContainer is the service container interface.
type ServiceContainer interface {
	GetKarma() karmacommand.Command // GetKarma returns the karma command.
	GetVersion() versioncommand.Command
}

type kernel struct{}

// GetKarma returns the karma command.
func (kernel) GetKarma() karmacommand.Command {
	return InitializeKarma()
}

func (kernel) GetVersion() versioncommand.Command {
	return InitializeVersion()
}

var (
	k             *kernel   // kernel is the service container
	containerOnce sync.Once // containerOnce is used to ensure that the service container is only initialized once
)

// GetServiceContainer returns the service container
func GetServiceContainer() ServiceContainer {
	containerOnce.Do(func() {
		k = &kernel{}
	})
	return k
}
