package admin

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

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	prefix := database.GetGuildValue(guild, "prefix")

	s.ChannelTyping(m.ChannelID)

	if util.HasPermission(s, m, discordgo.PermissionAdministrator) {
		if len(message) == 1 {

			muterole := database.GetGuildValue(guild, "mute_role_id")
			warnchannel := database.GetGuildValue(guild, "warn_channel_id")

			if muterole == "none" {
				muterole = "**Not set**"
			} else {
				muterole = "<@&" + muterole + ">"
			}
			if warnchannel == "none" {
				warnchannel = "**Not set**"
			} else {
				warnchannel = "<#" + warnchannel + ">"
			}

			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Settings",
					Value:  "`" + prefix + "settings prefix <Prefix>` - Change the bot's prefix\n`" + prefix + "settings muterole <Role-Mention>` - Change the mute role\n`" + prefix + "settings warnchannel <Channel>` - If set, the bot will announce Warns in that channel.",
					Inline: false,
				},
				{
					Name:   "Current values",
					Value:  "Prefix: `" + database.GetGuildValue(guild, "prefix") + "`\nMute-Role: " + muterole + "\nWarn-Channel: " + warnchannel,
					Inline: false,
				},
			}
			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("Bot Settings", embed.Green, field))
		} else if len(message) == 2 {

			setting := message[1]

			switch setting {
			case "logger":
				enabled := database.GetGuildLoggerValue(guild, "channel")
				var channel string
				temp := database.GetGuildLoggerValue(guild, "channel")
				if temp == "none" {
					channel = "`none`"
				} else {
					channel = temp
				}
				messages := database.GetGuildLoggerValue(guild, "messages")
				joins := database.GetGuildLoggerValue(guild, "joins")
				leaves := database.GetGuildLoggerValue(guild, "leaves")
				role := database.GetGuildLoggerValue(guild, "role")
				punishments := database.GetGuildLoggerValue(guild, "punishments")

				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Logger",
						Value:  prefix + "settings logger enabled toggle\n" + prefix + "settings logger channel <Log-Channel>\n" + prefix + "settings logger messages toggle\n" + prefix + "settings logger joins toggle\n" + prefix + "settings logger leaves toggle\n" + prefix + "settings logger role toggle",
						Inline: false,
					},
					{
						Name:   "Current Values",
						Value:  "\n***General***\nEnabled: `" + enabled+
							"`\nChannel: " + channel +
							"\n\n***Actions to log***\nMessages: `" + messages +
							"`\nJoins: `" + joins + "`\nLeaves: `" + leaves + "`\nRole-Changes: `" + role + "`\nPunishments: `" + punishments + "`",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("Bot Settings", embed.Green, field))
			default:

			}
		} else if len(message) == 3 {
			setting := message[1]
			value := message[2]

			guild, err := s.Guild(m.GuildID)
			if err != nil {
				panic(err)
			}

			switch setting {
			case "prefix":
				if len(value) > 3 {
					field := []*discordgo.MessageEmbedField{
						{
							Name:   "Prefix must be shorter than 3 characters.",
							Value:  "If you want to mention GoBot, use \"mention\" as your prefix.",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
					return
				}
				if value == "mention" {
					value = "<@" + s.State.SessionID + ">"
				}
				SettingsChangedEmbed(s, m, "Prefix", value)
				database.ChangeGuildValue(guild, "prefix", value)
			case "muterole":
				SettingsChangedEmbed(s, m, "Mute Role", value)
				database.ChangeGuildValue(guild, "mute_role_id", strings.Replace(strings.Replace(value, ">", "", 1), "<@&", "", 1))
			case "warnchannel":
				SettingsChangedEmbed(s, m, "Warn Channel", value)
				database.ChangeGuildValue(guild, "warn_channel_id", strings.Replace(strings.Replace(value, ">", "", 1), "<#", "", 1))
			default:
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Unknown setting.",
						Value:  "The setting you supplied does not exist.",
						Inline: false,
					},
					{
						Name:   "Possible options",
						Value:  "`prefix, muterole, warnchannel`, more Info using !settings",
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
	if setting == "Prefix" && len(value) <= 3 {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "Success.",
				Value:  "Changed the following setting: " + setting + ".",
				Inline: false,
			},
			{
				Name:   "Value",
				Value:  "New Value: `" + value + "`",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("Setting changed.", embed.Green, field))
	} else {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "Success.",
				Value:  "Changed the following setting: " + setting + ".",
				Inline: false,
			},
			{
				Name:   "Value",
				Value:  "New Value: " + value + "",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("Setting changed.", embed.Green, field))
	}
}
