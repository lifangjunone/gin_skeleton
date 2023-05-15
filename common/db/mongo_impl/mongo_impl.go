package mongo_impl

import (
	"context"
	"fmt"
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/conf"
	"gin_skeleton/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *MongoDb) InsertOne(colName string, data interface{}) {
	coll := m.MDB.Collection(colName)
	coll.InsertOne(context.Background(), data)
}

func (m *MongoDb) QueryByUsername(colName string, username string) interface{} {
	coll := m.MDB.Collection(colName)
	filter := bson.D{{"username", username}}
	data := coll.FindOne(context.Background(), filter)
	return data
}

func (m *MongoDb) QueryById(colName string, id string, user *models.User) {
	coll := m.MDB.Collection(colName)
	id_, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": id_}
	fmt.Println(filter)
	data := coll.FindOne(context.Background(), filter)
	data.Decode(user)
}

func (m *MongoDb) QueryByField(colName string, jsonParam map[string]interface{}) (users *[]models.User) {
	coll := m.MDB.Collection(colName)
	var key string
	var value interface{}
	for k, v := range jsonParam {
		key = k
		value = v
	}
	users = new([]models.User)
	filter := bson.M{key: value}
	cur, _ := coll.Find(context.TODO(), filter)
	err := cur.All(context.TODO(), users)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return
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
