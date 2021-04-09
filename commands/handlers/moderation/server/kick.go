package server

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Kick(ctx *commands.Context) {

	s := ctx.Session
	event := ctx.Event

	message := strings.Split(event.Message.Content, " ")
	guild, err := s.Guild(event.GuildID)
	if err != nil {
		embed.ThrowError(err.Error(), s, event)
	}

	prefix := database.GetGuildValue(guild, "prefix")

	s.ChannelTyping(event.ChannelID)

	if util.HasPermission(s, event, discordgo.PermissionKickMembers) {

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
						Value:  "`" + prefix + "kick <User> [Reason]` - Kicks a user",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			} else {
				err = s.GuildMemberDelete(event.GuildID, user.ID)

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

				_, err := s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Kicked "+user.Username, embed.Green, field))
				if err != nil {
					embed.ThrowError(err.Error(), s, event)
				}
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
						Value:  "`" + prefix + "kick <User> [Reason]` - Kicks a user",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			} else {
				err = s.GuildMemberDelete(event.GuildID, user.ID)

				field := []*discordgo.MessageEmbedField{
					{
						Name:   "User",
						Value:  "<@" + user.ID + ">",
						Inline: true,
					},
					{
						Name:   "Reason",
						Value:  "`Not supplied`",
						Inline: true,
					},
				}

				_, err := s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Kicked "+user.Username, embed.Green, field))
				if err != nil {
					embed.ThrowError(err.Error(), s, event)
				}
			}
		} else {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Invalid syntax.",
					Value:  "You didn't supply a user for me to ban.",
					Inline: false,
				},
				{
					Name:   "Correct syntax",
					Value:  "`" + prefix + "ban <User> [Reason]` - Bans a user",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
		}
	} else {
		embed.NoPermsEmbed(s, event, "Kick Members")
	}

}