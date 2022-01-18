package mongo

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var Client *mongo.Client

func Init() (err error) {
	dsn := fmt.Sprintf("mongodb://%s:%d/?readPreference=primary&ssl=false",
		viper.GetString("mongo.host"),
		viper.GetInt("mongo.port"),
	)
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		zap.L().Fatal("mongodb initialization failed", zap.Error(err))
		return
	}
	return nil
}

func Close() {
	err := Client.Disconnect(context.Background())
	if err != nil {
		return
	}
}

// Mongo defines a mongo database
type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a new mongo dao
func NewMongo(db *mongo.Database, collection string) *Mongo {
	return &Mongo{
		col: db.Collection(collection),
	}
}

type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}
