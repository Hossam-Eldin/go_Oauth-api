package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongodb://<dbuser>:<dbpassword>@ds055762.mlab.com:55762/golang
//user  go_root
//password go1579531

//GetClient : to start connection with mongodb
func GetClient() *mongo.Client {
	ClientOptions := options.Client().ApplyURI("mongodb://go_root:go1579531@ds055762.mlab.com:55762/golang?retryWrites=false")
	client, err := mongo.NewClient(ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client

}
