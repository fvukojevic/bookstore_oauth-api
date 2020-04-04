package services

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/domain/access_token"
	"github.com/fvukojevic/bookstore_oauth-api/src/repository/rest"
	"github.com/fvukojevic/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(tokenRequest access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type service struct {
	dbRepo       Repository
	restUserRepo rest.RestUserRepository
}

func NewService(dbRepo Repository, usersRepo rest.RestUserRepository) Service {
	return &service{
		dbRepo:       dbRepo,
		restUserRepo: usersRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(tokenRequest access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := tokenRequest.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both client_credentials and password grant_Types
	//Authenticate user against user api
	user, err := s.restUserRepo.LoginUser(tokenRequest.Username, tokenRequest.Password)
	if err != nil {
		return nil, err
	}

	token := access_token.GetNewAccessToken(user.Id)
	token.Generate()

	if err := s.dbRepo.Create(token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *service) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(token)
}
