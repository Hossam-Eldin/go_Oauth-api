package db

import (
	"context"
	"log"
	"github.com/Hossam-Eldin/go_Oauth-api/src/client/mongodb"
	"github.com/Hossam-Eldin/go_Oauth-api/src/domain/accesstoken"
	"github.com/Hossam-Eldin/go_Oauth-api/src/utils/errors"
	"go.mongodb.org/mongo-driver/bson"

)

//New : to return the methods in struct
func New() Repository {
	return &dbRepository{}
}

//Repository : interface
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

//GetByID : will tell the database to get accesstoken by this id
func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	log.Println("the coming id is ", id)

	
	c := mongodb.GetClient()
	collection := c.Database("golang").Collection("Oauth")

	var at accesstoken.AccessToken
	result := collection.FindOne(context.TODO(),bson.M{"accesstoken" : id})
	err := result.Decode(&at)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &at ,nil
}
