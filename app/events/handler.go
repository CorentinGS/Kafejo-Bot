package events

import (
	"github.com/corentings/kafejo-bot/app/handler"
)

func RegisterHandlers(handler *handler.CommandHandler) {
	message := Message{
		IHandler: handler,
	}
	// Register your handlers here
	handler.AddSyncHandler(message.MessageDeleteEvent())

	handler.AddSyncHandler(message.MessageUpdateEvent())

	handler.AddSyncHandler(message.MessageCreateEvent())

	handler.AddSyncHandler(message.MessageReactionAddEvent())

	member := Member{
		IHandler: handler,
	}

	handler.AddSyncHandler(member.GuildMemberAddEvent())

	handler.AddSyncHandler(member.GuildMemberRemoveEvent())

	handler.AddSyncHandler(member.GuildMemberUpdateEvent())

	handler.AddSyncHandler(member.GuildMemberBanEvent())

	handler.AddSyncHandler(member.GuildMemberUnbanEvent())
}
