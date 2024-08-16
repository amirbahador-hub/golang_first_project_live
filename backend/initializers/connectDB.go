package initializers

import (
	"fmt"
	"log"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *gorm.DB
var MONGO *mongo.Client
var ProductCollection *mongo.Collection

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	MONGO, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI(config.MongoDbUri))

	ProductCollection = MONGO.Database("db").Collection("product")
	fmt.Println(MONGO)
	fmt.Println(config.MongoDbUri)
	// defer func() {
	// 	if err := MONGO.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err != nil {
		panic(err)
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}