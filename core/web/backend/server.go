package backend

import (
	"GoBot-Recode/core/logger"
	"GoBot-Recode/core/web/backend/auth/flow"
	"fmt"
	"net/http"
	"os"
)

type WebServer struct {
	Address string
	Port    string
	TLS     bool
	ServeMux *http.ServeMux
}

var Server WebServer

func CreateServer(Address string, Port string, TLS bool) WebServer {
	var err error

	if Address == "" || Port == "" {
		logger.LogModule(logger.TypeError, "GoBot/Web", "Address and/or Port are missing. Can't continue.")
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir("html/assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/auth/login", flow.LoginHandler)
	mux.HandleFunc("/auth/callback", flow.CallbackHandler)
	if TLS {
		err = http.ListenAndServeTLS(Address + ":" + Port, "tls/certFile.crt", "tls/keyFile.key", mux)
	} else {
		err = http.ListenAndServe(Address + ":" + Port, mux)
	}

	if err != nil {
		panic(err)
	}

	Server = WebServer{
		Address:  Address,
		Port:	  Port,
		TLS:	  TLS,
		ServeMux: mux,
	}

	return Server
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\":\"alive\", \"app-name\": \"wtf.cxt.gobot\", \"app-description\": \"GoBot API Server\", \"version\":\"1.0\"}")
}