package mongo_impl

import (
	"context"
	"fmt"
	"gin_skeleton/conf"
	"gin_skeleton/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDb struct {
	Client *mongo.Client
	MDB    *mongo.Database
	Conf   *conf.Config
}

func NewMongoDb() *MongoDb {
	return &MongoDb{
		Conf: global.ConfObj,
	}
}

func (m *MongoDb) getMongoClient() *mongo.Client {

	MongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		m.Conf.Mongo.UserName, m.Conf.Mongo.Password, m.Conf.Mongo.Host, m.Conf.Mongo.Port,
	)
	fmt.Println(MongoURL)
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongoDB")
	return client
}

func (m *MongoDb) GetMongoClient() *mongo.Client {
	return m.getMongoClient()
}

func (m *MongoDb) GetMongoDb() *mongo.Database {
	return m.getMongoClient().Database(m.Conf.Mongo.Database)
}

func InitMongoDb() *MongoDb {
	mD := NewMongoDb()
	mD.Client = mD.getMongoClient()
	mD.MDB = mD.GetMongoDb()
	return mD
}
