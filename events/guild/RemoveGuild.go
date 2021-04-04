package guild

import (
	"GoBot/database"
	"github.com/bwmarrin/discordgo"
)

func Remove(session *discordgo.Session, event *discordgo.GuildDelete) {
	database.RemoveGuild(database.GetClient(), event.Guild)
}
