package karmacommand

import (
	"context"

	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/corentings/kafejo-bot/domain"
	"github.com/corentings/kafejo-bot/internal/karma"
	"github.com/corentings/kafejo-bot/views"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/rs/zerolog/log"
)

func NewCommand(karmaService karma.IUseCase, handler handler.IHandler) Command {
	return Command{karmaService, handler}
}

type Command struct {
	karma.IUseCase
	handler.IHandler
}

func (k *Command) AddKarmaCommand() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		options := data.Options

		if options[0].String() == data.Event.Member.User.ID.String() {
			log.Debug().Msgf("User %v tried to add karma to himself", data.Event.Member.User.ID.String())
			return &api.InteractionResponseData{
				Embeds: views.NewEmbeds(views.Forbidden()),
				Flags:  discord.EphemeralMessage,
			}
		}

		karma, err := k.IncrementKarma(ctx, options[0].String(), data.Event.GuildID.String())
		if err != nil {
			log.Warn().Msgf("Error incrementing karma: %v", err)
			return &api.InteractionResponseData{
				Embeds: views.NewEmbeds(
					views.Error("Error incrementing karma", err.Error())),
				Flags: discord.EphemeralMessage,
			}
		}

		return &api.InteractionResponseData{
			Embeds: views.NewEmbeds(
				views.Success("Karma added", karma.GetKarmaAsString())),
			Flags: discord.EphemeralMessage,
		}
	}
}

func (k *Command) ShowKarmaCommand() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		options := data.Options

		var userID string

		// Check if options is empty
		if len(options) == 0 {
			userID = data.Event.Member.User.ID.String()
		} else {
			userID = options[0].String()
		}

		karma, err := k.GetKarma(ctx, userID, data.Event.GuildID.String())
		if err != nil {
			karma, err = k.CreateKarma(ctx, domain.Karma{UserID: userID, GuildID: data.Event.GuildID.String(), Value: 0})
			if err != nil {
				log.Warn().Msgf("Error creating karma: %v", err)
				return &api.InteractionResponseData{
					Embeds: views.NewEmbeds(
						views.Error("Error creating karma", err.Error())),
					Flags: discord.EphemeralMessage,
				}
			}
		}

		return &api.InteractionResponseData{Content: option.NewNullableString(karma.GetKarmaAsString()), Flags: discord.EphemeralMessage}
	}
}

func (k *Command) RemoveKarmaCommand() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		options := data.Options

		log.Debug().Msgf("KarmaCommandHandler: %v", options)

		return &api.InteractionResponseData{Content: option.NewNullableString("totoro"), Flags: discord.EphemeralMessage}
	}
}

func (k *Command) ListKarmaCommand() cmdroute.CommandHandlerFunc {
	return func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		options := data.Options

		log.Debug().Msgf("KarmaCommandHandler: %v", options)

		return &api.InteractionResponseData{Content: option.NewNullableString("totoro"), Flags: discord.EphemeralMessage}
	}
}
