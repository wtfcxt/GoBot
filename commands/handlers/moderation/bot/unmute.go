package bot

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Unmute(ctx *commands.Context) {

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
		if database.GetGuildValue(guild, "mute_role_id") != "none" {
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
							Value:  "`!unmute <User>` - Unmutes a user",
							Inline: false,
						},
					}

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
					return
				} else {
					if database.GetUserValueBool(user, guild, "muted") {
						err := s.GuildMemberRoleRemove(event.GuildID, user.ID, database.GetGuildValue(guild, "mute_role_id"))
						if err != nil {
							field := []*discordgo.MessageEmbedField{
								{
									Name:   "User not muted.",
									Value:  "Looks like this user hasn't been muted.",
									Inline: false,
								},
							}

							s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
						}
						database.ChangeUserValueBool(user, guild, "muted", false)

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

						s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbed("Unmuted: "+user.Username, "This action has been performed successfully.", "https://files.cxt.wtf/GoBot/hammer_green.png", embed.Green, field))
					} else {
						field := []*discordgo.MessageEmbedField{
							{
								Name:   "User is not muted.",
								Value:  "The user you supplied isn't muted.",
								Inline: false,
							},
							{
								Name:   "...but the user still has the Muted Role?",
								Value:  "If that's the case, someone maybe gave it to them without using the mute command.",
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
						Value:  "You didn't supply a user for me to unmute.",
						Inline: false,
					},
					{
						Name:   "Correct syntax",
						Value:  "`" + prefix + "mute <user>` - Unmutes a user",
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
					Value:  "`" + prefix + "settings muterole <Ping Mute Role>` - Sets the mute role",
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
