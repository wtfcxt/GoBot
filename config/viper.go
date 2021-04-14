package config

import (
	"GoBot-Recode/core/data"
	"GoBot-Recode/core/logger"
	"github.com/spf13/viper"
	"os"
)

var (
	WebHost string
	WebPort string
	WebTLS string

	Token string
	Prefix string
	ClientID string
	ClientSecret string

	Debug bool
	MaxThreads float64

	Host string
	Port string
	Username string
	Password string
)

func Init() {

	logger.LogModuleNoNewline(logger.TypeInfo, "GoBot/Config", "Initializing config...")

	if _, err := os.Stat(data.RunDirectory + "/config/"); err == nil {} else if os.IsNotExist(err) {
		err := os.MkdirAll(data.RunDirectory + "/config/", os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		logger.AppendFail()
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		panic(err)
	}

	viper.SetDefault("webserver", map[string]string{"host": "127.0.0.1", "port": "8080", "encryption": "false"})
	viper.SetDefault("mongodb", map[string]string{"host": "localhost", "port": "27017", "username": "none", "password": "none"})
	viper.SetDefault("settings", map[string]string{"clientid": "", "clientsecret": "", "token": "empty", "prefix": "!"})
	viper.SetDefault("runtime", map[string]int{"max_threads": 1, "max_goroutines": 250000})
	viper.SetDefault("debug", map[string]bool{"debug": false})

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(data.RunDirectory + "/config/")
	if err := viper.ReadInConfig(); err != nil {
		viper.WriteConfig()
		viper.SafeWriteConfigAs(data.RunDirectory + "/config/config.json")
		viper.WriteConfigAs(data.RunDirectory + "/config/config.json")
		logger.AppendDone()
		logger.LogModule(logger.TypeInfo, "GoBot/Config", "Configuration file created. Please enter your values and restart the bot.")
		os.Exit(0)
	}

	webserver := viper.GetStringMap("webserver")
	settings := viper.GetStringMap("settings")
	runtime := viper.GetStringMap("runtime")
	debug := viper.GetStringMap("debug")
	mongodb := viper.GetStringMap("mongodb")

	WebHost = webserver["host"].(string)
	WebPort = webserver["port"].(string)
	WebTLS = webserver["encryption"].(string)

	Token = settings["token"].(string)
	Prefix = settings["prefix"].(string)
	ClientID = settings["clientid"].(string)
	ClientSecret = settings["clientsecret"].(string)

	Debug = debug["debug"].(bool)
	MaxThreads = runtime["max_threads"].(float64)

	Host = mongodb["host"].(string)
	Port = mongodb["port"].(string)
	Username = mongodb["username"].(string)
	Password = mongodb["password"].(string)

	logger.AppendDone()

}
