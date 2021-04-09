package guild

import (
	"GoBot/database"
	"github.com/bwmarrin/discordgo"
	"sync"
)

func Add(session *discordgo.Session, event *discordgo.GuildCreate) {

	var wg sync.WaitGroup

	if !database.GuildExists(event.Guild) {
		wg.Add(1)
		go database.RegisterGuild(event.Guild, session, &wg)
		return
	}

	wg.Add(1)
	for i := range event.Guild.Members {
		currentMember := event.Guild.Members[i].User
		if !database.UserExists(currentMember, event.Guild) {
			database.CreateUser(currentMember, event.Guild)
		}
	}
	defer wg.Done()
}
