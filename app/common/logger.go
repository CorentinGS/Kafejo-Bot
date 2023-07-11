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
	// LogDangerousMemberAdd is a dangerous member log
	LogDangerousMemberAdd
	// LogTypeMemberRoleAdd is a member role add log
	LogTypeMemberRoleAdd
	// LogTypeMemberRoleRemove is a member role remove log
	LogTypeMemberRoleRemove
	// LogTypeGuildMemberBan is a member ban log
	LogTypeGuildMemberBan
	// LogTypeGuildMemberUnban is a member unban log
	LogTypeGuildMemberUnban
)

func (l LogType) String() string {
	return [...]string{
		"Message deleted",
		"Message updated",
		"Message created",
		"Message reaction added",
		"Guild member added",
		"Guild member removed",
		"Danger Will Robinson 🤖",
		"Member role added",
		"Member role removed",
		"Member banned",
		"Member unbanned",
	}[l]
}

func (l LogType) Color() int {

	return [...]int{
		0xFF0000, // red (message delete)
		0xFFFF00, // yellow (message update)
		0x00FF00, // green (message create)
		0x0000FF, // blue (message reaction add)
		0x00FFFF, // cyan (guild member add)
		0xFF00FF, // magenta (guild member remove)
		0xFF0000, // red (dangerous member add)
		0x00FF00, // green (member role add)
		0xFF0000, // red 	(member role remove)
		0xFF0000, // red (member ban)
		0x00FF00, // green (member unban)
	}[l]

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

func MemberBanLogger(user *discord.User) Logger {
	return Logger{
		Type:    LogTypeGuildMemberBan,
		Message: fmt.Sprintf("Member: %s", user.Tag()),
		Footer:  &discord.EmbedFooter{Text: "Member ID: " + user.ID.String()},
	}
}

func MemberUnbanLogger(user *discord.User) Logger {
	return Logger{
		Type:    LogTypeGuildMemberUnban,
		Message: fmt.Sprintf("Member: %s", user.Tag()),
		Footer:  &discord.EmbedFooter{Text: "Member ID: " + user.ID.String()},
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

func GetModChannel() discord.ChannelID {
	channelID, _ := discord.ParseSnowflake(utils.ConfigAdminChannelID)
	return discord.ChannelID(channelID)
}

func DangerMemberLogger(member discord.User, danger MemberDangerLevel) Logger {
	return Logger{
		Type:   LogDangerousMemberAdd,
		Author: member,
		Footer: &discord.EmbedFooter{Text: "Member ID: " + member.ID.String()},
		Message: fmt.Sprintf("A suspicious member has joined: %s\n\nDanger level: **%s**\nAccount created: %s ago (%s)\n\n",
			member.Tag(), danger.String(), utils.FormatTimeSince(member.CreatedAt()),
			member.CreatedAt().Format("2006-01-02 15:04:05"),
		),
	}
}

func MemberRoleAddLogger(member discord.User, role discord.RoleID) Logger {
	return Logger{
		Type:   LogTypeMemberRoleAdd,
		Author: member,
		Footer: &discord.EmbedFooter{Text: "Role ID: " + role.String()},
		Message: fmt.Sprintf("A role has been added to a member: %s\n\nRole: <@&%s>",
			member.Tag(), role),
	}
}

func MemberRoleRemoveLogger(member discord.User, role discord.RoleID) Logger {
	return Logger{
		Type:   LogTypeMemberRoleRemove,
		Author: member,
		Footer: &discord.EmbedFooter{Text: "Role ID: " + role.String()},
		Message: fmt.Sprintf("A role has been removed from a member: %s\n\nRole: <@&%s>",
			member.Tag(), role),
	}
}
