package events

import (
	"github.com/corentings/kafejo-bot/interfaces"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/rs/zerolog/log"
)

type Message struct {
	interfaces.IHandler
}

func (m Message) MessageDeleteEvent() func(c *gateway.MessageDeleteEvent) {
	log.Debug().Msg("Registering MessageDeleteEvent")
	return func(c *gateway.MessageDeleteEvent) {
		log.Debug().Msgf("Message deleted: %s", c.ID)
	}
}

func (m Message) MessageUpdateEvent() func(c *gateway.MessageUpdateEvent) {
	log.Debug().Msg("Registering MessageUpdateEvent")
	return func(c *gateway.MessageUpdateEvent) {
		log.Debug().Msgf("Message updated: %s", c.ID)
	}
}

func (m Message) MessageCreateEvent() func(c *gateway.MessageCreateEvent) {
	log.Debug().Msg("Registering MessageCreateEvent")
	return func(c *gateway.MessageCreateEvent) {
		log.Debug().Msgf("Message created: %s", c.ID)
	}
}

func (m Message) MessageReactionAddEvent() func(c *gateway.MessageReactionAddEvent) {
	log.Debug().Msg("Registering MessageReactionAddEvent")
	return func(c *gateway.MessageReactionAddEvent) {
		log.Debug().Msgf("Message reaction added: %s", c.MessageID)
	}
}
