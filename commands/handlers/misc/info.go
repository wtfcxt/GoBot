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
			Name:   "Version",
			Value:  "Running `GoBot X` | Branch `Dev`",
			Inline: false,
		},
		{
			Name:   "Contributors",
			Value:  "<@704419523922493542> - Main developer\n<@813305613671989259> - Some ideas",
			Inline: false,
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("GoBot Info", embed.Green, field))

}
