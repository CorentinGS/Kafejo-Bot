package services

import (
	"github.com/corentings/kafejo-bot/app/commands/karma"
	"sync"
)

type IServiceContainer interface {
	InjectKarmaCommandHandler() karma.Command
}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func GetServiceContainer() IServiceContainer {
	containerOnce.Do(func() {
		k = &kernel{}
	})
	return k
}
