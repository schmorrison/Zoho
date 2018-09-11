package zoho

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Endpoint struct {
	Method        HttpMethod
	URL           string
	Name          string
	ResponseData  interface{}
	URLParameters map[string]Parameter
	RequestBody   interface{}
}

type Parameter string

func (z *Zoho) HttpRequest(endpoint *Endpoint) (err error) {
	endpointURL := endpoint.URL

	q := url.Values{}
	for k, v := range endpoint.URLParameters {
		if v != "" {
			q.Set(k, string(v))
		}
	}

	var reqBody io.Reader
	if endpoint.RequestBody != nil {
		b, err := json.Marshal(endpoint.RequestBody)
		if err != nil {
			return fmt.Errorf("Failed to create json from request body")
		}

		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(string(endpoint.Method), fmt.Sprintf("%s?%s", endpointURL, q.Encode()), reqBody)
	if err != nil {
		return fmt.Errorf("Failed to create a request for %s: %s", endpoint.Name, err)
	}

	req.Header.Add("Authorization", "Zoho-oauthtoken "+z.oauth.token.AccessToken)

	resp, err := z.client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to perform request for %s: %s", endpoint.Name, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Failed to read body of response for %s: got status %s: %s", endpoint.Name, checkStatus(resp), err)
	}

	err = json.Unmarshal(body, &endpoint.ResponseData)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal data from response for %s: got status %s: %s", endpoint.Name, checkStatus(resp), err)
	}

	return nil
}

type HttpStatusCode int

var HttpStatusCodes = map[HttpStatusCode]string{
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

func checkStatus(r *http.Response) string {
	if v, ok := HttpStatusCodes[HttpStatusCode(r.StatusCode)]; ok {
		return v
	}
	return ""
}

type HttpHeader string

const (
	rateLimit          HttpHeader = "X-RATELIMIT-LIMIT"
	rateLimitRemaining HttpHeader = "X-RATELIMIT-REMAINING"
	rateLimitReset     HttpHeader = "X-RATELIMIT-RESET"
)

func checkHeaders(r http.Response, header HttpHeader) string {
	value := r.Header.Get(string(header))

	if value != "" {
		return value
	}
	return ""
}

type HttpMethod string

const (
	HTTPGet    HttpMethod = "GET"
	HTTPPost   HttpMethod = "POST"
	HTTPPut    HttpMethod = "PUT"
	HTTPDelete HttpMethod = "DELETE"
)
