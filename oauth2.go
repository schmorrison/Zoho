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

func (z *Zoho) SetRefreshToken(refreshToken string) {
	z.oauth.token.RefreshToken = refreshToken
}

func (z *Zoho) SetClientID(clientID string) {
	z.oauth.clientID = clientID
}

func (z *Zoho) SetClientSecret(clientSecret string) {
	z.oauth.clientSecret = clientSecret
}

// RefreshTokenRequest is used to refresh the oAuth2 access token
func (z *Zoho) RefreshTokenRequest() (err error) {
	q := url.Values{}
	q.Set("client_id", z.oauth.clientID)
	q.Set("client_secret", z.oauth.clientSecret)
	q.Set("refresh_token", z.oauth.token.RefreshToken)
	q.Set("grant_type", "refresh_token")

	tokenURL := fmt.Sprintf("%s%s?%s", z.oauth.baseURL, oauthGenerateTokenRequestSlug, q.Encode())
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
		return fmt.Errorf("Failed to read request body on request to %s%s: %s", z.oauth.baseURL, oauthGenerateTokenRequestSlug, err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Got non-200 status code from request to refresh token: %s\n%s", resp.Status, string(body))
	}

	tokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal access token response from request to refresh token: %s", err)
	}
	//If the tokenResponse is not valid it should not update local tokens
	if tokenResponse.Error == "invalid_code" {
		return ErrTokenInvalidCode
	}

	z.oauth.token.AccessToken = tokenResponse.AccessToken
	z.oauth.token.APIDomain = tokenResponse.APIDomain
	z.oauth.token.ExpiresIn = tokenResponse.ExpiresIn
	z.oauth.token.TokenType = tokenResponse.TokenType

	err = z.SaveTokens(z.oauth.token)
	if err != nil {
		return fmt.Errorf("Failed to save access tokens: %s", err)
	}

	return nil
}

// GenerateTokenRequest will get the Access token and Refresh token and hold them in the Zoho struct. This function can be used rather than
// AuthorizationCodeRequest is you do not want to click on a link and redirect to a consent screen. Instead you can go to, https://accounts.zoho.com/developerconsole
// and click the kebab icon beside your clientID, and click 'Self-Client'; then you can define you scopes and an expiry, then provide the generated authorization code
// to this function which will generate your access token and refresh tokens.
func (z *Zoho) GenerateTokenRequest(clientID, clientSecret, code, redirectURI string) (err error) {

	z.oauth.clientID = clientID
	z.oauth.clientSecret = clientSecret
	z.oauth.redirectURI = redirectURI

	err = z.CheckForSavedTokens()
	if err == ErrTokenExpired {
		return z.RefreshTokenRequest()
	}

	q := url.Values{}
	q.Set("client_id", clientID)
	q.Set("client_secret", clientSecret)
	q.Set("code", code)
	q.Set("redirect_uri", redirectURI)
	q.Set("grant_type", "authorization_code")

	tokenURL := fmt.Sprintf("%s%s?%s", z.oauth.baseURL, oauthGenerateTokenRequestSlug, q.Encode())
	fmt.Printf(tokenURL)
	resp, err := z.client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return fmt.Errorf("Failed while requesting generate token: %s", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Failed to close request body: %s\n", err)
		}
	}()
	fmt.Printf(tokenURL)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read request body on request to %s%s: %s", z.oauth.baseURL, oauthGenerateTokenRequestSlug, err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Got non-200 status code from request to generate token: %s\n%s", resp.Status, string(body))
	}

	tokenResponse := AccessTokenResponse{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal access token response from request to generate token: %s", err)
	}

	//If the tokenResponse is not valid it should not update local tokens
	if tokenResponse.Error == "invalid_code" {
		return ErrTokenInvalidCode
	}

	z.oauth.clientID = clientID
	z.oauth.clientSecret = clientSecret
	z.oauth.redirectURI = redirectURI
	z.oauth.token = tokenResponse

	err = z.SaveTokens(z.oauth.token)
	if err != nil {
		return fmt.Errorf("Failed to save access tokens: %s", err)
	}

	return nil
}

