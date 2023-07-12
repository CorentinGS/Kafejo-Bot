package handler

import (
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/state"
	arikawaHandler "github.com/diamondburned/arikawa/v3/utils/handler"
)

type IHandler interface {
	AddFunc(name string, fn cmdroute.CommandHandlerFunc)
	GetState() *state.State
}

var commandHandler *CommandHandler

type CommandHandler struct {
	*cmdroute.Router
	S *state.State
}

func CreateHandler(s *state.State) *CommandHandler {
	commandHandler = &CommandHandler{
		Router: cmdroute.NewRouter(),
		S:      s,
	}
	commandHandler.Use(cmdroute.Deferrable(s, cmdroute.DeferOpts{}))
	commandHandler.Router.Use(cmdroute.UseContext(s.Context()))

	commandHandler.S.PreHandler = arikawaHandler.New()

	return commandHandler
}

func (h *CommandHandler) AddFunc(name string, fn cmdroute.CommandHandlerFunc) {
	h.Router.AddFunc(name, fn)
}

func (h *CommandHandler) AddSyncHandler(fn interface{}) {
	h.S.PreHandler.AddSyncHandler(fn)
}

func (h *CommandHandler) AddHandler(fn interface{}) {
	h.S.PreHandler.AddHandler(fn)
}

func (h *CommandHandler) GetState() *state.State {
	return h.S
}

func GetHandler() IHandler {
	return commandHandler
}
