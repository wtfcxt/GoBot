package bot

import (
	"GoBot/util"
	"github.com/bwmarrin/discordgo"
)

func ReadyEvent(s *discordgo.Session, e *discordgo.Ready) {

	version := "10.0"
	branch := "Dev"

	err := s.UpdateGameStatus(1, "GoBot X | "+version+"/"+branch)
	if err != nil {
		util.LogCrash(err)
	}

}
