package database

import (
	"GoBot/util"
	"context"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

var client *mongo.Client
var wg2 sync.WaitGroup

/*
	This function is used for connecting to the database.
*/
func Connect() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if util.Username == "none" && util.Password == "none" {
		clients, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+util.Host+":"+util.Port))
		err := clients.Ping(ctx, readpref.Primary())
		if err != nil {
			util.LogCrash(err)
		}

		util.LogModule(util.TypeInfo, "GoBot/Database", "Connected to database. [MongoDB]")

		client = clients
	} else {
		clients, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+util.Username+":"+util.Password+"@"+util.Host+":"+util.Port+"?authSource=admin"))
		err := clients.Ping(ctx, readpref.Primary())
		if err != nil {
			util.LogCrash(err)
		}

		util.LogModule(util.TypeInfo, "GoBot/Database", "Connected to database. [MongoDB]")

		client = clients
	}
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

	util.LogModule(util.TypeInfo, "GoBot/Mongo", "Database Connection closed. [MongoDB]")
}

/*
	This function is used for registering a new guild.
*/
func RegisterGuild(guild *discordgo.Guild, session *discordgo.Session, wg *sync.WaitGroup) {

	defer wg.Done()

	util.LogModule(util.TypeDebug, "GoBot/Debug", "Trying to register the current guild...")
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"guild_id", guild.ID}, {"settings", map[string]string{"prefix": "!", "warn_channel_id": "none", "mute_role_id": "none"}}, {"logger", map[string]string{"enabled": "no", "channel": "none", "messages": "no", "joins": "no", "leaves": "no", "role": "none", "punishments": "no"}}})

	if err != nil {
		util.LogCrash(err)
	}

	_ = res.InsertedID

	util.LogModule(util.TypeDebug, "GoBot/Debug", "Inserted Guild into MongoDB.")

	var collections []interface{}

	for i := range guild.Members {
		user := guild.Members[i].User
		if !UserExists(user, guild) && user != session.State.User {
			collections = append(collections, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}, {"muted", false}, {"warns", []string{""}}})
		}
	}
	CreateManyUsers(collections)

}

func GuildExists(guild *discordgo.Guild) bool {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.Find(ctx, bson.D{{"guild_id", guild.ID}})

	if err != nil {
		util.LogCrash(err)
	}

	var episodes []bson.D
	if err = res.All(ctx, &episodes); err != nil {
		util.LogCrash(err)
	}

	if len(episodes) == 0 {
		return false
	} else {
		return true
	}

}

func UserExists(user *discordgo.User, guild *discordgo.Guild) bool {
	var result bson.D

	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}, options.FindOne()).Decode(&result)
	if err != nil {
		return false
	}

	if result == nil {
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
		util.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func GetGuildValue(guild *discordgo.Guild, option string) string {
	var result bson.D

	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"guild_id", guild.ID}}).Decode(&result)
	if err != nil {
		util.LogCrash(err)
	}

	s := result.Map()["settings"].(bson.D).Map()[option].(string)

	return s
}

func GetGuildLoggerValue(guild *discordgo.Guild, option string) string {
	var result bson.D

	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"guild_id", guild.ID}}).Decode(&result)
	if err != nil {
		util.LogCrash(err)
	}

	s := result.Map()["logger"].(bson.D).Map()[option].(string)

	return s
}

func ChangeGuildLoggerValue(guild *discordgo.Guild, option string, value string) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"guild_id", guild.ID}}, bson.D{{"$set", bson.D{{"logger." + option, value}}}}, options.Update())

	if err != nil {
		util.LogCrash(err)
	}

	_ = res.ModifiedCount
}

/*
	User Part
*/

func ChangeUserValueString(user *discordgo.User, guild *discordgo.Guild, option string, value string) {
	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}, bson.D{{"$set", bson.D{{option, value}}}}, options.Update())

	if err != nil {
		util.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func ChangeUserValueBool(user *discordgo.User, guild *discordgo.Guild, option string, value bool) {
	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}, bson.D{{"$set", bson.D{{option, value}}}}, options.Update())

	if err != nil {
		util.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func CreateUser(user *discordgo.User, guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}, {"muted", false}, {"warns", []string{}}})

	if err != nil {
		util.LogCrash(err)
	}

	_ = res.InsertedID
}

func CreateManyUsers(users []interface{}) {
	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertMany(ctx, users, options.InsertMany())

	if err != nil {
		util.LogModule(util.TypeWarn, "GoBot/Database", "A non-critical MongoDB error occurred: " + err.Error())
	}
}

func GetUserValueString(user *discordgo.User, guild *discordgo.Guild, option string) string {
	var result bson.D

	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}).Decode(&result)
	if err != nil {
		util.LogCrash(err)
	}

	s := result.Map()[option].(string)

	return s
}

func GetUserValueJson(user *discordgo.User, guild *discordgo.Guild, option string) bson.D {
	var result bson.D

	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}).Decode(&result)
	if err != nil {
		util.LogCrash(err)
	}

	s := result.Map()[option].(bson.D)

	return s
}

func GetUserValueBool(user *discordgo.User, guild *discordgo.Guild, option string) bool {
	var result bson.D

	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}).Decode(&result)
	if err != nil {
		util.LogCrash(err)
	}

	s := result.Map()[option].(bool)

	return s
}

func AddWarning(user *discordgo.User, guild *discordgo.Guild, reason string) {

	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.UpdateOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}, bson.D{{"$push", bson.D{{"warns", reason}}}}, options.Update())

	if err != nil {
		util.LogCrash(err)
	}

	_ = res.MatchedCount

}

func RemoveWarning(user *discordgo.User, guild *discordgo.Guild, id string) {

}

func GetWarnings(user *discordgo.User, guild *discordgo.Guild) bson.A {

	var result bson.D

	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"user_id", user.ID}, {"guild_id", guild.ID}}, options.FindOne()).Decode(&result)
	if err != nil {
		util.LogCrash(err)
	}

	s := result.Map()["warns"].(bson.A)

	return s

}
