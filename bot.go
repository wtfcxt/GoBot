package main

import (
	"GoBot-Recode/logger"
	"GoBot-Recode/web"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	must(WriteFiles(&wg))
	wg.Wait()

	go web.CreateServer("127.0.0.1", "8080", false)
	if !web.Server.TLS {
		logger.LogModule(logger.TypeWarn, "GoBot/Web", "This server is not running in TLS mode. Please don't use it in production.")
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func WriteFiles(wg *sync.WaitGroup) {

	defer wg.Done()

	runDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	if _, err := os.Stat("/html/"); err == nil {
	} else if os.IsNotExist(err) {
		os.MkdirAll(runDir + "/html/", os.ModePerm)
	} else {
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		logger.LogCrash(err)
	}

	if _, err := os.Stat("/tls/"); err == nil {
	} else if os.IsNotExist(err) {
		os.MkdirAll(runDir + "/tls/", os.ModePerm)
	} else {
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		logger.LogCrash(err)
	}

	if _, err := os.Stat("/config/"); err == nil {
	} else if os.IsNotExist(err) {
		os.MkdirAll(runDir + "/config/", os.ModePerm)
	} else {
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		logger.LogCrash(err)
	}

	if _, err := os.Stat("/html/index.html"); err == nil {
	} else if os.IsNotExist(err) {
		example := []byte("<h1>Webserver running.</h1>")
		os.WriteFile(runDir + "/html/index.html", example, os.ModePerm)
	} else {
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		logger.LogCrash(err)
	}

}