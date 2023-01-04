package events

import (
	common2 "github.com/corentings/kafejo-bot/app/common"
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/rs/zerolog/log"
)

type Member struct {
	interfaces.IHandler
}

func (m Member) GuildMemberAddEvent() func(c *gateway.GuildMemberAddEvent) {
	log.Debug().Msgf("Registering GuildMemberAddEvent")
	return func(c *gateway.GuildMemberAddEvent) {
		if c.GuildID.String() != utils.ConfigGuildID {
			return
		}

		logEmbed := common2.MemberAddLogger(&c.Member).ToEmbed()
		common2.AddEmbedToQueue(common2.MessageItem{
			Embed:   logEmbed,
			Channel: common2.GetLoggerChannel(),
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
			logEmbed = common2.MemberRemoveLogger(&c.User, nil).ToEmbed()
		} else {
			logEmbed = common2.MemberRemoveLogger(&c.User, member.RoleIDs).ToEmbed()
		}

		common2.AddEmbedToQueue(common2.MessageItem{
			Embed:   logEmbed,
			Channel: common2.GetLoggerChannel(),
		})
	}
}
