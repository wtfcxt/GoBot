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

func NSFWCommand(ctx *commands.Context) {

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
			meme := GetNSFWSubreddit(s, m, message[1])
			if meme == "NSFW" {
				return
			}
			_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(ExtractValue(meme, "title"), "This image is from [r/" + message[1] + "](https://reddit.com/r/" + message[1] + ")", ExtractValue(meme, "url"), embed.Orange))
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
		} else if len(message) == 1 {
			meme := GetNSFW(s, m)
			_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(ExtractValue(meme, "title"), "This image is **NSFW**.", ExtractValue(meme, "url"), embed.Orange))
			if err != nil {
				embed.ThrowError(err.Error(), s, m)
			}
		}
	}

}

func GetNSFW(s *discordgo.Session, m *discordgo.MessageCreate) string {
	url := "https://meme-api.herokuapp.com/gimme/" + randomNSFWSubreddit()

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

func GetNSFWSubreddit(s *discordgo.Session, m *discordgo.MessageCreate, subreddit string) string {

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

func randomNSFWSubreddit() string {
	in := []string{"nsfw", "nsfw411", "omgbeckylookathiscock", "nudes", "legalteens", "amateur", "tiktoknsfw", "nsfw_wtf", "hentai", "hentai+memes", "gayporn"}
	randomIndex := rand.Intn(len(in))
	return in[randomIndex]
}