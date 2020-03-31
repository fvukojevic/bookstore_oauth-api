package access_token

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(token AccessToken) *errors.RestErr
	UpdateExpirationTime(token AccessToken) *errors.RestErr
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
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.reporsitory.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(token AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.reporsitory.Create(token)
}

func (s *service) UpdateExpirationTime(token AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.reporsitory.UpdateExpirationTime(token)
}
