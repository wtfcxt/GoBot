package embed

import "github.com/bwmarrin/discordgo"

var (
	Red = 16711685
	Green = 3657238
	Yellow = 14735630
	Orange = 13929741
)

func CreateEmbed(title string, description string, thumbnail string, colour int, fields []*discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	embedThumbnail := discordgo.MessageEmbedThumbnail{
		URL: thumbnail,
	}

	footer := discordgo.MessageEmbedFooter{
		Text: "(c) 2021 - GoBot X | Made by cxt#1234",
	}

	return &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       colour,
		Thumbnail:   &embedThumbnail,
		Fields:      fields,
		Footer:		 &footer,
	}
}

func CreateEmbedFieldsOnly(title string, colour int, fields []*discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	footer := discordgo.MessageEmbedFooter{
		Text: "(c) 2021 - GoBot X | Made by cxt#1234",
	}
	return &discordgo.MessageEmbed{
		Title:       title,
		Color:       colour,
		Fields:      fields,
		Footer:		 &footer,
	}
}

func ThrowError(error string, s *discordgo.Session, m *discordgo.MessageCreate) {
	field := []*discordgo.MessageEmbedField{
		{
			Name:   "An error occurred.",
			Value:  "I couldn't perform the action you wanted me to perform.",
			Inline: false,
		},
		{
			Name: 	"Error",
			Value:	"`" + error + "`",
			Inline: false,
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, CreateEmbedFieldsOnly("An error occurred.", Red, field))
}

func NoPermsEmbed(s *discordgo.Session, m*discordgo.MessageCreate, perm string) {
	field := []*discordgo.MessageEmbedField{
		{
			Name:   "No permissions.",
			Value:  "You do not have necessary permissions in order to execute this command.",
			Inline: false,
		},
		{
			Name: 	"Permission needed",
			Value:	"`" + perm + "`",
			Inline: false,
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, CreateEmbedFieldsOnly("An error occurred.", Red, field))
}