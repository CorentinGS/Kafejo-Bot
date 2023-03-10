package commands

import (
	"context"
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"time"
)

var (
	commands = []api.CreateCommandData{{Name: "ping", Description: "Ping!"},
		{Name: "get-version", Description: "Returns the version of the bot"}, {Name: "coffee", Description: "Get a cup of coffee"},
		KarmaCommand}
)

func GetCommands() []api.CreateCommandData {
	return commands
}

func NewHandler(s *state.State) *cmdHandler.HandlerModel {
	h := cmdHandler.CreateHandler(s)
	h.AddFunc("ping", pingCommandHandler())
	h.AddFunc("get-version", cmdVersion())
	h.AddFunc("coffee", CoffeeCommandHandler())
	registerKarma(h)
	return h
}

type Command struct {
	interfaces.IHandler
}

func pingCommandHandler() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return &api.InteractionResponseData{Content: option.NewNullableString("Pong!")}
	}
}

func CoffeeCommandHandler() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		// Make a goroutine to send the message in the queue after 10 seconds to simulate the coffee brewing
		go func() {
			time.Sleep(10 * time.Second)
			_, err := utils.GetSession().SendMessage(data.Event.ChannelID,
				"Your coffee is ready "+data.Event.Member.Mention()+"! :coffee:\nhttps://tenor.com/bHtFm.gif")
			if err != nil {
				return
			}
		}()

		return &api.InteractionResponseData{Content: option.NewNullableString("Your order has been taken, please wait while we brew your coffee!"), Flags: discord.EphemeralMessage}
	}
}
