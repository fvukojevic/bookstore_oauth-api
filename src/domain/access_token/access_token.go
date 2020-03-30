package access_token

import "time"

const (
	expirationTime = 24 // 24 h access token
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(token.Expires, 0)

	return now.After(expirationTime)
}
