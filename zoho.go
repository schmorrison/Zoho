package zoho

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// New initializes a Zoho structure
func New() *Zoho {
	z := Zoho{
		client: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
		ZohoTLD:    "com",
		tokensFile: "./.tokens.zoho",
		oauth: OAuth{
			baseURL: "https://accounts.zoho.com/oauth/v2/",
		},
	}

	return &z
}

// SetTokenManager can be used to provide a type which implements the TokenManager interface
// which will get/set AccessTokens/RefreshTokens using a persistence mechanism
func (z *Zoho) SetTokenManager(tm TokenLoaderSaver) {
	z.tokenManager = tm
}

// SetTokensFile can be used to set the file location of the token persistence location,
// by default tokens are stored in a file in the current directory called '.tokens.zoho'
func (z *Zoho) SetTokensFile(s string) {
	z.tokensFile = s
}

// SetZohoTLD can be used to set the TLD extension for API calls for example for Zoho in EU and China.
// by default this is set to "com", other options are "eu" and "ch"
func (z *Zoho) SetZohoTLD(s string) {
	z.ZohoTLD = s
	z.oauth.baseURL = fmt.Sprintf("https://accounts.zoho.%s/oauth/v2/", s)
}

// CustomHTTPClient can be used to provide a custom HTTP Client that replaces the once instantiated
// when executing New()
//
// A notable use case is AppEngine where a user must use the appengine/urlfetch packages provided http client
// when performing outbound http requests.
func (z *Zoho) CustomHTTPClient(c *http.Client) {
	z.client = c
}

// SetOrganizationID cab be used to add organization id in zoho struct
// which is needed for expense apis
func (z *Zoho) SetOrganizationID(orgID string) {
	z.organizationID = orgID
}

// Zoho is for accessing all APIs. It is used by subpackages to simplify passing authentication
// values between API subpackages.
type Zoho struct {
	oauth OAuth

	client         *http.Client
	tokenManager   TokenLoaderSaver
	tokensFile     string
	organizationID string

	ZohoTLD string
}

// OAuth is the OAuth part of the Zoho struct
type OAuth struct {
	scopes       []ScopeString
	clientID     string
	clientSecret string
	redirectURI  string
	token        AccessTokenResponse
	baseURL      string
}
