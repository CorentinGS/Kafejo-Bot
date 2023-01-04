package events

import (
	"github.com/corentings/kafejo-bot/app/commands/common"
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/corentings/kafejo-bot/views"
	"github.com/diamondburned/arikawa/v3/api"
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
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   logEmbed,
			Channel: common.GetLoggerChannel(),
		})
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
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   logEmbed,
			Channel: common.GetLoggerChannel(),
		})
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
				common.AddEmbedToQueue(common.MessageItem{
					Embed:   logEmbed,
					Channel: common.GetLoggerChannel(),
				})
			case "leave":
				log.Debug().Msg("Leave command")
				member, err := m.IHandler.GetState().Member(c.GuildID, c.Message.Author.ID)
				if err != nil {
					log.Error().Err(err).Msg("Error getting member from cache")
					return
				}

				logEmbed := common.MemberRemoveLogger(&c.Message.Author, member.RoleIDs).ToEmbed()
				common.AddEmbedToQueue(common.MessageItem{
					Embed:   logEmbed,
					Channel: common.GetLoggerChannel(),
				})
			}
			return
		}
	}
}

func (m Message) MessageReactionAddEvent() func(c *gateway.MessageReactionAddEvent) {
	log.Debug().Msg("Registering MessageReactionAddEvent")
	return func(c *gateway.MessageReactionAddEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}
		if c.ChannelID.String() != utils.ConfigGateKeepChannelID || c.MessageID.String() != utils.ConfigWelcomeMessageID {
			return
		}
		if c.Emoji.Name != "☕" {
			return
		}

		log.Debug().Msgf("Message reaction added: %s", c.MessageID)

		// Convert roleID to snowflake
		roleID, err := discord.ParseSnowflake(utils.ConfigMainRole)

		err = m.GetState().AddRole(c.GuildID, c.UserID, discord.RoleID(roleID), api.AddRoleData{AuditLogReason: "User reacted to welcome message"})
		if err != nil {
			log.Error().Err(err).Msg("Error adding role to user")
			return
		}

		log.Debug().Msgf("Added role %s to user %s", roleID, c.UserID)

		// Remove the reaction
		err = m.GetState().DeleteUserReaction(c.ChannelID, c.MessageID, c.UserID, discord.APIEmoji(c.Emoji.Name))
		if err != nil {
			log.Error().Err(err).Msg("Error removing reaction")
			return
		}

		welcomeChan, err := discord.ParseSnowflake(utils.ConfigWelcomeChannelID)

		welcomeEmbed := views.Welcome(c.Member)
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   welcomeEmbed,
			Channel: discord.ChannelID(welcomeChan),
		})
	}
}