// AuthorizationCodeRequest will request an authorization code from Zoho. This authorization code is then used to generate access and refresh tokens.
// This function will print a link that needs to be pasted into a browser to continue the oAuth2 flow. Then it will redirect to the redirectURL, it
// must be the same as the redirect URL that was provided to Zoho when generating your client ID and client secret. If the redirect URL was a localhost
// domain, the function will start a server that will get the code from the URL when the browser redirects.
// If the domain is not a localhost, you will be prompted to paste the code from the URL back into the terminal window,
// eg. https://domain.com/redirect-url?code=xxxxxxxxxx
func (z *Zoho) AuthorizationCodeRequest(clientID, clientSecret string, scopes []ScopeString, redirectURI string) (err error) {
	// check for existing tokens
	err = z.CheckForSavedTokens()
	if err == nil {
		z.oauth.clientID = clientID
		z.oauth.clientSecret = clientSecret
		z.oauth.redirectURI = redirectURI
		z.oauth.scopes = scopes
		return nil
	}

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
	locChan := make(chan string)
	var srv *http.Server

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
		srv = &http.Server{Addr: ":" + port}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Code retrieved, you can close this window to continue"))

			codeChan <- r.URL.Query().Get("code")
			locChan <- r.URL.Query().Get("location")
		})

		go func() {
			srvChan <- 1
			err := srv.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				fmt.Printf("Error while serving locally: %s\n", err)
			}
		}()

		<-srvChan
	}

	authURL := fmt.Sprintf("%s%s?%s", z.oauth.baseURL, oauthAuthorizationRequestSlug, q.Encode())
	fmt.Printf("Go to the following authentication URL to begin oAuth2 flow:\n %s\n\n", authURL)

	code := ""
	location := ""
	if localRedirect {
		// wait for code to be returned by the server
		code = <-codeChan
		location = <-locChan
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer func() {
			cancel()
		}()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("Error while shutting down local server: %s\n", err)
		}
	} else {
		fmt.Printf("Paste code and press enter:\n")
		_, err := fmt.Scan(&code)
		if err != nil {
			return fmt.Errorf("Failed to read code from input: %s", err)
		}
	}

	if code == "" {
		return fmt.Errorf("No code was recieved from oAuth2 flow")
	}
	fmt.Printf(code)
	z.SetZohoTLD(location)
	err = z.GenerateTokenRequest(clientID, clientSecret, code, redirectURI)
	if err != nil {
		return fmt.Errorf("Failed to retrieve oAuth2 token: %s", err)
	}

	return nil
}

