package bot

import (
	"GoBot/commands"
	"GoBot/util"
	"GoBot/util/embed"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Clear(ctx *commands.Context) {

	s := ctx.Session
	event := ctx.Event

	message := strings.Split(event.Message.Content, " ")

	if util.HasPermission(s, event, discordgo.PermissionManageMessages) {
		switch len(message) {
		case 3:
			method := message[1]
			amount, err := strconv.Atoi(message[2])
			if err != nil {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Invalid syntax.",
						Value:  "You didn't supply the amount of messages to delete.",
						Inline: false,
					},
					{
						Name: 	"Correct syntax",
						Value:	"`!clear [all|members|bots] <amount>` - You don't have to supply a method, it will default to all.",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			}

			if amount > 100 {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Integer too high.",
						Value:  "The amount of messages to delete can't be higher than 100.",
						Inline: false,
					},
					{
						Name: 	"Why?",
						Value:	"This is a discord limitation.",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			} else {
				if method == "all" {
					unconverted, err := s.ChannelMessages(event.ChannelID, amount, "", "", "")

					if err != nil {
						s.ChannelMessageSend(event.ChannelID, err.Error())
					}
					converted := make([]string, len(unconverted))
					for i, m := range unconverted {
						converted[i] = m.ID
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go deleteMessages(converted, &wg, s, event)

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) for you.", "", embed.Green, nil))

					ch, err := s.Channel(event.ChannelID)
					if(err != nil) {
						s.ChannelMessageSend(event.ChannelID, err.Error())
					}

					lastID := ch.LastMessageID
					time.Sleep(3 * time.Second)
					s.ChannelMessageDelete(event.ChannelID, lastID)
				} else if method == "members" {
					unconverted, err := s.ChannelMessages(event.ChannelID, amount, "", "", "")

					if err != nil {
						s.ChannelMessageSend(event.ChannelID, err.Error())
					}
					converted := make([]string, len(unconverted))
					for i, m := range unconverted {
						if !m.Author.Bot {
							converted[i] = m.ID
						}
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go deleteMessages(converted, &wg, s, event)

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) for you.", "", embed.Green, nil))

					ch, err := s.Channel(event.ChannelID)
					if(err != nil) {
						s.ChannelMessageSend(event.ChannelID, err.Error())
					}

					lastID := ch.LastMessageID
					time.Sleep(3 * time.Second)
					s.ChannelMessageDelete(event.ChannelID, lastID)
				} else if method == "bots" {
					unconverted, err := s.ChannelMessages(event.ChannelID, amount, "", "", "")

					if err != nil {
						s.ChannelMessageSend(event.ChannelID, err.Error())
					}
					converted := make([]string, len(unconverted))
					for i, msg := range unconverted {
						if msg.Author.Bot {
							converted[i] = msg.ID
						}
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go deleteMessages(converted, &wg, s, event)

					s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) from bots.", "", embed.Green, nil))

					ch, err := s.Channel(event.ChannelID)
					if(err != nil) {
						s.ChannelMessageSend(event.ChannelID, err.Error())
					}

					lastID := ch.LastMessageID
					time.Sleep(3 * time.Second)
					s.ChannelMessageDelete(event.ChannelID, lastID)
				}
			}
		case 2:
			amount, err := strconv.Atoi(message[1])
			if err != nil {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Invalid syntax.",
						Value:  "You didn't supply the amount of messages to delete.",
						Inline: false,
					},
					{
						Name: 	"Correct syntax",
						Value:	"`!clear [all|members|bots] <amount>` - You don't have to supply a method, it will default to all.",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			}

			if amount > 100 {
				field := []*discordgo.MessageEmbedField{
					{
						Name:   "Integer too high.",
						Value:  "The amount of messages to delete can't be higher than 100.",
						Inline: false,
					},
					{
						Name: 	"Why?",
						Value:	"This is a discord limitation.",
						Inline: false,
					},
				}

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			} else {
				unconverted, err := s.ChannelMessages(event.ChannelID, amount, "", "", "")

				if err != nil {
					s.ChannelMessageSend(event.ChannelID, err.Error())
				}
				converted := make([]string, len(unconverted))
				for i, m := range unconverted {
					converted[i] = m.ID
				}

				var wg sync.WaitGroup
				wg.Add(1)
				go deleteMessages(converted, &wg, s, event)

				s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) for you.", "", embed.Green, nil))

				ch, err := s.Channel(event.ChannelID)
				if(err != nil) {
					s.ChannelMessageSend(event.ChannelID, err.Error())
				}

				lastID := ch.LastMessageID
				time.Sleep(3 * time.Second)
				s.ChannelMessageDelete(event.ChannelID, lastID)
			}
		default:
			field := []*discordgo.MessageEmbedField{
				{
					Name:   "Invalid syntax.",
					Value:  "You didn't supply the amount of messages to delete.",
					Inline: false,
				},
				{
					Name: 	"Correct syntax",
					Value:	"`!clear [all|members|bots] <amount>` - You don't have to supply a method, it will default to all.",
					Inline: false,
				},
			}

			s.ChannelMessageSendEmbed(event.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
		}
	} else {
		embed.NoPermsEmbed(s, event, "Manage Messages")
	}

}

func deleteMessages(msgs []string, wg *sync.WaitGroup, s *discordgo.Session, m *discordgo.MessageCreate) {
	err := s.ChannelMessagesBulkDelete(m.ChannelID, msgs)
	if err != nil {
		embed.ThrowError(err.Error(), s, m)
		return
	}
	defer wg.Done()
}
