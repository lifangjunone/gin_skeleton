package tests

import (
	"context"
	"fmt"
	"gin_skeleton/SkeletonExample/form"
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func getMongoClient() *mongo.Client {
	confObj := conf.InitConfig("/home/lifangjun/Desktop/go_study/gin_skeleton/conf/env/dev.toml")
	mongoObj := confObj.Mongo
	MongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		mongoObj.UserName, mongoObj.Password, mongoObj.Host, mongoObj.Port,
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

func TestMongo(t *testing.T) {
	fmt.Println("mongo test")
	mCli := getMongoClient()
	coll := mCli.Database("go_skeleton").Collection("test")
	var user = models.User{
		Form: &form.UserForm{},
	}
	user.Form.Username = "root"
	user.Form.Password = "123456"
	user.Id = 1
	user.Hobby = "玩"
	user.Sex = "男"
	fmt.Printf("%#v", user)
	doc, _ := coll.InsertOne(context.Background(), user)
	fmt.Println(doc)

}

func TestEs(t *testing.T) {
	fmt.Println("es test")
}
