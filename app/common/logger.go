package common

import (
	"fmt"
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/discord"
	"strings"
)

// LogType is the type of the log
type LogType int64

const (
	// LogTypeMessageDelete is a delete message log
	LogTypeMessageDelete LogType = iota
	// LogTypeMessageUpdate is an update message log
	LogTypeMessageUpdate
	// LogTypeMessageCreate is a creation message log
	LogTypeMessageCreate
	// LogTypeMessageReactionAdd is an add reaction log
	LogTypeMessageReactionAdd
	// LogTypeGuildMemberAdd is an add member log
	LogTypeGuildMemberAdd
	// LogTypeGuildMemberRemove is a remove member log
	LogTypeGuildMemberRemove
)

func (l LogType) String() string {
	switch l {
	case LogTypeMessageDelete:
		return "Message deleted"
	case LogTypeMessageUpdate:
		return "Message updated"
	case LogTypeMessageCreate:
		return "Message created"
	case LogTypeMessageReactionAdd:
		return "Message reaction added"
	case LogTypeGuildMemberAdd:
		return "Guild member added"
	case LogTypeGuildMemberRemove:
		return "Guild member removed"
	default:
		return "Unknown"
	}
}

func (l LogType) Color() int {
	switch l {
	case LogTypeMessageDelete:
		return 0xFF0000
	case LogTypeMessageUpdate:
		return 0xFFFF00
	case LogTypeMessageCreate:
		return 0x00FF00
	case LogTypeMessageReactionAdd:
		return 0x0000FF
	case LogTypeGuildMemberAdd:
		return 0x00FFFF
	case LogTypeGuildMemberRemove:
		return 0xFF00FF
	default:
		return 0x000000
	}
}

// Logger is a logger
type Logger struct {
	Type    LogType
	Author  discord.User
	Message string
	Footer  *discord.EmbedFooter
}

func (l Logger) ToEmbed() discord.Embed {
	return createLoggerEmbed(l)
}

func createLoggerEmbed(logger Logger) discord.Embed {
	return discord.Embed{
		Title:       logger.Type.String(),
		Description: logger.Message,
		Color:       discord.Color(logger.Type.Color()),
		Author:      &discord.EmbedAuthor{Name: logger.Author.Username, Icon: logger.Author.AvatarURL()},
		Footer:      logger.Footer,
		Timestamp:   discord.NowTimestamp(),
	}
}

func MessageDeleteLogger(message *discord.Message) Logger {
	return Logger{
		Type:    LogTypeMessageDelete,
		Author:  message.Author,
		Message: fmt.Sprintf("Channel: <#%s>\nContent: **%s**", message.ChannelID, message.Content),
		Footer:  &discord.EmbedFooter{Text: "Message ID: " + message.ID.String()},
	}
}

func UnknownMessageDeleteLogger(channelID discord.ChannelID, guildID discord.GuildID, messageID discord.MessageID) Logger {
	return Logger{
		Type:    LogTypeMessageDelete,
		Author:  discord.User{},
		Message: fmt.Sprintf("Channel: <#%s>\nContent: **Unknown**", channelID),
		Footer:  &discord.EmbedFooter{Text: "Message ID: " + messageID.String()},
	}
}

func MessageUpdateLogger(message *discord.Message, oldContent string) Logger {
	return Logger{
		Type:    LogTypeMessageUpdate,
		Message: fmt.Sprintf("Channel: <#%s>\nOld content: %s\nNew content: %s", message.ChannelID, oldContent, message.Content),
		Footer:  &discord.EmbedFooter{Text: "Message ID: " + message.ID.String()},
	}
}

func MemberAddLogger(member *discord.Member) Logger {
	return Logger{
		Type: LogTypeGuildMemberAdd,
		Message: fmt.Sprintf("Member: %s (%s)\nAccount created: %s ago (%s)",
			member.Mention(),
			member.User.Tag(),
			utils.FormatTimeSince(member.User.CreatedAt()),
			member.User.CreatedAt().Format("2006-01-02 15:04:05")),
		Footer: &discord.EmbedFooter{Text: "Member ID: " + member.User.ID.String()},
		Author: member.User,
	}
}

func MemberRemoveLogger(user *discord.User, roles []discord.RoleID) Logger {
	// Create a string builder
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Member: %s", user.Tag()))
	// Get roles
	if roles != nil {
		sb.WriteString("\n\nRoles: ")
		for _, role := range roles {
			sb.WriteString(fmt.Sprintf("<@&%s> ", role))
		}
	}
	return Logger{
		Type:    LogTypeGuildMemberRemove,
		Message: sb.String(),
		Footer:  &discord.EmbedFooter{Text: "Member ID: " + user.ID.String()},
		Author:  *user,
	}
}

func CreateLoggerEmbeds(loggers []Logger) *[]discord.Embed {
	var embeds []discord.Embed
	for _, logger := range loggers {
		embeds = append(embeds, createLoggerEmbed(logger))
	}
	return &embeds
}

func GetLoggerChannel() discord.ChannelID {
	// Convert string to snowflake
	channelID, _ := discord.ParseSnowflake(utils.ConfigLogsChannelID)
	return discord.ChannelID(channelID)
}
