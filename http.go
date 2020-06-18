package zoho

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// Endpoint defines the data required to interact with most Zoho REST api endpoints
type Endpoint struct {
	Method        HTTPMethod
	URL           string
	Name          string
	ResponseData  interface{}
	RequestBody   interface{}
	URLParameters map[string]Parameter
	Headers       map[string]string
	JSONString    bool
}

// Parameter is used to provide URL Parameters to zoho endpoints
type Parameter string

// HTTPRequest is the function which actually performs the request to a Zoho endpoint as specified by the provided endpoint
func (z *Zoho) HTTPRequest(endpoint *Endpoint) (err error) {
	if reflect.TypeOf(endpoint.ResponseData).Kind() != reflect.Ptr {
		return fmt.Errorf("Failed, you must pass a pointer in the ResponseData field of endpoint")
	}

	var reqBody io.Reader
	if endpoint.RequestBody != nil {
		b, err := json.Marshal(endpoint.RequestBody)
		if err != nil {
			return fmt.Errorf("Failed to create json from request body")
		}
		if endpoint.JSONString {
			requestBodyString := string(b)
			reqBody = strings.NewReader("JSONString=" + requestBodyString)
		} else {
			reqBody = bytes.NewReader(b)
		}
	}

	// Load and renew access token if expired
	err = z.CheckForSavedTokens()
	if err == ErrTokenExpired {
		err := z.RefreshTokenRequest()
		if err != nil {
			return fmt.Errorf("Failed to refresh the access token: %s: %s", endpoint.Name, err)
		}
	}

	endpointURL := endpoint.URL
	q := url.Values{}
	for k, v := range endpoint.URLParameters {
		if v != "" {
			q.Set(k, string(v))
		}
	}

	req, err := http.NewRequest(string(endpoint.Method), fmt.Sprintf("%s?%s", endpointURL, q.Encode()), reqBody)
	if err != nil {
		return fmt.Errorf("Failed to create a request for %s: %s", endpoint.Name, err)
	}

	req.Header.Add("Authorization", "Zoho-oauthtoken "+z.oauth.token.AccessToken)

	// Add mandatory header for specific APIs
	if z.OrganizationID != "" {
		// Zoho doesn't use regular JSON content type
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		for k, v := range endpoint.Headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := z.client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to perform request for %s: %s", endpoint.Name, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read body of response for %s: got status %s: %s", endpoint.Name, resolveStatus(resp), err)
	}

	dataType := reflect.TypeOf(endpoint.ResponseData).Elem()
	data := reflect.New(dataType).Interface()

	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal data from response for %s: got status %s: %s", endpoint.Name, resolveStatus(resp), err)
	}

	endpoint.ResponseData = data

	return nil
}

// HTTPStatusCode is a type for resolving the returned HTTP Status Code Content
type HTTPStatusCode int

// HTTPStatusCodes is a map of possible HTTP Status Code and Messages
var HTTPStatusCodes = map[HTTPStatusCode]string{
	200: "The API request is successful.",
	201: "Request fulfilled for single record insertion.",
	202: "Request fulfilled for multiple records insertion.",
	204: "There is no content available for the request.",
	304: "The requested page has not been modified. In case \"If-Modified-Since\" header is used for GET APIs",
	400: "The request or the authentication considered is invalid.",
	401: "Invalid API key provided.",
	403: "No permission to do the operation.",
	404: "Invalid request.",
	405: "The specified method is not allowed.",
	413: "The server did not accept the request while uploading a file, since the limited file size has exceeded.",
	415: "The server did not accept the request while uploading a file, since the media/ file type is not supported.",
	429: "Number of API requests per minute/day has exceeded the limit.",
	500: "Generic error that is encountered due to an unexpected server error.",
}

func resolveStatus(r *http.Response) string {
	if v, ok := HTTPStatusCodes[HTTPStatusCode(r.StatusCode)]; ok {
		return v
	}
	return ""
}

// HTTPHeader is a type for defining possible HTTPHeaders that zoho request could return
type HTTPHeader string

const (
	rateLimit          HTTPHeader = "X-RATELIMIT-LIMIT"
	rateLimitRemaining HTTPHeader = "X-RATELIMIT-REMAINING"
	rateLimitReset     HTTPHeader = "X-RATELIMIT-RESET"
)

func checkHeaders(r http.Response, header HTTPHeader) string {
	value := r.Header.Get(string(header))

	if value != "" {
		return value
	}
	return ""
}

// HTTPMethod is a type for defining the possible HTTP request methods that can be used
type HTTPMethod string

const (
	// HTTPGet is the GET method for http requests
	HTTPGet HTTPMethod = "GET"
	// HTTPPost is the POST method for http requests
	HTTPPost HTTPMethod = "POST"
	// HTTPPut is the PUT method for http requests
	HTTPPut HTTPMethod = "PUT"
	// HTTPDelete is the DELETE method for http requests
	HTTPDelete HTTPMethod = "DELETE"
)
