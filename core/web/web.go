package web

import (
	"GoBot-Recode/config"
	"GoBot-Recode/core/logger"
	"GoBot-Recode/core/web/backend"
	"sync"
)

func Init() {
	var wg sync.WaitGroup
	wg.Add(1); go backend.WriteFiles(&wg); wg.Wait()

	tls := false
	if config.WebTLS != "false" {
		tls = true
	}

	go backend.CreateServer(config.WebHost, config.WebPort, tls)
	if !backend.Server.TLS {
		logger.LogModule(logger.TypeWarn, "GoBot/Web", "This server is not running in TLS mode. Please don't use it in production.")
	}
}
