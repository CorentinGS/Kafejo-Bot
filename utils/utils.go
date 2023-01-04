package utils

import (
	"fmt"
	"github.com/corentings/kafejo-bot/data/cmdHandler"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"strings"
	"time"
)

const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

func GetSession() *state.State {
	return cmdHandler.GetHandler().GetState()
}

func TimeSince(old time.Time) time.Duration {
	now := time.Now()
	diff := now.Sub(old)
	return diff
}

func FormatTimeSince(old time.Time) string {
	d := TimeSince(old)
	if d < day {
		return fmt.Sprintf("%d hours ago", int(d.Hours()))
	}

	var b strings.Builder
	if d >= year {
		years := d / year
		_, err := fmt.Fprintf(&b, "%dy ", years)
		if err != nil {
			return ""
		}
		d -= years * year
	}

	months := d / (30 * day)
	d -= months * (30 * day)
	_, err := fmt.Fprintf(&b, "%dm ", months)
	if err != nil {
		return ""
	}

	days := d / day
	d -= days * day
	_, err = fmt.Fprintf(&b, "%dd", days)
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
