package db

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/clients/cassandra"
	"github.com/fvukojevic/bookstore_oauth-api/src/domain/access_token"
	"github.com/fvukojevic/bookstore_util-go/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateAccessToken = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

type DbRepository interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var token access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&token.AccessToken,
		&token.UserId,
		&token.ClientId,
		&token.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access found for given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &token, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateAccessToken,
		token.Expires,
		token.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
