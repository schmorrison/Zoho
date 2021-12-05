package zoho

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOAuth2(t *testing.T) {
	t.Log("Testing OAuth2 flow")

	var (
		testClientID     = "test-client-id"
		testClientSecret = "test-client-secret"
		testCode         = "test-code"
		testRedirectURI  = "http://localhost:9646"
	)

	z := New()
	z.SetTokensFile(testTokensFile)
	z.LoadAccessAndRefreshToken()

	config := NewEndpointTestGroup(t).Fail(TestFailWarn).URI([]*TestZohoEndpoint{
		NewTestEndpoint(fmt.Sprintf("/oauth/v2/%s", oauthGenerateTokenRequestSlug)).
			Request(nil, &NullFormatter{}).Response(testTokenWrapper.Token).SetOnRequest(
			func(t *testing.T, r *http.Request, e error) error {
				grant := r.FormValue("grant_type")
				switch grant {
				case "authorization_code", "refresh_token":
					// test query params
					assert.Equalf(t, testClientID, z.oauth.clientID, "test http request failed: clientId '%s' == '%s'", testClientID, z.oauth.clientID)
					assert.Equalf(t, testClientSecret, z.oauth.clientSecret, "test http request failed: clientSecret '%s' == '%s'", testClientSecret, z.oauth.clientSecret)
					assert.Equalf(t, testRedirectURI, z.oauth.redirectURI, "test http request failed: redirectURI '%s' == '%s'", testRedirectURI, z.oauth.redirectURI)
					t.Logf("test http request passed for 'grant_type'= %s", grant)
				default:
					return fmt.Errorf("invalid grant type: %s", grant)
				}
				return nil
			}),
	})

	z.TestHTTPClient(config)
	config.Exec(func(t *testing.T) {
		var wg sync.WaitGroup
		t.Run("oauth-generate-request", func(t *testing.T) {
			wg.Add(1)
			err := z.GenerateTokenRequest(testClientID, testClientSecret, testCode, testRedirectURI)
			if err != nil {
				t.Errorf("GenerateTokenRequest failed: %s", err)
			}
			wg.Done()
		})
		t.Run("oauth-refresh-request", func(t *testing.T) {
			wg.Add(1)
			err := z.RefreshTokenRequest()
			if err != nil {
				t.Errorf("RefreshTokenRequest failed: %s", err)
			}
			wg.Done()
		})

		t.Run("oauth-authorization-code", func(t *testing.T) {
			wg.Add(1)
			z.clean()
			assert.NoFileExists(t, testTokensFile, "test http request failed: tokens file '%s' does exist", testTokensFile)
			go func() {
				// wait for local redirect server to start
				time.Sleep(50 * time.Millisecond)

				// send the authorization code to the local redirect server
				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?code=%s", testRedirectURI, testCode), nil)
				if err != nil {
					t.Errorf("build http request to localhost redirect URI failed: %s", err)
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					t.Errorf("http request to localhost redirect URI failed: %s", err)
				}

				// read the response body from the local redirect server
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("read http response body failed: %s", err)
				}

				// check the response body
				assert.Equal(t, fmt.Sprintf("Code [%s] retrieved, you can close this window to continue", testCode), string(respBody), "test http request failed: code '%s' == '%s'", testCode, string(respBody))

				t.Logf("code successfully sent to local redirect server")
				wg.Done()

			}()
			wg.Add(1)
			// start the oauth flow
			err := z.AuthorizationCodeRequest(testClientID, testClientSecret, []ScopeString{
				BuildScope(Crm, FullAccessScope, AllMethod, All),
			}, testRedirectURI)
			if err != nil {
				t.Errorf("AuthorizationCode failed: %s", err)
			}
			assert.Equalf(t, testClientID, z.oauth.clientID, "test http request failed: clientId '%s' == '%s'", testClientID, z.oauth.clientID)
			assert.Equalf(t, testClientSecret, z.oauth.clientSecret, "test http request failed: clientSecret '%s' == '%s'", testClientSecret, z.oauth.clientSecret)
			assert.Equalf(t, testRedirectURI, z.oauth.redirectURI, "test http request failed: redirectURI '%s' == '%s'", testRedirectURI, z.oauth.redirectURI)
			assert.FileExists(t, testTokensFile, "test http request failed: tokens file '%s' does not exist", testTokensFile)
			t.Logf("AuthorizationCode completed, and internal call to GenerateToken completed")
			wg.Done()
		})

		wg.Wait()
		config.complete <- true
	})
}
