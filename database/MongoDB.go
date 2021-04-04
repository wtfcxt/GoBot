package database

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
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

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

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)

	if err != nil {
		panic(err)
	}

	logger.LogModule(logger.TypeInfo, "GoBot/Mongo", "Database Connection closed. [MongoDB]")
}

func AddGuild(client *mongo.Client, guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"id", guild.ID}, {"prefix", "!"}, {"warnch", "none"}, {"muterole", "none"}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedID
	logger.LogModule(logger.TypeInfo, "GoBot/Mongo", "Added server with ID \"" + guild.ID + "\" to database.")
}

func RemoveGuild(client *mongo.Client, guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, bson.D{{"id", guild.ID}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.DeletedCount
	logger.LogModule(logger.TypeInfo, "GoBot/Mongo", "Removed server with ID \"" + guild.ID + "\" from database.")
}

func GuildExists(client *mongo.Client, guild *discordgo.Guild) bool {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.Find(ctx, bson.D{{"id", guild.ID}})

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

func AddMember(client *mongo.Client, user *discordgo.User, guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err :=
		collection.InsertOne(ctx, bson.D{
		{"userid", user.ID},
		{"guildid", guild.ID},
		{"muted", false},
		{"warns", []string{}},
		})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedID
}

func AddAllMembers(client *mongo.Client, guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var collections []interface{}
	for i := range guild.Members {
		member := guild.Members[i]
		collections = append(collections, bson.D{
			{"userid", member.User.ID},
			{"guildid", guild.ID},
			{"muted", false},
			{"warns", []string{}},
		})
	}
	res, err := collection.InsertMany(ctx, collections)

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedIDs
}

func GetClient() *mongo.Client {
	return client
}

func ChangeGuildSetting(client *mongo.Client, guild *discordgo.Guild, setting string, value string) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"id", guild.ID}}, bson.D{{"$set", bson.D{{setting, value}}}}, options.Update())

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func ChangeMemberOption(client *mongo.Client, user *discordgo.User, option string, value string) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"userid", user.ID}}, bson.D{{"$set", bson.D{{option, value}}}}, options.Update())

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func ChangeMemberOptionBool(client *mongo.Client, user *discordgo.User, option string, value bool) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"userid", user.ID}}, bson.D{{"$set", bson.D{{option, value}}}}, options.Update())

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}


func AddWarning(client *mongo.Client, guild *discordgo.Guild, user *discordgo.User, reason string) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, bson.D{{"userid", user.ID}, {"guildid", guild.ID}}, bson.D{{"$push", bson.D{{"warns", []string{reason}}}}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}

func GetWarnings(client *mongo.Client, guild *discordgo.Guild, user *discordgo.User) {
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.Find(ctx, bson.D{{"userid", user.ID}, {"guildid", guild.ID}}, options.Find())
	if err != nil {
		logger.LogCrash(err)
	}

	var warns []string

	err = res.All(ctx, warns)
	if err != nil {
		logger.LogCrash(err)
	}

	fmt.Println(warns)

}

func GetSetting(client *mongo.Client,guild *discordgo.Guild, setting string) string {

	var result bson.D
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"id", guild.ID}}).Decode(&result)
	if err != nil {
		logger.LogCrash(err)
	}

	s := result.Map()[setting].(string)

	return s

}

func GetMemberValueString(client *mongo.Client, user *discordgo.User, value string) string {

	var result bson.D
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"userid", user.ID}}).Decode(&result)
	if err != nil {
		logger.LogCrash(err)
	}

	s := result.Map()[value].(string)

	return s

}

func GetMemberValueBoolean(client *mongo.Client, user *discordgo.User, value string) bool {

	var result bson.D
	collection := client.Database("gobot").Collection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"userid", user.ID}}).Decode(&result)
	if err != nil {
		logger.LogCrash(err)
	}

	s := result.Map()[value].(bool)

	return s

}