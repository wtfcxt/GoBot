package bot

import (
	"GoBot/commands"
	new2 "GoBot/database/new"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Warn(ctx *commands.Context) {
	s := ctx.Session
	event := ctx.Event

	message := strings.Split(event.Message.Content, " ")
	guild, _ := s.Guild(event.GuildID)

	s.ChannelTyping(event.ChannelID)

	if util.HasPermission(s, event, discordgo.PermissionManageMessages) {
			if len(message) >= 3 {
				user := event.Mentions[0]
				reason := message[2]
				if user == nil {
					field := []*discordgo.MessageEmbedField{
						{
							Name:   "No user supplied.",
							Value:  "You didn't supply a user.",
							Inline: false,
						},
						{
							Name:   "Correct syntax",
							Value:  "`!warn <user> [reason]` - Warns a user",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
					return
				} else {
						new2.AddWarning(user, guild, reason)

						field := []*discordgo.MessageEmbedField{
							{
								Name:   "User",
								Value:  "<@" + user.ID + ">.",
								Inline: true,
							},
							{
								Name: 	"Reason",
								Value:  "`" + reason + "`",
								Inline: true,
							},
						}

						s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Warned " + user.Username, embed.Green, field))
				}
			} else if len(message) == 2 {
				member := event.Mentions[0]
				if member == nil {
					field := []*discordgo.MessageEmbedField{
						{
							Name:   "No member supplied.",
							Value:  "You didn't supply a user.",
							Inline: false,
						},
						{
							Name:   "Correct syntax",
							Value:  "`!warn <user> [reason]` - Warns a user",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
					return
				} else {
					new2.AddWarning(member, guild, "None")

					field := []*discordgo.MessageEmbedField{
						{
							Name:   "User",
							Value:  "<@" + member.ID + ">.",
							Inline: true,
						},
						{
							Name: 	"Reason",
							Value:  "`None`",
							Inline: true,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Warned " + member.Username, embed.Green, field))
				}
			} else {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Invalid syntax.",
						Value:  "You didn't supply a user for me to warn",
						Inline: false,
					},
					{
						Name:   "Correct syntax",
						Value:  "`!warn <User>` - Warns a user",
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