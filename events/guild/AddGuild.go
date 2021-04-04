package guild

import (
	new2 "GoBot/database/new"
	"github.com/bwmarrin/discordgo"
)

func Add(session *discordgo.Session, event *discordgo.GuildCreate) {
	/*if !database.GuildExists(database.GetClient(), event.Guild) {
		database.AddGuild(database.GetClient(), event.Guild)
		database.AddAllMembers(database.GetClient(), event.Guild)
	}*/
	if !new2.GuildExists(event.Guild) {
		new2.RegisterGuild(event.Guild)
	} else {
		return
	}
}
