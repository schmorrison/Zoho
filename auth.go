package zoho

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//GenerateAuthToken takes an email, password, and scope and retrieves an
// authentication token from zoho for the account and scope pair
func (z *Zoho) GenerateAuthToken(email, password, scope string) error {
	if email == "" || password == "" || scope == "" {
		return fmt.Errorf("Must set email, password, and scope when requesting an authtoken, got: Email='%s' Password='%s' Scope='%s'", email, password, scope)
	}
	resource := "https://accounts.zoho.com/apiauthtoken/nb/create"
	zr := z.NewRequest(resource, "GET")
	zr.Add("EMAIL_ID", email)
	zr.Add("PASSWORD", password)
	zr.Add("SCOPE", scope)

	resp, err := z.Request(zr)
	if err != nil {
		return fmt.Errorf("Error requesting authtoken: %s", err.Error())
	}

	z.authtoken, err = extractAuthToken(resp)
	if err != nil {
		return fmt.Errorf("Error extracting token: %s", err.Error())
	}
	z.user = email
	z.password = password
	z.scope = scope
	return nil
}

// extactAuthtoken will use regex to extract the token from the response provided by zoho
// this regex might be subject to change
func extractAuthToken(r *http.Response) (string, error) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading body from response: %s", err.Error())
	}

	token := regexp.MustCompile(`AUTHTOKEN=([a-zA-Z0-9]+)?\n`).FindStringSubmatch(string(b))
	if len(token) == 0 {
		return "", fmt.Errorf("There was no authtoken in the response")
	}
	return token[0], nil
}

// SetAuthToken will use a provided token and scope as for HTTP Request authentication
func (z *Zoho) SetAuthToken(token, scope string) error {
	if token == "" {
		return fmt.Errorf("Error: Must provide authtoken, provided: '%s'", token)
	}
	if scope == "" {
		return fmt.Errorf("Error: Must provide scope, provided '%s'", scope)
	}
	z.authtoken = token
	z.scope = scope
	return nil
}

// GetAuthToken will return the token that is set on the receiver Zoho object
func (z *Zoho) GetAuthToken() string {
	if t := z.authtoken; t != "" {
		return t
	}
	return ""
}

// SetOrganizationID will set the id as provided for use in HTTP request authentication
func (z *Zoho) SetOrganizationID(id string) error {
	if id == "" {
		return fmt.Errorf("Error: Must provide id, provided: '%s'", id)
	}
	z.orgID = id
	return nil
}
