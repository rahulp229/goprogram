package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	DbName             string
	DbHost             string
	DbPort             string
	DbUsername         string
	DbPassword         string
	DbConnectionString string
}

type DBConnection struct {
	Client *mongo.Client
	DBName string
}

func ConnectDb() *DBConnection {

	db := DBConfig{DbName: "microservice_db", DbHost: "localhost", DbPort: "2717", DbUsername: "root", DbPassword: "password"}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db.DbConnectionString = "mongodb://" + db.DbHost + ":" + db.DbPort
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DbConnectionString))
	if err != nil {
		panic("db connection failed " + err.Error())
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalln("mongo db not reachable " + err.Error())
	}
	log.Println("database connected")
	dbConnection := DBConnection{Client: client, DBName: db.DbName}

	return &dbConnection
}

func CloseDb(c *mongo.Client) {
	if err := c.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
