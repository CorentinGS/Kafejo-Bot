package events

import (
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/rs/zerolog/log"
)

type Member struct {
	interfaces.IHandler
}

func (m Member) GuildMemberAddEvent() func(c *gateway.GuildMemberAddEvent) {
	log.Debug().Msgf("Registering GuildMemberAddEvent")
	return func(c *gateway.GuildMemberAddEvent) {
		log.Debug().Msgf("Member added: %s", c.User.ID)
	}
}

func (m Member) GuildMemberRemoveEvent() func(c *gateway.GuildMemberRemoveEvent) {
	log.Debug().Msgf("Registering GuildMemberRemoveEvent")
	return func(c *gateway.GuildMemberRemoveEvent) {
		log.Debug().Msgf("Member removed: %s", c.User.ID)
	}
}
