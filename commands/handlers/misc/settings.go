package misc

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Settings(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	message := strings.Split(m.Message.Content, " ")

	if util.HasPermission(s, m, discordgo.PermissionAdministrator) {
		if len(message) <= 2 {
			field := []*discordgo.MessageEmbedField{
				{
					Name: "Settings",
					Value: "`!settings prefix <Prefix>` - Change the bot's prefix\n`!settings muterole <Role-Mention>` - Change the mute role\n`!settings warnch <Channel>` - If set, the bot will announce Warns in that channel.",
					Inline: false,
				},
				{
					Name: "Current values",
					Value: "Prefix: `" + database.GetSetting(database.GetClient(), "prefix") + "`\nMute-Role: `" + database.GetSetting(database.GetClient(), "muterole") + "`\nWarn-Channel: `" + database.GetSetting(database.GetClient(), "warnch") + "`",
					Inline: false,
				},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("Bot Settings", embed.Green, field))
		} else if len(message) == 3 {
			setting := message[1]
			value := message[2]

			guild, err := s.Guild(m.GuildID)
			if err != nil {
				panic(err)
			}

			switch setting {
			case "prefix":
				SettingsChangedEmbed(s, m, setting, value)
				database.ChangeGuildSetting(database.GetClient(), guild, setting, value)
			case "muterole":
				SettingsChangedEmbed(s, m, setting, strings.Replace(strings.Replace(value, ">", "", 1), "<@&", "", 1))
				database.ChangeGuildSetting(database.GetClient(), guild, setting, strings.Replace(strings.Replace(value, ">", "", 1), "<@&", "", 1))
			case "warnch":
				SettingsChangedEmbed(s, m, setting, strings.Replace(strings.Replace(value, ">", "", 1), "<#", "", 1))
				database.ChangeGuildSetting(database.GetClient(), guild, setting, strings.Replace(strings.Replace(value, ">", "", 1), "<#", "", 1))
			default:
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Unknown setting.",
						Value:  "The setting you supplied does not exist.",
						Inline: false,
					},
					{
						Name: 	"Possible options",
						Value:	"`prefix, muterole, warnch`, more Info using !settings",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			}
		}
	} else {
		embed.NoPermsEmbed(s, m, "Administrator")
	}

}

func SettingsChangedEmbed(s *discordgo.Session, m *discordgo.MessageCreate, setting string, value string) {
	field := []*discordgo.MessageEmbedField{
		{
			Name:   "Success.",
			Value:  "I changed setting `" + setting + "`.",
			Inline: false,
		},
		{
			Name: 	"Value",
			Value:	"New Value: `" + value + "`",
			Inline: false,
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("Setting changed.", embed.Green, field))
}