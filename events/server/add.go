package server

import (
	new2 "GoBot/database/new"
	"github.com/bwmarrin/discordgo"
)

func GuildAdd(session *discordgo.Session, event *discordgo.GuildCreate) {
	/*if !database.GuildExists(database.GetClient(), event.Guild) {
		database.AddGuild(database.GetClient(), event.Guild)
		database.AddAllMembers(database.GetClient(), event.Guild)
	}*/
	new2.RegisterGuild(event.Guild)
	new2.RegisterUserBulk(event.Guild)
}
