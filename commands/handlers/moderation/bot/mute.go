package bot

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Mute(ctx *commands.Context) {

	s := ctx.Session
	event := ctx.Event

	message := strings.Split(event.Message.Content, " ")
	guild, _ := s.Guild(event.GuildID)

	if util.HasPermission(s, event, discordgo.PermissionManageMessages) {
		if len(message) == 2 {
			member := event.Mentions[0]
			if member == nil {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "No member supplied.",
						Value:  "You didn't supply a member.",
						Inline: false,
					},
					{
						Name:   "Correct syntax",
						Value:  "`!mute <user> [reason]` - Mutes a member",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			} else {
				if !database.GetMemberValueBoolean(database.GetClient(), member, "muted") {
					err := s.GuildMemberRoleAdd(event.GuildID, member.ID, database.GetSetting(database.GetClient(), guild, "muterole"))
					if err != nil {
						embed.ThrowError(err.Error(), s, event)
					}
					database.ChangeMemberOptionBool(database.GetClient(), member, "muted", true)

					field := []*discordgo.MessageEmbedField{
						{
							Name:   "Member muted.",
							Value:  "I muted <@" + member.ID + "> for you.",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Action successful.", embed.Green, field))
				} else {
					field := []*discordgo.MessageEmbedField{
						{
							Name:   "Member is already muted.",
							Value:  "The member you supplied has already been muted.",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				}
			}
		} else if len(message) == 1 {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Invalid syntax.",
					Value:  "You didn't supply the amount of messages to delete.",
					Inline: false,
				},
				{
					Name:   "Correct syntax",
					Value:  "`!mute <user>` - Mutes a member",
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
