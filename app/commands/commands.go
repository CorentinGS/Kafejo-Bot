package commands

import (
	"context"
	"math/rand"
	"time"

	"github.com/corentings/kafejo-bot/app/commands/versioncommand"
	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

var commands = []api.CreateCommandData{
	{Name: "ping", Description: "Ping!"},
	{Name: "coffee", Description: "Get a cup of coffee"},
	{Name: "matcha", Description: "Get a cup of matcha"},
	{Name: "beer", Description: "Get a cup of beer"},
	versioncommand.GetCommandData(),
}

func GetCommands() []api.CreateCommandData {
	return commands
}

func NewHandler(s *state.State) *handler.CommandHandler {
	h := handler.CreateHandler(s)
	h.AddFunc("ping", pingCommandHandler())
	h.AddFunc("coffee", CoffeeCommandHandler())
	h.AddFunc("matcha", MatchaCommandHandler())
	h.AddFunc("beer", BeerCommandHandler())

	registerVersion(h)
	return h
}

type Command struct {
	handler.IHandler
}

func pingCommandHandler() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return &api.InteractionResponseData{Content: option.NewNullableString("Pong!")}
	}
}

// GetRandomInt generates a random number between min and max
func GetRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func BeerCommandHandler() cmdroute.CommandHandlerFunc {
	var beers = []string{"https://tenor.com/view/schwanchi-gif-10646964323290635570", "https://tenor.com/view/beer-garden-beer-waitress-gif-14167773785568845141"}
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		go func() {
			time.Sleep(10 * time.Second)
			// get a random beer gif
			beergif := beers[GetRandomInt(0, len(beers))]
			_, err := utils.GetSession().SendMessage(data.Event.ChannelID,
				"Your beer is ready "+data.Event.Member.Mention()+"! :beers:\n"+beergif)
			if err != nil {
				return
			}
		}()
		return &api.InteractionResponseData{Content: option.NewNullableString("Your order has been taken, please wait while we pour your beer!"), Flags: discord.EphemeralMessage}
	}
}

func MatchaCommandHandler() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		// Make a goroutine to send the message in the queue after 10 seconds to simulate the matcha brewing
		go func() {
			time.Sleep(10 * time.Second)
			_, err := utils.GetSession().SendMessage(data.Event.ChannelID,
				"Your matcha is ready "+data.Event.Member.Mention()+"! :tea:\nhttps://tenor.com/view/tea-matcha-green-tea-stir-matcha-green-tea-gif-17550598")
			if err != nil {
				return
			}
		}()

		return &api.InteractionResponseData{Content: option.NewNullableString("Your order has been taken, please wait while we brew your matcha!"), Flags: discord.EphemeralMessage}
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
