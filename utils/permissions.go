package utils

import "github.com/diamondburned/arikawa/v3/discord"

// checks if the user is an admin
func isAdmin(member discord.Member) bool {
	for _, role := range member.RoleIDs {
		if role.String() == ConfigRoleAdmin {
			return true
		}
	}
	return false
}

// checks  if the user is a moderator
func isMod(member discord.Member) bool {
	for _, role := range member.RoleIDs {
		if role.String() == ConfigRoleMod {
			return true
		}
	}
	return false
}

func isOwner(member discord.Member) bool {
	return member.User.ID.String() == ConfigOwnerID
}

// IsAdminOrMod checks if the user is an admin or a moderator
func IsAdminOrMod(member discord.Member) bool {
	return isAdmin(member) || isMod(member)
}

// HasAdminPermission checks if the user has admin permission
func HasAdminPermission(member discord.Member) bool {
	return isAdmin(member) || isOwner(member)
}

// HasModPermission checks if the user has mod permission
func HasModPermission(member discord.Member) bool {
	return isMod(member) || HasAdminPermission(member)
}

// HasOwnerPermission checks if the user has owner permission
func HasOwnerPermission(member discord.Member) bool {
	return isOwner(member)
}