// AccessTokenResponse is the data returned when generating AccessTokens, or Refreshing the token
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	APIDomain    string `json:"api_domain,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	Error        string `json:"error,omitempty"`
}

const (
	oauthAuthorizationRequestSlug = "auth"
	oauthGenerateTokenRequestSlug = "token"
	oauthRevokeTokenRequestSlug   = "revoke"
)

// ScopeString is a type for defining scopes for oAuth2 flow
type ScopeString string

// BuildScope is used to generate a scope string for oAuth2 flow
func BuildScope(service Service, scope Scope, method Method, operation Operation) ScopeString {
	var built string = ""
	if method != "" {
		built += fmt.Sprintf("%s.%s.%s", service, scope, method)
	} else {
		built += fmt.Sprintf("%s.%s", service, scope)
	}
	if operation != "" {
		built += "." + string(operation)
	}
	return ScopeString(built)
}

// Service is a type for building scopes
type Service string

const (
	// Crm is the Service portion of the scope string
	Crm Service = "ZohoCRM"
	// Expense is the Service portion of the scope string
	Expense Service = "ZohoExpense"
	// Bookings is the Service portion of the scope string
	Bookings Service = "zohobookings"
)

// Scope is a type for building scopes
type Scope string

const (
	// UsersScope is a possible Scope portion of the scope string
	UsersScope Scope = "users"
	// OrgScope is a possible Scope portion of the scope string
	OrgScope Scope = "org"
	// SettingsScope is a possible Scope portion of the scope string
	SettingsScope Scope = "settings"
	// ModulesScope is a possible Scope portion of the scope string
	ModulesScope Scope = "modules"

	// Additional Scopes related to expense APIs

	// FullAccessScope is a possible Method portion of the scope string
	FullAccessScope Scope = "fullaccess"
	// ExpenseReportScope is a possible Method portion of the scope string
	ExpenseReportScope Scope = "expensereport"
	// ApprovalScope is a possible Method portion of the scope string
	ApprovalScope Scope = "approval"
	// ReimbursementScope is a possible Method portion of the scope string
	ReimbursementScope Scope = "reimbursement"
	// AdvanceScope is a possible Method portion of the scope string
	AdvanceScope Scope = "advance"
	// DataScope is a possible Method portion of the scope string
	DataScope Scope = "data"
)

// Method is a type for building scopes
type Method string

// SettingsMethod is a type for building scopes
type SettingsMethod = Method

// ModulesMethod is a type for building scopes
type ModulesMethod = Method

const (
	// AllMethod is a possible Method portion of the scope string
	AllMethod Method = "ALL"

	// Territories is a possible Method portion of the scope string
	Territories SettingsMethod = "territories"
	// CustomViews is a possible Method portion of the scope string
	CustomViews SettingsMethod = "custom_views"
	// RelatedLists is a possible Method portion of the scope string
	RelatedLists SettingsMethod = "related_lists"
	// Modules is a possible Method portion of the scope string
	Modules SettingsMethod = "modules"
	// TabGroups is a possible Method portion of the scope string
	TabGroups SettingsMethod = "tab_groups"
	// Fields is a possible Method portion of the scope string
	Fields SettingsMethod = "fields"
	// Layouts is a possible Method portion of the scope string
	Layouts SettingsMethod = "layouts"
	// Macros is a possible Method portion of the scope string
	Macros SettingsMethod = "macros"
	// CustomLinks is a possible Method portion of the scope string
	CustomLinks SettingsMethod = "custom_links"
	// CustomButtons is a possible Method portion of the scope string
	CustomButtons SettingsMethod = "custom_buttons"
	// Roles is a possible Method portion of the scope string
	Roles SettingsMethod = "roles"
	// Profiles is a possible Method portion of the scope string
	Profiles SettingsMethod = "profiles"

	// Approvals is a possible Method portion of the scope string
	Approvals ModulesMethod = "approvals"
	// Leads is a possible Method portion of the scope string
	Leads ModulesMethod = "leads"
	// Accounts is a possible Method portion of the scope string
	Accounts ModulesMethod = "accounts"
	// Contacts is a possible Method portion of the scope string
	Contacts ModulesMethod = "contacts"
	// Deals is a possible Method portion of the scope string
	Deals ModulesMethod = "deals"
	// Campaigns is a possible Method portion of the scope string
	Campaigns ModulesMethod = "campaigns"
	// Tasks is a possible Method portion of the scope string
	Tasks ModulesMethod = "tasks"
	// Cases is a possible Method portion of the scope string
	Cases ModulesMethod = "cases"
	// Events is a possible Method portion of the scope string
	Events ModulesMethod = "events"
	// Calls is a possible Method portion of the scope string
	Calls ModulesMethod = "calls"
	// Solutions is a possible Method portion of the scope string
	Solutions ModulesMethod = "solutions"
	// Products is a possible Method portion of the scope string
	Products ModulesMethod = "products"
	// Vendors is a possible Method portion of the scope string
	Vendors ModulesMethod = "vendors"
	// PriceBooks is a possible Method portion of the scope string
	PriceBooks ModulesMethod = "pricebooks"
	// Quotes is a possible Method portion of the scope string
	Quotes ModulesMethod = "quotes"
	// SalesOrders is a possible Method portion of the scope string
	SalesOrders ModulesMethod = "salesorders"
	// PurchaseOrders is a possible Method portion of the scope string
	PurchaseOrders ModulesMethod = "purchaseorders"
	// Invoices is a possible Method portion of the scope string
	Invoices ModulesMethod = "invoices"
	// Custom is a possible Method portion of the scope string
	Custom ModulesMethod = "custom"
	// Dashboards is a possible Method portion of the scope string
	Dashboards ModulesMethod = "dashboards"
	// Notes is a possible Method portion of the scope string
	Notes ModulesMethod = "notes"
	// Activities is a possible Method portion of the scope string
	Activities ModulesMethod = "activities"
	// Search is a possible Method portion of the scope string
	Search ModulesMethod = "search"
)

// Operation is a type for building scopes
type Operation string

const (
	// NoOp is a possible Operation portion of the scope string
	NoOp Operation = ""
	// All is a possible Operation portion of the scope string
	All Operation = "ALL"
	// Read is a possible Operation portion of the scope string
	Read Operation = "READ"
	// Create is a possible Operation portion of the scope string
	Create Operation = "CREATE"
	// Update is a possible Operation portion of the scope string
	Update Operation = "UPDATE"
	// Delete is a possible Operation portion of the scope string
	Delete Operation = "DELETE"
)
