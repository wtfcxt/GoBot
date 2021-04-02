package bot

import (
	"GoBot/commands"
	"GoBot/util"
	"GoBot/util/embed"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Clear(ctx *commands.Context) {

	s := ctx.Session
	m := ctx.Event

	message := strings.Split(m.Message.Content, " ")

	if util.HasPermission(s, m, discordgo.PermissionManageMessages) {
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

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
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

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
				return
			} else {
				if method == "all" {
					unconverted, err := s.ChannelMessages(m.ChannelID, amount, "", "", "")

					if err != nil {
						s.ChannelMessageSend(m.ChannelID, err.Error())
					}
					converted := make([]string, len(unconverted))
					for i, m := range unconverted {
						converted[i] = m.ID
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go deleteMessages(converted, &wg, s, m)

					s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) for you.", "", embed.Green, nil))

					ch, err := s.Channel(m.ChannelID)
					if(err != nil) {
						s.ChannelMessageSend(m.ChannelID, err.Error())
					}

					lastID := ch.LastMessageID
					time.Sleep(3 * time.Second)
					s.ChannelMessageDelete(m.ChannelID, lastID)
				} else if method == "members" {
					fmt.Println(amount)
					unconverted, err := s.ChannelMessages(m.ChannelID, amount, "", "", "")

					if err != nil {
						s.ChannelMessageSend(m.ChannelID, err.Error())
					}
					converted := make([]string, len(unconverted))
					for i, m := range unconverted {
						if !m.Author.Bot {
							continue
						}
						converted[i] = m.ID
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go deleteMessages(converted, &wg, s, m)

					s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) from members.", "", embed.Green, nil))

					ch, err := s.Channel(m.ChannelID)
					if(err != nil) {
						s.ChannelMessageSend(m.ChannelID, err.Error())
					}

					lastID := ch.LastMessageID
					time.Sleep(3 * time.Second)
					s.ChannelMessageDelete(m.ChannelID, lastID)
				} else if method == "bots" {
					unconverted, err := s.ChannelMessages(m.ChannelID, amount, "", "", "")

					if err != nil {
						s.ChannelMessageSend(m.ChannelID, err.Error())
					}
					converted := make([]string, len(unconverted))
					for i, m := range unconverted {
						if m.Author.Bot {
							continue
						}
						converted[i] = m.ID
					}

					var wg sync.WaitGroup
					wg.Add(1)
					go deleteMessages(converted, &wg, s, m)

					s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) from bots.", "", embed.Green, nil))

					ch, err := s.Channel(m.ChannelID)
					if(err != nil) {
						s.ChannelMessageSend(m.ChannelID, err.Error())
					}

					lastID := ch.LastMessageID
					time.Sleep(3 * time.Second)
					s.ChannelMessageDelete(m.ChannelID, lastID)
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

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
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

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
			} else {
				unconverted, err := s.ChannelMessages(m.ChannelID, amount, "", "", "")

				if err != nil {
					s.ChannelMessageSend(m.ChannelID, err.Error())
				}
				converted := make([]string, len(unconverted))
				for i, m := range unconverted {
					converted[i] = m.ID
				}

				var wg sync.WaitGroup
				wg.Add(1)
				go deleteMessages(converted, &wg, s, m)

				s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbed("Deleted " + strconv.Itoa(amount) + " message(s).", "I successfully deleted **" + strconv.Itoa(amount) + "** message(s) for you.", "", embed.Green, nil))

				ch, err := s.Channel(m.ChannelID)
				if(err != nil) {
					s.ChannelMessageSend(m.ChannelID, err.Error())
				}

				lastID := ch.LastMessageID
				time.Sleep(3 * time.Second)
				s.ChannelMessageDelete(m.ChannelID, lastID)
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

			s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateEmbedFieldsOnly("An error occurred.", embed.Red, field))
		}
	} else {
		embed.NoPermsEmbed(s, m, "Manage Messages")
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
