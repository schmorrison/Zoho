package zoho

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (z *Zoho) RefreshTokenRequest() (err error) {
	q := url.Values{}
	q.Set("client_id", z.oauth.clientID)
	q.Set("client_secret", z.oauth.clientSecret)
	q.Set("refresh_token", z.oauth.token.RefreshToken)
	q.Set("grant_type", "refresh_token")

	tokenURL := fmt.Sprintf("%s%s?%s", oauthBaseURL, oauthGenerateTokenRequestSlug, q.Encode())
	resp, err := z.client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return fmt.Errorf("Failed while requesting refresh token: %s", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Failed to close request body: %s\n", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read request body on request to %s%s: %s", oauthBaseURL, oauthGenerateTokenRequestSlug, err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Got non-200 status code from request to refresh token: %s\n%s", resp.Status, string(body))
	}

	tokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal access token response from request to refresh token: %s", err)
	}

	z.oauth.token.AccessToken = tokenResponse.AccessToken
	z.oauth.token.APIDomain = tokenResponse.APIDomain
	z.oauth.token.ExpiresIn = tokenResponse.ExpiresIn
	z.oauth.token.ExpiresInSeconds = tokenResponse.ExpiresInSeconds
	z.oauth.token.TokenType = tokenResponse.TokenType

	return nil
}

func (z *Zoho) GenerateTokenRequest(clientID, clientSecret, code, redirectURI string) (err error) {
	q := url.Values{}
	q.Set("client_id", clientID)
	q.Set("client_secret", clientSecret)
	q.Set("code", code)
	q.Set("redirect_uri", redirectURI)
	q.Set("grant_type", "authorization_code")

	tokenURL := fmt.Sprintf("%s%s?%s", oauthBaseURL, oauthGenerateTokenRequestSlug, q.Encode())
	resp, err := z.client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return fmt.Errorf("Failed while requesting generate token: %s", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Failed to close request body: %s\n", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read request body on request to %s%s: %s", oauthBaseURL, oauthGenerateTokenRequestSlug, err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Got non-200 status code from request to generate token: %s\n%s", resp.Status, string(body))
	}

	tokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal access token response from request to generate token: %s", err)
	}

	z.oauth.clientID = clientID
	z.oauth.clientSecret = clientSecret
	z.oauth.redirectURI = redirectURI
	z.oauth.token = tokenResponse

	return nil
}

func (z *Zoho) AuthorizationCodeRequest(clientID, clientSecret string, scopes []ScopeString, redirectURI string) (err error) {
	scopeStr := ""
	for i, a := range scopes {
		scopeStr += string(a)
		if i < len(scopes)-1 {
			scopeStr += ","
		}
	}

	z.oauth.scopes = scopes

	q := url.Values{}
	q.Set("scope", scopeStr)
	q.Set("client_id", clientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("response_type", "code")
	q.Set("access_type", "offline")

	srvChan := make(chan int)
	codeChan := make(chan string)

	localRedirect := strings.Contains(redirectURI, "localhost")
	if localRedirect {
		// start a localhost server that will handle the redirect url
		u, err := url.Parse(redirectURI)
		if err != nil {
			return fmt.Errorf("Failed to parse redirect URI: %s", err)
		}
		_, port, err := net.SplitHostPort(u.Host)
		if err != nil {
			return fmt.Errorf("Failed to split redirect URI into host and port segments: %s", err)
		}
		srv := &http.Server{Addr: ":" + port}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("Code retrieved, you can close this window to continue"))

			codeChan <- r.URL.Query().Get("code")
			time.Sleep(500 * time.Millisecond)
			ctx := context.Background()
			if err := srv.Shutdown(ctx); err != nil {
				fmt.Printf("Error while shutting down local server: %s\n", err)
			}
		})

		go func() {
			srvChan <- 1
			err := srv.ListenAndServe()
			if err != nil {
				fmt.Printf("Error while serving locally: %s\n", err)
			}
		}()

		<-srvChan
	}

	authURL := fmt.Sprintf("%s%s?%s", oauthBaseURL, oauthAuthorizationRequestSlug, q.Encode())
	fmt.Printf("Go to the following authentication URL to begin oAuth2 flow:\n %s\n\n", authURL)

	code := ""

	if localRedirect {
		// wait for code to be returned by the server
		code = <-codeChan
	} else {
		fmt.Println("Paste code and press enter:\n")
		_, err := fmt.Scan(&code)
		if err != nil {
			return fmt.Errorf("Failed to read code from input: %s", err)
		}
	}

	if code == "" {
		return fmt.Errorf("No code was recieved from oAuth2 flow")
	}

	err = z.GenerateTokenRequest(clientID, clientSecret, code, redirectURI)
	if err != nil {
		return fmt.Errorf("Failed to retrieve oAuth2 token: %s", err)
	}

	return nil
}

type AccessTokenResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	ExpiresInSeconds int    `json:"expires_in_sec,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	APIDomain        string `json:"api_domain,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	Error            string `json:"error,omitempty"`
}

const (
	oauthBaseURL                  = "https://accounts.zoho.com/oauth/v2/"
	oauthAuthorizationRequestSlug = "auth"
	oauthGenerateTokenRequestSlug = "token"
	oauthRevokeTokenRequestSlug   = "revoke"
)

type ScopeString string

func BuildScope(service Service, scope Scope, method Method, operation Operation) ScopeString {
	built := fmt.Sprintf("%s.%s.%s", service, scope, method)
	if operation != "" {
		built += "." + string(operation)
	}
	return ScopeString(built)
}

type Service string

const (
	Crm Service = "ZohoCRM"
)

type Scope string

const (
	UsersScope    Scope = "users"
	OrgScope      Scope = "org"
	SettingsScope Scope = "settings"
	ModulesScope  Scope = "modules"
)

type Method string
type SettingsMethod = Method
type ModulesMethod = Method

const (
	AllMethod Method = "ALL"

	Territories   SettingsMethod = "territories"
	CustomViews   SettingsMethod = "custom_views"
	RelatedLists  SettingsMethod = "related_lists"
	Modules       SettingsMethod = "modules"
	TabGroups     SettingsMethod = "tab_groups"
	Fields        SettingsMethod = "fields"
	Layouts       SettingsMethod = "layouts"
	Macros        SettingsMethod = "macros"
	CustomLinks   SettingsMethod = "custom_links"
	CustomButtons SettingsMethod = "custom_buttons"
	Roles         SettingsMethod = "roles"
	Profiles      SettingsMethod = "profiles"

	Approvals      ModulesMethod = "approvals"
	Leads          ModulesMethod = "leads"
	Accounts       ModulesMethod = "accounts"
	Contacts       ModulesMethod = "contacts"
	Deals          ModulesMethod = "deals"
	Campaigns      ModulesMethod = "campaigns"
	Tasks          ModulesMethod = "tasks"
	Cases          ModulesMethod = "cases"
	Events         ModulesMethod = "events"
	Calls          ModulesMethod = "calls"
	Solutions      ModulesMethod = "solutions"
	Products       ModulesMethod = "products"
	Vendors        ModulesMethod = "vendors"
	PriceBooks     ModulesMethod = "pricebooks"
	Quotes         ModulesMethod = "quotes"
	SalesOrders    ModulesMethod = "salesorders"
	PurchaseOrders ModulesMethod = "purchaseorders"
	Invoices       ModulesMethod = "invoices"
	Custom         ModulesMethod = "custom"
	Dashboards     ModulesMethod = "dashboards"
	Notes          ModulesMethod = "notes"
	Activities     ModulesMethod = "activities"
	Search         ModulesMethod = "search"
)

type Operation string

const (
	NoOp   Operation = ""
	All    Operation = "ALL"
	Read   Operation = "READ"
	Create Operation = "CREATE"
	Update Operation = "UPDATE"
	Delete Operation = "DELETE"
)
