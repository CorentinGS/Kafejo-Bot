package common

import (
	"strings"
	"time"

	"github.com/corentings/kafejo-bot/utils"

	"github.com/diamondburned/arikawa/v3/discord"
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
		flag++
	}

	// if the member hasn't a custom avatar
	if utils.IsDefaultAvatar(member.User.AvatarURL()) {
		flag++
	}

	// if the member has MFA enabled
	if member.User.MFA {
		flag--
	}

	if member.User.Nitro > discord.NoUserNitro {
		flag--
	}

	// if the member has no email verified
	if member.User.EmailVerified == false {
		flag = MemberDangerLevelHigh
	}

	// if the member is flagged as a spammer
	if member.User.PublicFlags&discord.LikelySpammer != 0 {
		flag = MemberDangerLevelHigh
	}

	cleanUsername := strings.ToLower(member.User.Username)
	// if the username has an invitation link
	if strings.Contains(cleanUsername, "discord.gg/") || strings.Contains(cleanUsername, "discord.com/invite/") {
		flag = MemberDangerLevelHigh
	}

	// if the username has a link
	//goland:noinspection HttpUrlsUsage
	if strings.Contains(cleanUsername, "http://") || strings.Contains(cleanUsername, "https://") {
		flag = MemberDangerLevelHigh
	}

	return flag
}
