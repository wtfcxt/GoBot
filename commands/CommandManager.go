package commands

import (
	"GoBot/database"
	"GoBot/util/cfg"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command struct {
	Command			string
	CommandHandler	func(context *Context)
}

type Context struct {
	Implementation 	*Command
	Event 			*discordgo.MessageCreate
	Session 		*discordgo.Session
	Label 			string
}

type CommandManager struct {
	Prefix			string
	Commands		[]Command
}

func NewCommandManager() CommandManager {
	return CommandManager{
		Prefix: cfg.Prefix,
		Commands: []Command{},
	}
}

func (manager *CommandManager) RegisterCommand(command string, handler func(context *Context)) {

	commandToRegister := Command{
		Command:        command,
		CommandHandler: handler,
	}
	manager.Commands = append(manager.Commands, commandToRegister)

}

func (manager *CommandManager) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	guild, _ := s.Guild(m.GuildID)

	input := strings.ToLower(strings.Split(m.Content, " ")[0])
	var commandImpl Command
	isCmd := false

	if strings.HasPrefix(input, database.GetGuildValue(guild, "prefix")) {
		for _, v := range manager.Commands {
			if strings.Contains(strings.ToLower(input), strings.ToLower(v.Command)) {
				isCmd = true
				commandImpl = v
			}
		}
		if isCmd == true {
			context := Context{
				Implementation: &commandImpl,
				Event:			m,
				Session:		s,
				Label:          input,
			}
			commandImpl.CommandHandler(&context)
		} else {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, "🚫")
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
		}
	}
}
