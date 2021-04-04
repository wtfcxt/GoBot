package misc

import (
	"GoBot/commands"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
)

func Help(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	if util.HasPermission(s, m, discordgo.PermissionAdministrator) {
		field := []*discordgo.MessageEmbedField{
			{
				Name: "Administration",
				Value: "`!settings` - Change Server-specific Bot Settings\n`!module` - Disable a module (e.g. Moderation Module)",
				Inline: false,
			},
			{
				Name: "Moderation",
				Value: "`!clear` - Clears the chat\n`!mute <Member> [Reason]` - Mutes a member\n`!unmute` - Unmutes a member\n`!warn` - Warns a member\n`!ban` - Bans a member\n`!kick` - Kicks a member",
				Inline: false,
			},
			{
				Name: "Fun",
				Value: "`!meme` - Generate a random meme\n`!zap` - Zap another member",
				Inline: false,
			},
			{
				Name: "Music",
				Value: "_Soon..._",
				Inline: false,
			},
			{
				Name: "Misc",
				Value: "`!help` - Shows this help page\n`!info` - Shows the GoBot Info Page",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("GoBot Help", embed.Green, field))
	} else if util.HasPermission(s, m, discordgo.PermissionManageMessages) {
		field := []*discordgo.MessageEmbedField{
			{
				Name: "Moderation",
				Value: "`!clear` - Clears the chat\n`!mute <Member> [Reason]` - Mutes a member\n`!unmute` - Unmutes a member\n`!warn` - Warns a member\n`!ban` - Bans a member\n`!kick` - Kicks a member",
				Inline: false,
			},
			{
				Name: "Fun",
				Value: "`!meme` - Generate a random meme\n`!zap` - Zap another member",
				Inline: false,
			},
			{
				Name: "Music",
				Value: "_Soon..._",
				Inline: false,
			},
			{
				Name: "Misc",
				Value: "`!help` - Shows this help page\n`!info` - Shows the GoBot Info Page",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("GoBot Help", embed.Green, field))
	} else {
		field := []*discordgo.MessageEmbedField{
			{
				Name: "Fun",
				Value: "`!meme` - Generate a random meme\n`!zap` - Zap another member",
				Inline: false,
			},
			{
				Name: "Music",
				Value: "_Soon..._",
				Inline: false,
			},
			{
				Name: "Misc",
				Value: "`!help` - Shows this help page\n`!info` - Shows the GoBot Info Page",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("GoBot Help", embed.Green, field))
	}

}
