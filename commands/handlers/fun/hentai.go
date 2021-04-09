package fun

import (
	"GoBot/commands"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

/*
	This command is using https://meme-api.herokuapp.com/gimme
*/

func HentaiCommand(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	message := strings.Split(m.Message.Content, " ")
	channel, _ := s.Channel(m.ChannelID)

	err := s.ChannelTyping(m.ChannelID)

	if channel.NSFW == false {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "No NSFW channel.",
				Value:  "The content you requested is **N**ot **S**afe **F**or **W**ork. You need a NSFW channel in order to view it.",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
	} else {
		if len(message) == 2 {
			meme := GetHentaiSubreddit(s, m, message[1])
			if meme == "NSFW" {
				return
			}
			_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(ExtractValue(meme, "title"), "This image is from [r/" + message[1] + "](https://reddit.com/r/" + message[1] + ")", ExtractValue(meme, "url"), embed.Orange))
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
		} else if len(message) == 1 {
			meme := GetHentai(s, m)
			msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(ExtractValue(meme, "title"), ":smirk: - Generate a new Image", ExtractValue(meme, "url"), embed.Orange))
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
			s.MessageReactionAdd(m.ChannelID, msg.ID, "😏")
		}
	}

}

func GetHentaiNoThrow() string {
	url := "https://meme-api.herokuapp.com/gimme/" + randomHentaiSubreddit()

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	return string(content)
}

func GetHentai(s *discordgo.Session, m *discordgo.MessageCreate) string {
	url := "https://meme-api.herokuapp.com/gimme/" + randomHentaiSubreddit()

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

func GetHentaiSubreddit(s *discordgo.Session, m *discordgo.MessageCreate, subreddit string) string {

	url := "https://meme-api.herokuapp.com/gimme/" + subreddit

	channel, _ := s.Channel(m.ChannelID)

	resp, err := http.Get(url)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
	}

	if ExtractValue(string(content), "nsfw") == "true" && channel.NSFW == false {
		field := []*discordgo.MessageEmbedField{
			{
				Name:   "No NSFW channel.",
				Value:  "The content you requested is **N**ot **S**afe **F**or **W**ork. You need a NSFW channel in order to view it.",
				Inline: false,
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
		return "NSFW"

	} else {
		return string(content)
	}

}

func randomHentaiSubreddit() string {
	in := []string{"hentai", "hentai+memes", "hentaimemes"}
	randomIndex := rand.Intn(len(in))
	return in[randomIndex]
}