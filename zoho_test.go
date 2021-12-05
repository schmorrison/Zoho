package zoho

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testTokenWrapper = TokenWrapper{
		Token: AccessTokenResponse{
			AccessToken:  "access_token:12345",
			RefreshToken: "refresh_token:12345",
			ExpiresIn:    5000,
			APIDomain:    "api.domain.com",
			TokenType:    "cat",
			Error:        "error",
		},
		Expires: time.Now(),
	}
	testTokensFile = "./_testdata/tmp/tokensfile.gob"
)

func TestZoho(t *testing.T) {
	z := New()

	t.Run("set-tokens", func(t *testing.T) {
		z.SetTokensFile(testTokensFile)
		assert.Equalf(t, testTokensFile, z.tokensFile, "SetTokensFile failed: %s", z.tokensFile)
	})

	t.Run("save-tokens-file", func(t *testing.T) {
		err := z.SaveTokens(testTokenWrapper.Token)
		if err != nil {
			t.Errorf("save-tokens-file: SaveTokens failed: %s", err)
		}
	})

	t.Run("read-tokens-file", func(t *testing.T) {
		err := z.CheckForSavedTokens()
		if err != nil {
			t.Errorf("CheckForSavedTokens failed: %s", err)
		}

		assert.Equalf(t, testTokenWrapper.Token.AccessToken, z.oauth.token.AccessToken, "LoadAccessAndRefreshToken failed: '%s' == '%s'", testTokenWrapper.Token.AccessToken, z.oauth.token.AccessToken)
		assert.Equalf(t, testTokenWrapper.Token.RefreshToken, z.oauth.token.RefreshToken, "LoadAccessAndRefreshToken failed: '%s' == '%s'", testTokenWrapper.Token.RefreshToken, z.oauth.token.RefreshToken)
	})

	t.Run("set-token-manager", func(t *testing.T) {
		z.SetTokenManager(&TestTokenManager{})
		if err := z.SaveTokens(AccessTokenResponse{}); err == nil || err != ErrTokenManagerSuccess {
			assert.Equalf(t, ErrTokenManagerSuccess, err, "SetTokenManager -> SaveTokens failed: did not return 'success' error: %s", err)
		}

		if _, err := z.LoadAccessAndRefreshToken(); err == nil || err != ErrTokenManagerSuccess {
			assert.Equalf(t, ErrTokenManagerSuccess, err, "SetTokenManager -> Load... failed: did not return 'success' error: %s", err)
		}
	})

	t.Run("set-tld", func(t *testing.T) {
		tld := "zh"
		z.SetZohoTLD(tld)
		assert.Equalf(t, tld, z.ZohoTLD, "SetZohoTLD failed: %s", z.ZohoTLD)
		assert.Equalf(t, fmt.Sprintf("https://accounts.zoho.%s/oauth/v2/", tld), z.oauth.baseURL, "SetZohoTLD failed: %s", z.oauth.baseURL)
	})

	t.Run("set-organization-id", func(t *testing.T) {
		orgID := "11235813"
		z.SetOrganizationID(orgID)
		assert.Equalf(t, orgID, z.OrganizationID, "SetOrganizationID failed: %s", z.OrganizationID)
	})

	t.Run("set-refresh-token", func(t *testing.T) {
		refreshToken := "11235813"
		z.SetRefreshToken(refreshToken)
		assert.Equalf(t, refreshToken, z.oauth.token.RefreshToken, "SetRefreshToken failed: %s", z.oauth.token.RefreshToken)
	})

	t.Run("set-client-id", func(t *testing.T) {
		clientId := "11235813"
		z.SetClientID(clientId)
		assert.Equalf(t, clientId, z.oauth.clientID, "SetClientID failed: %s", z.oauth.clientID)
	})

	t.Run("set-client-secret", func(t *testing.T) {
		clientSecret := "11235813"
		z.SetClientSecret(clientSecret)
		assert.Equalf(t, clientSecret, z.oauth.clientSecret, "SetClientSecret failed: %s", z.oauth.clientSecret)
	})

}

var ErrTokenManagerSuccess = errors.New("success")

type TestTokenManager struct {
}

func (T *TestTokenManager) SaveTokens(t AccessTokenResponse) error {
	return ErrTokenManagerSuccess
}

func (T *TestTokenManager) LoadAccessAndRefreshToken() (AccessTokenResponse, error) {
	return AccessTokenResponse{}, ErrTokenManagerSuccess
}
