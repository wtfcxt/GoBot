package server

import (
	"GoBot/database"
	"github.com/bwmarrin/discordgo"
)

func GuildAdd(session *discordgo.Session, event *discordgo.GuildCreate) {
	if !database.GuildExists(database.GetClient(), event.Guild) {
		database.AddGuild(database.GetClient(), event.Guild)
		database.AddAllMembers(database.GetClient(), event.Guild)
	}
}
