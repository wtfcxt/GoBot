package web

import (
	"GoBot-Recode/logger"
	"net/http"
	"os"
)

type WebServer struct {
	Address string
	Port    string
	TLS     bool
}

var Server WebServer

func CreateServer(Address string, Port string, TLS bool) WebServer {
	var err error

	if Address == "" || Port == "" {
		logger.LogModule(logger.TypeError, "GoBot/Web", "Address and/or Port are missing. Can't continue.")
		os.Exit(1)
	}

	if TLS {
		mux := http.NewServeMux()

		mux.HandleFunc("/", handler)
		err = http.ListenAndServeTLS(Address + ":" + Port, "tls/certFile.crt", "tls/keyFile.key", mux)
	} else {
		mux := http.NewServeMux()

		mux.HandleFunc("/", handler)
		err = http.ListenAndServe(Address + ":" + Port, mux)
	}

	if err != nil {
		logger.LogCrash(err)
	}

	Server = WebServer{
		Address: Address,
		Port:	 Port,
		TLS:	 TLS,
	}

	return Server
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
	    logger.LogCrash(err)
	}
}