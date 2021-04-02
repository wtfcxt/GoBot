package bot

import (
	"GoBot/commands"
	"GoBot/database"
	"GoBot/util/embed"
)

func Warnings(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	guild, err := s.Guild(m.GuildID)

	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	database.GetWarnings(database.GetClient(), guild, m.Author)

}