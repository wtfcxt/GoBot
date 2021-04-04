package fun

import (
	"GoBot/commands"
	"GoBot/util/embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

/*
	This command is using https://meme-api.herokuapp.com/gimme
 */

func MemeCommand(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	err := s.ChannelTyping(m.ChannelID)

	meme := GetMeme(s, m)
	fmt.Println(extractValue(meme, "url"))
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(extractValue(meme, "title"), "I found this on reddit.", extractValue(meme, "url"), embed.Orange))
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

}

func GetMeme(s *discordgo.Session, m *discordgo.MessageCreate) string {
	url := "https://meme-api.herokuapp.com/gimme"

	resp, err := http.Get(url)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	return string(content)
}

func extractValue(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.SplitN(match, ":", 2)
	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}