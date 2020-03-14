package db

import (
	"github.com/Hossam-Eldin/go_Oauth-api/src/domain/accesstoken"
	"github.com/Hossam-Eldin/go_Oauth-api/src/utils/errors"
)

//New : to return the methods in struct 
func New() Repository{
	return &dbRepository{}
}

//Repository : interface 
type Repository interface {
		GetByID(string) (*accesstoken.AccessToken ,*errors.RestErr)
}


type dbRepository struct{

} 

//GetByID : will tell the database to get accesstoken by this id
func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr){
	return nil , errors.NewInternalServerError("database connection not implemented")
}

