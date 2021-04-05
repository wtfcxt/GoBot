package misc

import (
	"GoBot/commands"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
)

func Info(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	s.ChannelTyping(m.ChannelID)

	field := []*discordgo.MessageEmbedField{
		{
			Name: "Developer",
			Value: "`cxt#1234` made this awesome bot!",
			Inline: false,
		},
		{
			Name: "Version",
			Value: "Running `GoBot X` | Branch `Dev`",
			Inline: false,
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("GoBot Info", embed.Green, field))

}