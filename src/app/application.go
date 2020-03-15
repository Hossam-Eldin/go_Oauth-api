package app

import (
	"context"
	"fmt"
	"log"

	"github.com/Hossam-Eldin/go_Oauth-api/src/client/mongodb"
	"github.com/Hossam-Eldin/go_Oauth-api/src/domain/accesstoken"
	"github.com/Hossam-Eldin/go_Oauth-api/src/http"
	"github.com/Hossam-Eldin/go_Oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

type At struct {
	AccessToken string
	UserID      int64
	ClientID    int64
	Expires     int64
}

var (
	router = gin.Default()
)

//StartApplication : start the app with this func
func StartApplication() {

	// Check the connection
	client := mongodb.GetClient()

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("golang").Collection("Oauth")

	ash := At{"Ash", 10, 25, 30}

	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	atHandler := http.NewHandler(accesstoken.NewService(db.New()))
	router.GET("/oauth/at/:access_token_id", atHandler.GetByID)
	router.Run(":3000")
}
