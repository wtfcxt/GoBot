package util

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func HasPermission(s *discordgo.Session, m *discordgo.MessageCreate, permission int64) bool {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		fmt.Println(err.Error())
	}

	if p&permission == permission {
		return true
	} else {
		return false
	}
}