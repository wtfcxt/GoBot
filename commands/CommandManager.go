package commands

import (
	"GoBot/util/cfg"
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

	input := strings.ToLower(strings.Split(m.Content, " ")[0])
	var commandImpl Command
	isCmd := false

	if strings.HasPrefix(input, manager.Prefix) {
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
		}
	}
}
