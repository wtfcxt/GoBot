package new

import (
	"GoBot/util/cfg"
	"GoBot/util/logger"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

/*
	This function is used for connecting to the database.
 */
func Connect() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clients, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port))

	err := clients.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.LogCrash(err)
	}

	logger.LogModule(logger.TypeInfo, "GoBot/Init", "Connected to database. [MongoDB]")

	client = clients
}

/*
	This function is used for disconnecting from the database.
 */
func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}

	logger.LogModule(logger.TypeInfo, "GoBot/Mongo", "Database Connection closed. [MongoDB]")
}

/*
	This function is used for registering a new guild.
 */
func RegisterGuild(guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"guild_id", guild.ID}, {"settings", map[string]string{"prefix": "!", "warn_channel_id": "none", "mute_role_id": "none"}}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedID

	RegisterUserBulk(guild)
}

/*
	This function is used for registering a new user that isn't currently in the database.
 */
func RegisterUser(guild *discordgo.Guild, user *discordgo.User) {
	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"user_id", user.ID}, {"guilds", []bson.D{{}}}})
	AddGuildToUser(guild, user)

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedID
}

/*
	This function is for registering many new users which currently are not in the database.
 */
func RegisterUserBulk(guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var collections []interface{}
	for i := range guild.Members {
		user := guild.Members[i].User
		if !UserExists(user) {
			collections = append(collections, bson.D{{"user_id", user.ID}, {"guilds", []bson.D{{}}}})
		}
	}

	if collections != nil {
		res, err := collection.InsertMany(ctx, collections)
		if err != nil {
			logger.LogCrash(err)
		}

		_ = res.InsertedIDs

		for i := range guild.Members {
			user := guild.Members[i].User
			AddGuildToUser(guild, user)
		}
	} else {
		for i := range guild.Members {
			user := guild.Members[i].User
			if !IsUserInGuild(guild, user) {
				fmt.Println("should add rn")
				AddGuildToUser(guild, user)
			}
		}
	}
 }

/*
	This function is used for registering a guild in a user's document.
 */
func AddGuildToUser(guild *discordgo.Guild, user *discordgo.User) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.UpdateOne(ctx,
		bson.D{{"id", user.ID}},
		bson.D{{"$push",bson.D{{"guilds", bson.D{
			{"guild", guild.ID},
			{"muted", false},
			{"warns", []bson.D{{}}}}}}}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func GuildExists(guild *discordgo.Guild) bool {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.Find(ctx, bson.D{{"guild_id", guild.ID}})

	if err != nil {
		logger.LogCrash(err)
	}

	var episodes []bson.D
	if err = res.All(ctx, &episodes); err != nil {
		logger.LogCrash(err)
	}

	if len(episodes) == 0 {
		return false
	} else {
		return true
	}

}

func UserExists(user *discordgo.User) bool {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.Find(ctx, bson.D{{"user_id", user.ID}})

	if err != nil {
		logger.LogCrash(err)
	}

	var episodes []bson.D
	if err = res.All(ctx, &episodes); err != nil {
		logger.LogCrash(err)
	}

	if len(episodes) == 0 {
		return false
	} else {
		return true
	}

}

func ChangeGuildValue(guild *discordgo.Guild, option string, value string) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"guild_id", guild.ID}}, bson.D{{"$set", bson.D{{"settings." + option, value}}}}, options.Update())

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func ChangeUserValue(user *discordgo.User, option string, value string) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"id", user.ID}}, bson.D{{"$set", bson.D{{"guilds." + option, value}}}}, options.Update())

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}

/*func GetUserValueString(user *discordgo.User, option string) string {

}*/

/*func GetUserValueBool(user *discordgo.User, option string) bool {
	var result bson.D
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}}).Decode(&result)
	if err != nil {
		logger.LogCrash(err)
	}

	s := result.Map()["guilds"].(bson.D).Map()[""]
}*/

func GetGuildValue(guild *discordgo.Guild, option string) string {
	var result bson.D

	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"guild_id", guild.ID}}).Decode(&result)
	if err != nil {
		logger.LogCrash(err)
	}

	// fmt.Println(result.Map()["settings"].(bson.D).Map()[option].(string))
	s := result.Map()["settings"].(bson.D).Map()[option].(string)

	return s
}

func IsUserInGuild(guild *discordgo.Guild, user *discordgo.User) bool {

	/* TODO: Abfrage ob Nutzer in der Gilde ist
	    -> wenn der Nutzer bereits die Gilde in seinem Profil hat, returnen
	    -> wenn der Nutzer nicht in der derzeitigen Gilde ist, fortfahren
	*/

	var result bson.D

	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}}).Decode(&result)
	if err != nil {
		logger.LogCrash(err)
	}

	s := result.Map()["guilds"].(bson.A)
	var found bool = false
	if s == nil {
		found = false
		fmt.Println("skipped")
		goto skip
	}
	for i := range s {
		i2 := s[i]
		fmt.Println(i2)
		if i2 == guild.ID {
			found = true
		}
	}

	skip:
	fmt.Println(found)
	return found

}