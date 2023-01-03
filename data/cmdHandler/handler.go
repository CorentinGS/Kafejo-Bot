package cmdHandler

import (
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/state"
	arikawaHandler "github.com/diamondburned/arikawa/v3/utils/handler"
)

var handler *HandlerModel

type HandlerModel struct {
	*cmdroute.Router
	S *state.State
}

func CreateHandler(s *state.State) *HandlerModel {
	handler = &HandlerModel{
		Router: cmdroute.NewRouter(),
		S:      s,
	}
	handler.Use(cmdroute.Deferrable(s, cmdroute.DeferOpts{}))
	handler.Router.Use(cmdroute.UseContext(s.Context()))

	handler.S.PreHandler = arikawaHandler.New()

	return handler
}

func (h *HandlerModel) AddFunc(name string, fn cmdroute.CommandHandlerFunc) {
	h.Router.AddFunc(name, fn)
}

func (h *HandlerModel) AddSyncHandler(fn interface{}) {
	h.S.PreHandler.AddSyncHandler(fn)
}

func (h *HandlerModel) AddHandler(fn interface{}) {
	h.S.PreHandler.AddHandler(fn)
}

func (h *HandlerModel) GetState() *state.State {
	return h.S
}

func GetHandler() *HandlerModel {
	return handler
}
