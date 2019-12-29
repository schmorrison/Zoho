package zoho

import (
	"encoding/gob"
	"errors"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"os"
	"time"
)

// TokenLoaderSaver is an interface that can be implemented when using a system that does
// not allow disk persistence, or a different type of persistence is required.
// The use case that was in mind was AppEngine where datastore is the only persistence option.
type TokenLoaderSaver interface {
	SaveTokens(t AccessTokenResponse) error
	LoadTokens() (AccessTokenResponse, error)
}

// SaveTokens will check for a provided 'TokenManager' interface
// if one exists it will use its provided method
func (z Zoho) SaveTokens(t AccessTokenResponse) error {
	if z.tokenManager != nil {
		return z.tokenManager.SaveTokens(t)
	}

	// Save the token response as GOB to file
	file, err := os.OpenFile(z.tokensFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Failed to open file '%s': %s", z.tokensFile, err)
	}
	enc := gob.NewEncoder(file)

	v := TokensWrapper{
		Tokens: z.oauth.token,
	}
	v.SetExpiry()

	err = enc.Encode(v)
	if err != nil {
		return fmt.Errorf("Failed to encode tokens to file '%s': %s", z.tokensFile, err)
	}

	return nil
}

// LoadTokens will check for a provided 'TokenManager' interface
// if one exists it will use its provided method
func (z Zoho) LoadTokens() (AccessTokenResponse, error) {
	if z.tokenManager != nil {
		return z.tokenManager.LoadTokens()
	}

	// Load the GOB and decode to AccessToken
	file, err := os.OpenFile(z.tokensFile, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("Failed to open file '%s': %s", z.tokensFile, err)
	}
	dec := gob.NewDecoder(file)

	var v TokensWrapper
	err = dec.Decode(&v)
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("Failed to decode tokens from file '%s': %s", z.tokensFile, err)
	}

	if v.CheckExpiry() {
		return v.Tokens, ErrTokenExpired
	}

	return v.Tokens, nil
}

// ErrTokenExpired should be returned when the token is expired but still exists in persistence
var ErrTokenExpired = errors.New("zoho: oAuth2 token already expired")

// ErrTokenInvalidCode is turned when the autorization code in a request is invalid
var ErrTokenInvalidCode = errors.New("zoho: authorization-code is invalid ")

// TokensWrapper should be used to provide the time.Time corresponding to the expiry of an access token
type TokensWrapper struct {
	Tokens  AccessTokenResponse
	Expires time.Time
}

// SetExpiry sets the TokenWrappers expiry time to now + seconds until expiry
func (t *TokensWrapper) SetExpiry() {
	t.Expires = time.Now().Add(time.Duration(t.Tokens.ExpiresIn) * time.Second)
}

// CheckExpiry if the token expired before this instant
func (t *TokensWrapper) CheckExpiry() bool {
	return t.Expires.Before(time.Now())
}

func (z *Zoho) checkForSavedTokens() error {
	t, err := z.LoadTokens()
	z.oauth.token = t

	if err != nil && err == ErrTokenExpired {
		return err
	}

	if (t != AccessTokenResponse{}) && err != ErrTokenExpired {
		return nil
	}
	return fmt.Errorf("No saved tokens")
}

// DatastoreManager is an example TokenManager that satisfies the TokenManager interface
// When instantiating, user must provide the *http.Request for the current app engine request
// and the token key where the tokens are to be saved to/loaded from.
type DatastoreManager struct {
	Request         *http.Request
	EntityNamespace string
	TokensKey       string
}

// LoadTokens will use datastore package to get tokens from the datastore under the entity namespace
// 'ZohoAccessTokens' unless a value is provided to the EntityNamespace field
func (d DatastoreManager) LoadTokens() (AccessTokenResponse, error) {
	t := TokensWrapper{}
	if d.Request == nil || d.TokensKey == "" {
		return AccessTokenResponse{}, fmt.Errorf("Must provide the *http.Request for the current request and a valid token key")
	}

	entity := "ZohoAccessTokens"
	if d.EntityNamespace != "" {
		entity = d.EntityNamespace
	}

	ctx := appengine.NewContext(d.Request)
	k := datastore.NewKey(ctx, entity, d.TokensKey, 0, nil)

	if err := datastore.Get(ctx, k, &t); err != nil {
		return AccessTokenResponse{}, fmt.Errorf("Failed to retrieve tokens from datastore: %s", err)
	}

	if t.CheckExpiry() {
		return AccessTokenResponse{}, ErrTokenExpired
	}

	return t.Tokens, nil
}

// SaveTokens will use datastore package to put tokens to the datastore under the entity namespace
// 'ZohoAccessTokens' unless a value is provided to the EntityNamespace field
func (d DatastoreManager) SaveTokens(t AccessTokenResponse) error {
	if d.Request == nil || d.TokensKey == "" {
		return fmt.Errorf("Must provide the *http.Request for the current request and a valid token key")
	}

	entity := "ZohoAccessTokens"
	if d.EntityNamespace != "" {
		entity = d.EntityNamespace
	}

	ctx := appengine.NewContext(d.Request)
	k := datastore.NewKey(ctx, entity, d.TokensKey, 0, nil)

	v := TokensWrapper{
		Tokens: t,
	}
	v.SetExpiry()

	if _, err := datastore.Put(ctx, k, v); err != nil {
		return fmt.Errorf("Failed to save tokens to datastore: %s", err)
	}

	return nil
}
