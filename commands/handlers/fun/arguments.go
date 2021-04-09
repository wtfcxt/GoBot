package fun

import (
	"GoBot/commands"
	"time"
)

func ArgumentCommand(ctx *commands.Context) {

	s := ctx.Session
	event := ctx.Event

	s.ChannelTyping(event.ChannelID)

	time.Sleep(5*time.Second)
	s.ChannelMessageSend(event.ChannelID, "I don't have any arguments...")
}