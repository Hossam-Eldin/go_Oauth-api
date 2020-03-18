package db

import (
	"context"
	"fmt"
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
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

//GetByID : will tell the database to get accesstoken by this id
func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	log.Println("the coming id is ", id)

	c := mongodb.GetClient()
	collection := c.Database("golang").Collection("Oauth")

	var at accesstoken.AccessToken
	result := collection.FindOne(context.TODO(), bson.M{"accesstoken": id})
	err := result.Decode(&at)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &at, nil
}

//Create : accesst token creation
func (r *dbRepository) Create(at accesstoken.AccessToken) *errors.RestErr {

	c := mongodb.GetClient()
	collection := c.Database("golang").Collection("Oauth")

	insertResult, err := collection.InsertOne(context.TODO(), bson.M{
		"accesstoken": at.AccessToken,
		"userid":      at.UserID,
		"clientid":    at.ClientID,
		"expires":     at.Expires,
	})
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	err = c.Disconnect(context.TODO())
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	c := mongodb.GetClient()
	collection := c.Database("golang").Collection("Oauth")
	updateData := bson.M{"$set": bson.M{
		"accesstoken": at.AccessToken,
		"expires":     at.Expires,
	},
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"accesstoken": at.AccessToken}, updateData)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	} else {
		log.Println("updated", result)
	}
	return nil
}
