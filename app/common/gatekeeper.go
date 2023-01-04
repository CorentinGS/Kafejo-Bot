package common

import (
	"github.com/corentings/kafejo-bot/utils"
	"github.com/diamondburned/arikawa/v3/discord"
	"strings"
	"time"
)

type MemberDangerLevel int64

const (
	MemberDangerLevelNone MemberDangerLevel = iota
	MemberDangerLevelLow
	MemberDangerLevelMedium
	MemberDangerLevelHigh
)

func (m MemberDangerLevel) String() string {
	switch m {
	case MemberDangerLevelNone:
		return "None"
	case MemberDangerLevelLow:
		return "Low"
	case MemberDangerLevelMedium:
		return "Medium"
	case MemberDangerLevelHigh:
		return "High"
	default:
		return "Unknown"
	}
}

func VerifyMember(member *discord.Member) MemberDangerLevel {
	flag := MemberDangerLevelNone
	now := time.Now()

	// if member account has been created less than 1 week ago or 1 month ago
	if now.Sub(member.User.CreatedAt()) < time.Hour*24*7 {
		flag += 2
	} else if now.Sub(member.User.CreatedAt()) < time.Hour*24*30 {
		flag += 1
	}

	// if the member hasn't a custom avatar
	if utils.IsDefaultAvatar(member.User.AvatarURL()) {
		flag += 1
	}

	cleanUsername := strings.ToLower(member.User.Username)
	// if the username has an invitation link
	if strings.Contains(cleanUsername, "discord.gg/") || strings.Contains(cleanUsername, "discord.com/invite/") {
		flag = MemberDangerLevelHigh
	}

	// if the username has a link
	if strings.Contains(cleanUsername, "http://") || strings.Contains(cleanUsername, "https://") {
		flag = MemberDangerLevelHigh
	}

	return flag
}
