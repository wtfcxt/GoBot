package main

import (
	"GoBot-Recode/config"
	"GoBot-Recode/core/data"
	"GoBot-Recode/core/web"
	"GoBot-Recode/core/web/backend/auth"
	"GoBot-Recode/database/cache"
	"GoBot-Recode/database/mongo"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {

	fmt.Println(color.HiMagentaString("  _____     ___       __    _      __    __ \n / ___/__  / _ )___  / /_  | | /| / /__ / / \n/ (_ / _ \\/ _  / _ \\/ __/  | |/ |/ / -_) _ \\\n\\___/\\___/____/\\___/\\__/   |__/|__/\\__/_.__/\n                                            "))

	runDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	data.RunDirectory = runDir

	config.Init() // Initializing configuration
	mongo.Init()  // Initializing MongoDB Connection
	cache.Init()  // Initializing Cache
	auth.RandomState()
	web.Init()    // Initializing Webserver (creatin' files and stuff)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}