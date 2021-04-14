package mongo

import (
	"GoBot-Recode/config"
	"GoBot-Recode/core/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var client *mongo.Client

var ctx context.Context
var cancel context.CancelFunc

func Init() {

	logger.LogModuleNoNewline(logger.TypeInfo, "GoBot/Mongo", "Initializing MongoDB...")

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if config.Username == "none" && config.Password == "none" {
		clients, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port))
		err := clients.Ping(ctx, readpref.Primary())
		if err != nil {
			logger.AppendFail()
			panic(err)
		}

		logger.AppendDone()

		client = clients
	} else {
		clients, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "?authSource=admin"))
		err := clients.Ping(ctx, readpref.Primary())
		if err != nil {
			logger.AppendFail()
			panic(err)
		}

		logger.AppendDone()

		client = clients
	}
}