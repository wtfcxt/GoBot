package bot

import "github.com/bwmarrin/discordgo"

func ReadyEvent(s *discordgo.Session, e *discordgo.Ready) {

	version := "10.0"
	branch := "Dev"

	s.UpdateGameStatus(1, "GoBot X | " + version + "/" + branch)
}