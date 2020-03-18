package accesstoken

import (
	"strings"

	"github.com/Hossam-Eldin/go_Oauth-api/src/utils/errors"
)

//Repository : method
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

//Service :interface for the service methods to test and call from outside
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

//NewService : to handle the intferace struct
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

//GetByID : this method will handle return the accessToken the
// logic point bettewn the repostiory and controller http
func (s *service) GetByID(accesstTokenID string) (*AccessToken, *errors.RestErr) {
	accesstTokenID = strings.TrimSpace(accesstTokenID)
	if len(accesstTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid data request need access token")
	}

	accesstoken, err := s.repository.GetByID(accesstTokenID)
	if err != nil {
		return nil, err
	}
	return accesstoken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.UpdateExpirationTime(at)
}
