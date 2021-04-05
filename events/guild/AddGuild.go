package guild

import (
	"GoBot/database"
	"GoBot/util/logger"
	"github.com/bwmarrin/discordgo"
	"sync"
)

func Add(session *discordgo.Session, event *discordgo.GuildCreate) {
	var wg sync.WaitGroup
	if !database.GuildExists(event.Guild) {
		wg.Add(1)
		logger.LogModule(logger.TypeDebug, "GoBot/Debug", "Guild does not exist. Registering... (Guild: " + event.Guild.ID + ")")
		go database.RegisterGuild(event.Guild, session, &wg)
	} else {
		logger.LogModule(logger.TypeDebug, "GoBot/Debug", "Guild already exists, not registering. (Guild: " + event.Guild.ID + ")")
	}
}
