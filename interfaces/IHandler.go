package interfaces

import (
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/state"
)

type IHandler interface {
	AddFunc(name string, fn cmdroute.CommandHandlerFunc)
	GetState() *state.State
}
