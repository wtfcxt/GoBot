package new

import (
	"GoBot/util/cfg"
	"GoBot/util/logger"
	"context"
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
	res, err := collection.InsertOne(ctx, bson.D{{"id", guild.ID}, {"settings", map[string]string{"prefix": "!", "warnch": "none", "muterole": "none"}}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedID
}

/*
	This function is used for unregistering a guild because they for example removed the bot from the server.
 */
func DeregisterGuild(guild *discordgo.Guild) {
	collection := client.Database("gobot").Collection("guilds")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, bson.D{{"id", guild.ID}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.DeletedCount
}

/*
	This function is used for registering a new user that isn't currently in the database.
 */
func RegisterUser(user *discordgo.User) {
	collection := client.Database("gobot").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"id", user.ID}, {"guilds", []bson.E{}}})

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
		collections = append(collections, bson.D{{"id", user.ID}, {"guilds", []bson.E{}}})
	}
	res, err := collection.InsertMany(ctx, collections)

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.InsertedIDs
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
		bson.D{{"$push",bson.D{
		{"guild", guild.ID},
		{"muted", false},
		{"warns", []string{}}}}})

	if err != nil {
		logger.LogCrash(err)
	}

	_ = res.ModifiedCount
}