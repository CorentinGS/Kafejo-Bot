package internal

import (
	"github.com/corentings/kafejo-bot/app/commands/karmaCmd"
	"sync"
)

// ServiceContainer is the service container interface.
type ServiceContainer interface {
	GetKarma() karmaCmd.Command // GetKarma returns the karma command.
}

type kernel struct{}

// GetUser returns the user controller.
func (kernel) GetKarma() karmaCmd.Command {
	return InitializeKarma()
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
