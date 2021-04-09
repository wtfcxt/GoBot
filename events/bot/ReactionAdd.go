package bot

import (
	"GoBot/commands/handlers/fun"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
)

func AddReaction(s *discordgo.Session, event *discordgo.MessageReactionAdd) {
	user, err := s.User(event.MessageReaction.UserID)
	message, err := s.ChannelMessage(event.ChannelID ,event.MessageID)
	channel, err := s.Channel(event.ChannelID)

	if err != nil {
		util.LogCrash(err)
	}

	if user.ID != s.State.User.ID {
		if event.MessageReaction.Emoji.Name == "🔄" && message.Reactions[0].Me {
			s.MessageReactionRemove(event.ChannelID, message.ID, "🔄", user.ID)
			newMeme := fun.GetMemeNoThrow()
			if newMeme == "NSFW" {
				return
			}
			_, err := s.ChannelMessageEditEmbed(event.ChannelID, message.ID, embed.CreateEmbedImage(fun.ExtractValue(newMeme, "title"), ":arrows_counterclockwise: - Generate a new Image", fun.ExtractValue(newMeme, "url"), embed.Orange))
			if err != nil {
				util.LogCrash(err)
			}
		} else if event.MessageReaction.Emoji.Name == "😏" && message.Reactions[0].Me && channel.NSFW {
			s.MessageReactionRemove(event.ChannelID, message.ID, "😏", user.ID)
			newMeme := fun.GetHentaiNoThrow()
			_, err := s.ChannelMessageEditEmbed(event.ChannelID, message.ID, embed.CreateEmbedImage(fun.ExtractValue(newMeme, "title"), ":smirk: - Generate a new Image", fun.ExtractValue(newMeme, "url"), embed.Orange))
			if err != nil {
				util.LogCrash(err)
			}
		}
	}
}