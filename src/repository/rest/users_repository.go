package rest

import (
	"encoding/json"
	"github.com/fvukojevic/bookstore_oauth-api/src/domain/users"
	"github.com/fvukojevic/bookstore_util-go/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	UsersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(email string, password string) (*users.User, *errors.RestErr)
}

type restUserRepository struct {
}

func NewRepository() RestUserRepository {
	return &restUserRepository{}
}

func (rup restUserRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := UsersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		apiErr, err := errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}

		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user response")
	}

	return &user, nil
}
