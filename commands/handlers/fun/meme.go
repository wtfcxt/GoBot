package fun

import (
	"GoBot/commands"
	"GoBot/util/embed"
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

	message := strings.Split(m.Message.Content, " ")

	_ = s.ChannelTyping(m.ChannelID)

	if len(message) >= 2 {
		meme := GetMemeSubreddit(s, m, message[1])
		if meme == "NSFW" {
			return
		}
		if meme == "" {
			return
		}

		title := ExtractValue(meme, "title")
		memeURL := ExtractValue(meme, "url")

		if memeURL == "" {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Couldn't get image.",
					Value:  "The content you searched for does not exist.",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			return
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(title, "This image is from [r/" + message[1] + "](https://reddit.com/r/" + message[1] + ")", memeURL, embed.Orange))
		if err != nil {
			embed.ThrowError(err.Error(), s, m)
		}

	} else if len(message) == 1 {

		meme := GetMeme(s, m)
		title := ExtractValue(meme, "title")
		memeURL := ExtractValue(meme, "url")

		msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedImage(title, ":arrows_counterclockwise: - Generate a new Image", memeURL, embed.Orange))
		if err != nil {
			embed.ThrowError(err.Error(), s, m)
		}

		s.MessageReactionAdd(m.ChannelID, msg.ID, "🔄")

	}

}

func GetMemeNoThrow() string {
	url := "https://meme-api.herokuapp.com/gimme"

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	return string(content)
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

func GetMemeSubreddit(s *discordgo.Session, m *discordgo.MessageCreate, subreddit string) string {

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
		if err != nil {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Couldn't get image.",
					Value:  "The content you searched for does not exist.",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			return ""
		}
		return "NSFW"

	} else {
		if err != nil {
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Couldn't get image.",
					Value:  "The content you searched for does not exist.",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			return ""
		}
		return string(content)
	}

}

func ExtractValue(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.SplitN(match, ":", 2)
	if len(keyValMatch) == 1 {
		return ""
	}
	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}