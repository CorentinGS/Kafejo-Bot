package events

import (
	"fmt"

	"github.com/corentings/kafejo-bot/app/common"
	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/corentings/kafejo-bot/views"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/rs/zerolog/log"
)

type Member struct {
	handler.IHandler
}

func (m Member) GuildMemberBanEvent() func(c *gateway.GuildBanAddEvent) {
	log.Debug().Msgf("Registering GuildMemberBanEvent")
	return func(c *gateway.GuildBanAddEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		logEmbed := common.MemberBanLogger(&c.User).ToEmbed()
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   logEmbed,
			Channel: common.GetLoggerChannel(),
		})
	}
}

func (m Member) GuildMemberUnbanEvent() func(c *gateway.GuildBanRemoveEvent) {
	log.Debug().Msgf("Registering GuildMemberUnbanEvent")
	return func(c *gateway.GuildBanRemoveEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		logEmbed := common.MemberUnbanLogger(&c.User).ToEmbed()
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   logEmbed,
			Channel: common.GetLoggerChannel(),
		})
	}
}

func (m Member) GuildMemberAddEvent() func(c *gateway.GuildMemberAddEvent) {
	log.Debug().Msgf("Registering GuildMemberAddEvent")
	return func(c *gateway.GuildMemberAddEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		logEmbed := common.MemberAddLogger(&c.Member).ToEmbed()
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   logEmbed,
			Channel: common.GetLoggerChannel(),
		})

		flag := common.VerifyMember(&c.Member)
		log.Debug().Msgf("Member %s is verified: %s", c.Member.User.Username, flag.String())
		if flag >= common.MemberDangerLevelMedium {
			common.AddEmbedToQueue(common.MessageItem{
				Embed:   common.DangerMemberLogger(c.Member.User, flag).ToEmbed(),
				Channel: common.GetModChannel(),
				Content: fmt.Sprintf("<@&%s>", utils.ConfigRoleMod),
			})
		}

		welcomeChan, _ := discord.ParseSnowflake(utils.ConfigWelcomeChannelID)

		welcomeEmbed := views.Welcome(&c.Member)
		common.AddEmbedToQueue(common.MessageItem{
			Embed:   welcomeEmbed,
			Channel: discord.ChannelID(welcomeChan),
		})
	}
}

func (m Member) GuildMemberRemoveEvent() func(c *gateway.GuildMemberRemoveEvent) {
	log.Debug().Msgf("Registering GuildMemberRemoveEvent")
	return func(c *gateway.GuildMemberRemoveEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		var logEmbed discord.Embed

		member, err := m.IHandler.GetState().Member(c.GuildID, c.User.ID)
		if err != nil {
			logEmbed = common.MemberRemoveLogger(&c.User, nil).ToEmbed()
		} else {
			logEmbed = common.MemberRemoveLogger(&c.User, member.RoleIDs).ToEmbed()
		}

		common.AddEmbedToQueue(common.MessageItem{
			Embed:   logEmbed,
			Channel: common.GetLoggerChannel(),
		})
	}
}

func (m Member) GuildMemberUpdateEvent() func(c *gateway.GuildMemberUpdateEvent) {
	log.Debug().Msgf("Registering GuildMemberUpdateEvent")
	return func(c *gateway.GuildMemberUpdateEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		member, _ := m.GetState().Member(c.GuildID, c.User.ID)
		newMember := c.RoleIDs

		// Check the difference between the old and new roles
		// If the difference is 1, then it's a role add or remove
		// If the difference is more than 1, then it's a role update
		if len(member.RoleIDs) != len(newMember) {
			if len(member.RoleIDs) > len(newMember) {
				// Role removed
				role := utils.GetRoleDifference(member.RoleIDs, newMember)
				common.AddEmbedToQueue(common.MessageItem{
					Embed:   common.MemberRoleRemoveLogger(c.User, role).ToEmbed(),
					Channel: common.GetLoggerChannel(),
				})
			} else {
				// Role added
				role := utils.GetRoleDifference(newMember, member.RoleIDs)
				common.AddEmbedToQueue(common.MessageItem{
					Embed:   common.MemberRoleAddLogger(c.User, role).ToEmbed(),
					Channel: common.GetLoggerChannel(),
				})
			}
		}
	}
}
