package guild

import (
	"GoBot/database"
	"github.com/bwmarrin/discordgo"
)

func UserJoin(session *discordgo.Session, event *discordgo.GuildMemberAdd) {

	guild, _ := session.Guild(event.GuildID)

	if !database.UserExists(event.User, guild) {
		database.CreateUser(event.User, guild)
	}
}