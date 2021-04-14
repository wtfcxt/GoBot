package backend

import (
	"GoBot-Recode/core/data"
	"GoBot-Recode/core/logger"
	"os"
	"sync"
)

func WriteFiles(wg *sync.WaitGroup) {

	defer wg.Done()

	logger.LogModuleNoNewline(logger.TypeInfo, "GoBot/Web", "Initializing webserver...")

	if _, err := os.Stat(data.RunDirectory + "/html/"); err == nil {} else if os.IsNotExist(err) {
		err := os.MkdirAll(data.RunDirectory + "/html/", os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		logger.AppendFail()
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		panic(err)
	}


	if _, err := os.Stat(data.RunDirectory + "/html/assets/"); err == nil {} else if os.IsNotExist(err) {
		err := os.MkdirAll(data.RunDirectory + "/html/assets/", os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		logger.AppendFail()
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		panic(err)
	}

	if _, err := os.Stat(data.RunDirectory + "/tls/"); err == nil {} else if os.IsNotExist(err) {
		err := os.MkdirAll(data.RunDirectory + "/tls/", os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		logger.AppendFail()
		logger.LogModule(logger.TypeError, "GoBot/Web", "Unable to check if file exists or does not exist. Panic!")
		panic(err)
	}

	logger.AppendDone()

}