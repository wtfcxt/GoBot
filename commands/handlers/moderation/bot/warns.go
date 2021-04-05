package bot

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util/embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func Warnings(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	if len(strings.Split(m.Content, " ")) == 2 {

		target := m.Mentions[0]

		s.ChannelTyping(m.ChannelID)

		guild, err := s.Guild(m.GuildID)

		if err != nil {
			embed.ThrowError(err.Error(), s, m)
		}

		output := database.GetWarnings(target, guild)
		sorted := make(map[int]interface{})

		field := []*discordgo.MessageEmbedField {
			{
				Name:   strconv.Itoa(len(sorted)) + " warning(s).",
				Value:  "<@" + target.ID + "> currently has " + strconv.Itoa(len(sorted)) + "warn(s).",
				Inline: false,
			},
		}

		for i := range output {
			if output[i] == "" {
				continue
			} else {
				sorted[i] = output[i]

				appendField := discordgo.MessageEmbedField {
					Name:   "Warning #" + strconv.Itoa(i),
					Value:  fmt.Sprint(output[i]),
					Inline: true,
				}
				field = append(field, &appendField)
			}
		}

		if len(sorted) == 0 {
			field := []*discordgo.MessageEmbedField {
				{
					Name:   "0 warning(s).",
					Value:  "Woah! This user is clean. _Are you?_",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed(target.Username + "'s warnings", "This command will show their warnings. _if they have any..._", "", embed.Orange, field))
			return
		} else {
			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed(target.Username + "'s warnings", "This command will show their warnings.", "", embed.Orange, field))
		}
	} else {
		s.ChannelTyping(m.ChannelID)

		guild, err := s.Guild(m.GuildID)

		if err != nil {
			embed.ThrowError(err.Error(), s, m)
		}

		output := database.GetWarnings(m.Author, guild)
		sorted := make(map[int]interface{})

		var field []*discordgo.MessageEmbedField

		for i := range output {
			if output[i] == "" {
				continue
			} else {
				sorted[i] = output[i]

				appendField := discordgo.MessageEmbedField {
					Name:   "Warning #" + strconv.Itoa(i),
					Value:  fmt.Sprint(output[i]),
					Inline: true,
				}
				field = append(field, &appendField)
			}
		}

		if len(sorted) == 0 {
			field := []*discordgo.MessageEmbedField {
				{
					Name:   "0 warning(s).",
					Value:  "Woah! You are clean. Keep up the good work.",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed(m.Author.Username + "'s warnings", "This command will show your warnings. _if you have any..._", "https://files.cxt.wtf/GoBot/hammer_green.png", embed.Green, field))
			return
		} else {
			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed(m.Author.Username + "'s warnings", "This command will show your warnings.", "https://files.cxt.wtf/GoBot/hammer_red.png", embed.Red, field))
		}
	}

}