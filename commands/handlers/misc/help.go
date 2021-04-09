package misc

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
)

func Help(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	prefix := database.GetGuildValue(guild, "prefix")

	s.ChannelTyping(m.ChannelID)

	if util.HasPermission(s, m, discordgo.PermissionAdministrator) {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "Admin",
				Value:  "`" + prefix + "settings` `" + prefix + "module`",
				Inline: false,
			},
			{
				Name:   "Mod",
				Value:  "`" + prefix + "clear` `" + prefix + "mute` `" + prefix + "unmute` `" + prefix + "warn` `" + prefix + "ban` `" + prefix + "kick`",
				Inline: false,
			},
			{
				Name:   "Fun",
				Value:  "`" + prefix + "meme` `" + prefix + "zap` `" + prefix + "bite` `" + prefix + "hack` `" + prefix + "hug` `" + prefix + "ping` `" + prefix + "say` `" + prefix + "arguments`",
				Inline: false,
			},
			{
				Name:   "Misc",
				Value:  "`" + prefix + "help` `" + prefix + "info`",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("GoBot Help", "This is my help page. You can try every command on this list!", "https://files.cxt.wtf/GoBot/exclamationmark_green.png", embed.Green, field))
	} else if util.HasPermission(s, m, discordgo.PermissionManageMessages) {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "Moderation",
				Value:  "`" + prefix + "clear` `" + prefix + "mute <Member> [Reason]` `" + prefix + "unmute` `" + prefix + "warn` `" + prefix + "ban` `" + prefix + "kick`",
				Inline: false,
			},
			{
				Name:   "Fun",
				Value:  "`" + prefix + "meme` `" + prefix + "zap` `" + prefix + "bite` `" + prefix + "hack` `" + prefix + "hug` `" + prefix + "ping` `" + prefix + "say`",
				Inline: false,
			},
			{
				Name:   "Misc",
				Value:  "`" + prefix + "help` `" + prefix + "info`",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("GoBot Help", "This is my help page. You can try every command on this list!", "https://files.cxt.wtf/GoBot/exclamationmark_green.png", embed.Green, field))
	} else {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "Fun",
				Value:  "`" + prefix + "meme` `" + prefix + "zap` `" + prefix + "bite` `" + prefix + "hack` `" + prefix + "hug` `" + prefix + "ping` `" + prefix + "say`",
				Inline: false,
			},
			{
				Name:   "Misc",
				Value:  "`" + prefix + "help` `" + prefix + "info`",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("GoBot Help", "This is my help page. You can try every command on this list!", "https://files.cxt.wtf/GoBot/exclamationmark_green.png", embed.Green, field))
	}

}
