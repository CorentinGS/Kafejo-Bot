package events

import (
	"github.com/corentings/kafejo-bot/app/commands/common"
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/rs/zerolog/log"
)

type Message struct {
	interfaces.IHandler
}

func (m Message) MessageDeleteEvent() func(c *gateway.MessageDeleteEvent) {
	log.Debug().Msg("Registering MessageDeleteEvent")
	return func(c *gateway.MessageDeleteEvent) {
		var logEmbed discord.Embed

		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}
		// Get the message from the cache
		message, err := m.IHandler.GetState().Message(c.ChannelID, c.ID)
		if err != nil {
			logEmbed = common.UnknownMessageDeleteLogger(c.ChannelID, c.GuildID, c.ID).ToEmbed()
		} else {
			logEmbed = common.MessageDeleteLogger(message).ToEmbed()
		}
		common.AddEmbedToQueue(logEmbed)
	}
}

func (m Message) MessageUpdateEvent() func(c *gateway.MessageUpdateEvent) {
	log.Debug().Msg("Registering MessageUpdateEvent")
	return func(c *gateway.MessageUpdateEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		message, err := m.IHandler.GetState().Message(c.ChannelID, c.ID)
		if err != nil {
			log.Error().Err(err).Msg("Error getting message from cache")
			return
		}
		logEmbed := common.MessageUpdateLogger(&c.Message, message.Content).ToEmbed()
		common.AddEmbedToQueue(logEmbed)
	}
}

func (m Message) MessageCreateEvent() func(c *gateway.MessageCreateEvent) {
	log.Debug().Msg("Registering MessageCreateEvent")
	return func(c *gateway.MessageCreateEvent) {
		log.Debug().Msgf("Message created: %s", c.ID)
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		if c.Author.Bot || c.Content == "" {
			return
		}

		// Check if the message is a command
		if c.Content[0] == '?' && utils.HasOwnerPermission(c.Message.Author.ID) {
			log.Debug().Msg("Message is a command")
			// Get the command
			//  is the first word after the prefix
			command := c.Content[1:]
			switch command {
			case "join":
				log.Debug().Msg("Join command")
				member, _ := m.IHandler.GetState().Member(c.GuildID, c.Author.ID)
				logEmbed := common.MemberAddLogger(member).ToEmbed()
				common.AddEmbedToQueue(logEmbed)
			case "leave":
				log.Debug().Msg("Leave command")
				member, err := m.IHandler.GetState().Member(c.GuildID, c.Message.Author.ID)
				if err != nil {
					log.Error().Err(err).Msg("Error getting member from cache")
					return
				}

				logEmbed := common.MemberRemoveLogger(&c.Message.Author, member.RoleIDs).ToEmbed()
				common.AddEmbedToQueue(logEmbed)
			}
			return
		}
	}
}

func (m Message) MessageReactionAddEvent() func(c *gateway.MessageReactionAddEvent) {
	log.Debug().Msg("Registering MessageReactionAddEvent")
	return func(c *gateway.MessageReactionAddEvent) {
		log.Debug().Msgf("Message reaction added: %s", c.MessageID)
	}
}
