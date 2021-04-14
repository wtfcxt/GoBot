package flow

import (
	"GoBot-Recode/config"
	"context"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

var State string
var conf *oauth2.Config

func CallbackHandler(w http.ResponseWriter, r *http.Request) {

	conf = &oauth2.Config{
		RedirectURL: "http://" + config.WebHost + ":" + config.WebPort + "/auth/callback",
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}

	if r.FormValue("state") != State {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"InvalidStateException\", \"description\": \"The state does not match. Please try again later.\"}"))
		return
	}
	token, err := conf.Exchange(context.Background(), r.FormValue("code"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := conf.Client(context.Background(), token).Get("https://discordapp.com/api/users/@me")

	if err != nil || res.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(res.Status))
		}
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(body)
}