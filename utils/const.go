package utils

import "github.com/diamondburned/arikawa/v3/discord"

const (
	GITHUB  = "https://github.com/CorentinGS/Kafejo-Bot"
	VERSION = "v0.1.0"
	GREEN   = 0x00ff00
	RED     = 0xff0000
	BLUE    = 0x0000ff
	ORANGE  = 0xffa500
	PURPLE  = 0x800080

	SQLMaxOpenConns    = 50
	SQLMaxIdleConns    = 10
	CacheExpireMinutes = 2
)

const (
	ConfigRoleAdmin         = "1059086726565462136"
	ConfigRoleMod           = "1059086750921801749"
	ConfigOwnerID           = "282233191916634113"
	ConfigMainRole          = "1059086753425805354"
	ConfigGuildID           = "560798438099255296"
	ConfigGateKeepChannelID = "1059086776368631878"
	ConfigLogsChannelID     = "960577800589434950"
	ConfigWelcomeChannelID  = "1059086801203105863"
	ConfigAdminChannelID    = "1059086820442382348"
	ConfigWelcomeMessageID  = "1060225318453002351"
)

func GetGuildID() discord.GuildID {
	guildID, _ := discord.ParseSnowflake(ConfigGuildID)
	return discord.GuildID(guildID)
}
