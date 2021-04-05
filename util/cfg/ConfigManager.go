package cfg

import (
	"GoBot/util/logger"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {

	viper.SetDefault("general", map[string]string{"token": "empty", "prefix": "!", "debug": "disabled"})
	viper.SetDefault("database", map[string]string{"host": "localhost", "port": "27017", "username": "none", "password": "none"})

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		viper.WriteConfig()
		viper.SafeWriteConfig()
		viper.WriteConfigAs("config.json")
		logger.LogModule(logger.TypeInfo, "GoBot/Init", "Config created. Please enter your values and restart the bot.")
		os.Exit(0)
	}

	general := viper.GetStringMap("general")
	database := viper.GetStringMap("database")

	Token = general["token"].(string)
	Prefix = general["prefix"].(string)
	Debug = general["debug"].(string)

	Host = database["host"].(string)
	Port = database["port"].(string)
	Username = database["username"].(string)
	Password = database["password"].(string)

	logger.LogModule(logger.TypeInfo, "GoBot/Init", "Loaded config successfully.")

}