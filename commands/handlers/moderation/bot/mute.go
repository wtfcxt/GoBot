package bot

import (
	"GoBot/commands"
	new2 "GoBot/database/new"
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

	s.ChannelTyping(event.ChannelID)

	if util.HasPermission(s, event, discordgo.PermissionManageMessages) {
		if new2.GetGuildValue(guild, "mute_role_id") != "none" {
			if len(message) >= 2 {
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
							Value:  "`!mute <User> [Reason]` - Mutes a user",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
					return
				} else {
					if !new2.GetUserValueBool(user, guild, "muted") {
						err := s.GuildMemberRoleAdd(event.GuildID, user.ID, new2.GetGuildValue(guild, "mute_role_id"))
						if err != nil {
							field := []*discordgo.MessageEmbedField{
								{
									Name:   "Invalid Mute role.",
									Value:  "You either didn't set the Mute role or the one you set is invalid.",
									Inline: false,
								},
								{
									Name:   "How to fix?",
									Value:  "`!settings muterole <Ping Mute Role>` - Sets the mute role",
									Inline: false,
								},
							}

							s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
						}
						new2.ChangeUserValueBool(user, guild, "muted", true)

						field := []*discordgo.MessageEmbedField{
							{
								Name:   "User",
								Value:  "<@" + user.ID + ">",
								Inline: true,
							},
							{
								Name:   "Moderator",
								Value:  "<@" + event.Author.ID + ">",
								Inline: true,
							},
						}

						s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("Muted " + user.Username, embed.Green, field))
					} else {
						field := []*discordgo.MessageEmbedField{
							{
								Name:   "User is already muted.",
								Value:  "The user you supplied has already been muted.",
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
						Value:  "You didn't supply a user for me to mute.",
						Inline: false,
					},
					{
						Name:   "Correct syntax",
						Value:  "`!mute <User>` - Mutes a user",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			}
		} else {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Invalid Mute role.",
					Value:  "You either didn't set the Mute role or the one you set is invalid.",
					Inline: false,
				},
				{
					Name:   "How to fix?",
					Value:  "`!settings muterole <Role-Mention>` - Sets the mute role",
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
