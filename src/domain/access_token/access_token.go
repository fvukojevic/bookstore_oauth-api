package access_token

import (
	"fmt"
	"github.com/fvukojevic/bookstore_oauth-api/src/utils/crypto"
	"github.com/fvukojevic/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24 // 24 h access token
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	//Used for password grant type
	Username string `json:"Username"`
	Password string `json:"password"`

	//Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (token *AccessTokenRequest) Validate() *errors.RestErr {
	//TODO: Validate parameters for each grant_type
	switch token.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (token AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}

	if token.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}

	if token.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}

	if token.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(token.Expires, 0)

	return now.After(expirationTime)
}

func (token *AccessToken) Generate() {
	token.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", token.UserId, token.Expires))
}
