package main

import (
	"GoBot/commands"
	"GoBot/commands/handlers/fun"
	"GoBot/commands/handlers/misc"
	"GoBot/commands/handlers/moderation/bot"
	"GoBot/database"
	bot2 "GoBot/events/bot"
	"GoBot/events/server"
	"GoBot/util/cfg"
	"GoBot/util/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.LogLogo()

	cfg.LoadConfig()              // Initializing Config
	database.Connect()                // Initializing Database
	manager := registerCommands() // Registering Commands

	bot, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		logger.LogCrash(err)
	}

	bot.Identify.Intents =
			discordgo.IntentsGuildPresences |
			discordgo.IntentsGuildBans |
			discordgo.IntentsGuildMessageReactions |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildMessageTyping |
			discordgo.IntentsGuildMembers |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsGuilds |
			discordgo.IntentsGuildVoiceStates

	bot.AddHandler(manager.MessageCreate)
	bot.AddHandler(server.GuildAdd)
	bot.AddHandler(server.GuildRemove)
	bot.AddHandler(bot2.ReadyEvent)
	// bot.AddHandler(events.Ready)
	// bot.AddHandler(events.GuildJoin)

	if cfg.Token == "empty" {
		logger.LogModule(logger.TypeError, "GoBot/Init", "You didn't enter a token. Can't continue.")
		os.Exit(1)
	}

	err = bot.Open()
	if err != nil {
		logger.LogCrash(err)
	}

	fmt.Print("\n")
	logger.LogModule(logger.TypeInfo, "GoBot/Init", "Bot running. (Quit using Ctrl + C)")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
	database.Disconnect()
}

func registerCommands() commands.CommandManager {
	manager := commands.NewCommandManager()

	manager.RegisterCommand("clear", bot.Clear)
	manager.RegisterCommand("warns", bot.Warnings)
	manager.RegisterCommand("mute", bot.Mute)

	manager.RegisterCommand("settings", misc.Settings)
	manager.RegisterCommand("info", misc.Info)
	manager.RegisterCommand("help", misc.Help)

	manager.RegisterCommand("meme", fun.MemeCommand)

	logger.LogModule(logger.TypeInfo, "GoBot/Init", "Registered commands.")
	return manager
}