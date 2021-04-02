package server

import (
	"GoBot/database"
	"github.com/bwmarrin/discordgo"
)

func GuildRemove(session *discordgo.Session, event *discordgo.GuildDelete) {
	database.RemoveGuild(database.GetClient(), event.Guild)
}
