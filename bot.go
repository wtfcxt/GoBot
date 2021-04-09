package main

import (
	"GoBot/commands"
	"GoBot/commands/handlers/admin"
	"GoBot/commands/handlers/fun"
	"GoBot/commands/handlers/misc"
	"GoBot/commands/handlers/moderation/bot"
	"GoBot/commands/handlers/moderation/server"
	"GoBot/database"
	bot2 "GoBot/events/bot"
	"GoBot/events/guild"
	"GoBot/util"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

func main() {
	util.LogLogo()

	util.LoadConfig()              // Initializing Config

	if util.Debug {
		util.LogModule(util.TypeWarn, "GoBot/Debug", "############################################################")
		util.LogModule(util.TypeWarn, "GoBot/Debug", "The bot is running in debug mode. DO NOT USE IN PRODUCTION!")
		util.LogModule(util.TypeWarn, "GoBot/Debug", "############################################################")
	}

	runtime.GOMAXPROCS(int(util.MaxThreads))
	util.LogModule(util.TypeInfo, "GoBot/Runtime", "Limiting Maximum CPU Threads to " + strconv.Itoa(int(util.MaxThreads)) + "...")
	database.Connect()            // Initializing Database
	manager := registerCommands() // Registering Commands

	bot, err := discordgo.New("Bot " + util.Token) // Creating the bot
	if err != nil {
		util.LogCrash(err)
	}

	bot.Identify.Intents = discordgo.IntentsAll // This makes the bot use all Intents. (discord.com/developers)

	bot.AddHandler(manager.MessageCreate) // When a command has been executed, this event is called
	bot.AddHandler(guild.Add)             // When the bot registers a new guild, this event is called
	bot.AddHandler(bot2.ReadyEvent)       // When the bot is ready, this event is called
	bot.AddHandler(guild.UserJoin)        // When a new user joins, this event is called
	bot.AddHandler(bot2.AddReaction)	  // When a user adds a reaction to a message

	// If the token in the configuration file is empty, won't continue.
	if util.Token == "empty" {
		util.LogModule(util.TypeError, "GoBot/Bot", "You didn't enter a token. Can't continue.")
		os.Exit(1)
	}

	// If an error occurs whilst trying to open the bot, it will crash.
	err = bot.Open()
	if err != nil {
		util.LogCrash(err)
	}

	fmt.Print("\n")
	util.LogModule(util.TypeInfo, "GoBot/Bot", "Bot running. (Quit using Ctrl + C)") // The bot is now running.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()           // Closing the bot's connection
	fmt.Println("\n")
	util.LogModule(util.TypeInfo, "GoBot/Bot", "Shutting down Bot...")
	database.Disconnect() // Disconnecting from Database
}

func registerCommands() commands.CommandManager {
	manager := commands.NewCommandManager()

	/*
		Admin commands
	*/

	manager.RegisterCommand("settings", admin.Settings)

	/*
		Moderation commands
	*/

	manager.RegisterCommand("clear", bot.Clear)
	manager.RegisterCommand("warn", bot.Warn)
	manager.RegisterCommand("warns", bot.Warnings)
	manager.RegisterCommand("mute", bot.Mute)
	manager.RegisterCommand("unmute", bot.Unmute)

	manager.RegisterCommand("kick", server.Kick)
	manager.RegisterCommand("ban", server.Ban)

	/*
		Miscellaneous commands
	*/

	manager.RegisterCommand("info", misc.Info)
	manager.RegisterCommand("help", misc.Help)

	/*
		Fun commands
	*/

	manager.RegisterCommand("meme", fun.MemeCommand)
	manager.RegisterCommand("nsfw", fun.NSFWCommand)
	manager.RegisterCommand("arguments", fun.ArgumentCommand)
	manager.RegisterCommand("hack", fun.Hack)
	manager.RegisterCommand("hentai", fun.HentaiCommand)

	util.LogModule(util.TypeInfo, "GoBot/Commands", "Registered commands.")
	return manager
}
