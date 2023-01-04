package events

import (
	"fmt"
	"github.com/corentings/kafejo-bot/app/common"
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
