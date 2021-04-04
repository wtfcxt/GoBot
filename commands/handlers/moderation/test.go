package moderation

import (
	"GoBot/commands"
	new2 "GoBot/database/new"
)

func TestCommand(ctx *commands.Context) {
	m := ctx.Event
	s := ctx.Session

	guild, _ := s.Guild(m.GuildID)

	new2.IsUserInGuild(guild, m.Author)
}