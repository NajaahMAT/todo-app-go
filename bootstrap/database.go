package bootstrap

import (
	"context"
	"time"

	"todo-app-go/mongo"

	"log"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// mongodbURI := env.MongoURI

	// dbHost := env.DBHost
	// dbPort := env.DBPort
	// dbUser := env.DBUser
	// dbPass := env.DBPass

	// mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	mongodbURI := "mongodb+srv://Meezan:Meezan123@cluster0.tk8k6.mongodb.net/Meezan?retryWrites=true&w=majority"
	//mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", "Najaah", "Najaah123", "localhost", "27017")

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
