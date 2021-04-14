package flow

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, conf.AuthCodeURL(State), http.StatusTemporaryRedirect)
}