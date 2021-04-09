package bot

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Warn(ctx *commands.Context) {
	s := ctx.Session
	event := ctx.Event

	message := strings.Split(event.Message.Content, " ")
	guild, err := s.Guild(event.GuildID)
	if err != nil {
		embed.ThrowError(err.Error(), s, event)
	}

	prefix := database.GetGuildValue(guild, "prefix")

	s.ChannelTyping(event.ChannelID)

	if util.HasPermission(s, event, discordgo.PermissionManageMessages) {


		if len(message) >= 3 {
			user := event.Mentions[0]
			reason := message[1]



			if user == nil {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "No user supplied.",
						Value:  "You didn't supply a user.",
						Inline: false,
					},
					{
						Name:   "Correct syntax",
						Value:  "`" + prefix + "warn <user> [reason]` - Warns a user",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			} else {
				database.AddWarning(user, guild, reason)

				field := []*discordgo.MessageEmbedField{
					{
						Name:   "User",
						Value:  "<@" + user.ID + ">",
						Inline: true,
					},
					{
						Name:   "Reason",
						Value:  "`" + reason + "`",
						Inline: true,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Warned "+user.Username, embed.Green, field))
			}
		} else if len(message) == 2 {
			user := event.Mentions[0]
			if user == nil {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "No user supplied.",
						Value:  "You didn't supply a user.",
						Inline: false,
					},
					{
						Name:   "Correct syntax",
						Value:  "`" + prefix + "warn <user> [reason]` - Warns a user",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			} else {
				database.AddWarning(user, guild, "None")

				field := []*discordgo.MessageEmbedField{
					{
						Name:   "User",
						Value:  "<@" + user.ID + ">",
						Inline: true,
					},
					{
						Name:   "Reason",
						Value:  "`None`",
						Inline: true,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbed("Warned: "+user.Username, "This action has been performed successfully.", "https://files.cxt.wtf/GoBot/hammer_green.png", embed.Green, field))
			}
		} else {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Invalid syntax.",
					Value:  "You didn't supply a user for me to warn.",
					Inline: false,
				},
				{
					Name:   "Correct syntax",
					Value:  "`" + prefix + "warn <User> [Reason]` - Warns a user",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			return
		}
	} else {
		embed.NoPermsEmbed(s, event, "Manage Messages")
	}
}
