package access_token

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	reporsitory Repository
}

func NewService(repo Repository) Service {
	return &service{
		reporsitory: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.reporsitory.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
