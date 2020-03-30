package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	token := GetNewAccessToken()
	assert.False(t, token.IsExpired(), "Brand new access token should not be expired")
	assert.EqualValues(t, "", token.AccessToken, "new access token should not have access token field")
	assert.True(t, token.UserId == 0, "new access token should not have associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	token := AccessToken{}
	assert.True(t, token.IsExpired(), "empty access token should be expired by default")
	token.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, token.IsExpired(), "access token that has active time of 3 hours of now should not be expired")
}
