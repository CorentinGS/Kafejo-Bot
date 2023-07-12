package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/corentings/kafejo-bot/app/handler"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
)

const (
	day    = time.Minute * 60 * 24
	minute = time.Minute
	hour   = time.Hour
	year   = 365 * day
)

func GetSession() *state.State {
	return handler.GetHandler().GetState()
}

func FormatTimeSince(old time.Time) string {
	duration := time.Since(old)

	if duration < time.Second {
		return "just now"
	}

	if duration < minute {
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	}

	if duration < hour {
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	}

	if duration < day {
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	}

	var b strings.Builder

	if duration >= year {
		years := duration / year
		_, err := fmt.Fprintf(&b, "%dy ", years)
		if err != nil {
			return ""
		}
		duration -= years * year
	}

	months := duration / (30 * day)
	duration -= months * (30 * day)
	_, err := fmt.Fprintf(&b, "%dm ", months)
	if err != nil {
		return ""
	}

	days := duration / day
	duration -= days * day
	_, err = fmt.Fprintf(&b, "%dd ", days)
	if err != nil {
		return ""
	}

	hours := duration / time.Hour
	duration -= hours * time.Hour
	_, err = fmt.Fprintf(&b, "%dh ", hours)
	if err != nil {
		return ""
	}

	minutes := duration / time.Minute
	duration -= minutes * time.Minute
	_, err = fmt.Fprintf(&b, "%dm ", minutes)
	if err != nil {
		return ""
	}

	seconds := duration / time.Second
	_, err = fmt.Fprintf(&b, "%ds", seconds)
	if err != nil {
		return ""
	}

	return b.String()
}

func IsDefaultAvatar(avatar string) bool {
	// Default avatar is "https://cdn.discordapp.com/embed/avatars/" + [0-5] + ".png"
	return strings.HasPrefix(avatar, "https://cdn.discordapp.com/embed/avatars/") && strings.HasSuffix(avatar, ".png") && avatar[len(avatar)-5] >= '0' && avatar[len(avatar)-5] <= '5'
}

func GetRoleDifference(newRoles, oldRoles []discord.RoleID) discord.RoleID {
	for _, newRole := range newRoles {
		found := false
		for _, oldRole := range oldRoles {
			if newRole == oldRole {
				found = true
				break
			}
		}
		if !found {
			return newRole
		}
	}
	return 0
}
