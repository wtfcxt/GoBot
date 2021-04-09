package fun

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func Hack(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	message := strings.Split(m.Message.Content, " ")

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	prefix := database.GetGuildValue(guild, "prefix")

	switch len(message) {
	case 2:
		user := m.Mentions[0]
		if user == nil {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Invalid syntax.",
					Value:  "You can't hack the air...",
					Inline: false,
				},
				{
					Name:   "Correct syntax",
					Value:  "`" + prefix + "hack <User>` - Hacks a user",
					Inline: false,
				},
			}

			_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
		} else {
			user := m.Mentions[0]
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Grabbing IP address...",
					Value:  "[||#||##########]",
					Inline: false,
				},
			}
			msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("Hacking "+user.Username+"...", "This process will only take a few seconds...", "https://files.cxt.wtf/GoBot/hack.png", embed.Green, field))
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
			time.Sleep(2 * time.Second)
			field2 := []*discordgo.MessageEmbedField{
				{
					Name:   "Hacking computer...",
					Value:  "[||####||#######]",
					Inline: false,
				},
			}
			msg, err = s.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, embed.CreateEmbed("Hacking "+user.Username+"...", "This process will only take a few seconds...", "https://files.cxt.wtf/GoBot/hack.png", embed.Green, field2))
			time.Sleep(2 * time.Second)
			field3 := []*discordgo.MessageEmbedField{
				{
					Name:   "Hacking phone...",
					Value:  "[||#######||####]",
					Inline: false,
				},
			}
			msg, err = s.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, embed.CreateEmbed("Hacking "+user.Username+"...", "This process will only take a few seconds...", "https://files.cxt.wtf/GoBot/hack.png", embed.Green, field3))
			time.Sleep(2 * time.Second)
			field4 := []*discordgo.MessageEmbedField{
				{
					Name:   "Uploading to website...",
					Value:  "[||##########||#]",
					Inline: false,
				},
			}
			msg, err = s.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, embed.CreateEmbed("Hacking "+user.Username+"...", "This process will only take a few seconds...", "https://files.cxt.wtf/GoBot/hack.png", embed.Green, field4))
			time.Sleep(2 * time.Second)
			msg, err = s.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, embed.CreateEmbed("Hacking "+user.Username+"...", "**This hack was successful.**\n\nIP-Address: ||Never Gonna Give You Up||\nPhone Number: ||Never Gonna Let You Down||\nAddress: \n||Never Gonna Make You Cry||\n||Never Gonna|| ||Say Goodbye||", "https://files.cxt.wtf/GoBot/hack.png", embed.Green, nil))
		}
	default:
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "Invalid syntax.",
				Value:  "You can't hack yourself... right?",
				Inline: false,
			},
			{
				Name:   "Correct syntax",
				Value:  "`" + prefix + "hack <User>` - Hacks a user",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
	}
}
