package bot

import (
	"GoBot/commands"
	"github.com/bwmarrin/discordgo"
	"sync"
)

func Clear(ctx *commands.Context) {
		
}

func deleteMessages(msgs []string, wg *sync.WaitGroup, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessagesBulkDelete(m.ChannelID, msgs)
	defer wg.Done()
}
