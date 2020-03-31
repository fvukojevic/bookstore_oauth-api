package db

import (
	"github.com/fvukojevic/bookstore_oauth-api/src/clients/cassandra"
	"github.com/fvukojevic/bookstore_oauth-api/src/domain/access_token"
	"github.com/fvukojevic/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	// TODO: implement access token from cassandra db

	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("database connection not implementet yet")
}
