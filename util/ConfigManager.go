package util

import (
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {

	viper.SetDefault("mongodb", map[string]string{"host": "localhost", "port": "27017", "username": "none", "password": "none"})
	viper.SetDefault("settings", map[string]string{"token": "empty", "prefix": "!"})
	viper.SetDefault("runtime", map[string]int{"max_threads": 1, "max_goroutines": 250000})
	viper.SetDefault("debug", map[string]bool{"debug": false})

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		viper.WriteConfig()
		viper.SafeWriteConfig()
		viper.WriteConfigAs("config.json")
		LogModule(TypeInfo, "GoBot/Config", "Config created. Please enter your values and restart the bot.")
		os.Exit(0)
	}

	settings := viper.GetStringMap("settings")
	runtime := viper.GetStringMap("runtime")
	debug := viper.GetStringMap("debug")
	mongodb := viper.GetStringMap("mongodb")

	Token = settings["token"].(string)
	Prefix = settings["prefix"].(string)
	Debug = debug["debug"].(bool)
	MaxThreads = runtime["max_threads"].(float64)

	Host = mongodb["host"].(string)
	Port = mongodb["port"].(string)
	Username = mongodb["username"].(string)
	Password = mongodb["password"].(string)

}
