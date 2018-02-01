package zoho

import (
	"fmt"
	"net/http"
	"net/url"
)

//NewRequest uses the resource and method to initialize a ZohoRequest for usage
func (z *Zoho) NewRequest(resource, method string) *Request {
	zr := &Request{
		Resource: resource,
		Method:   method,
		vals:     url.Values{},
	}
	return zr
}

//Request will perform the provided request with the account provided in the Zoho pointer
func (z *Zoho) Request(zr *Request) (*http.Response, error) {
	req, err := http.NewRequest(zr.Method, zr.URL(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error initializing http request: %s", err.Error())
	}
	resp, err := z.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error performing http request: %s", err.Error())
	}

	return resp, nil
}

//CustomClient replaces the preinitialized http Client with the provided Client
func (z *Zoho) CustomClient(c *http.Client) {
	z.client = c
}

func (zr *Request) URL() string {
	u, err := url.Parse(zr.Resource)
	if err != nil {
		fmt.Println("Got error parsing resource URL: ", err.Error())
		return ""
	}
	u.RawQuery = zr.vals.Encode()
	return u.String()
}

//Add will add the provided key/val pair as a URL value to the request
func (zr *Request) Add(key string, val interface{}) error {
	valStr := ""
	switch t := val.(type) {
	case bool:
		valStr = fmt.Sprintf("%t", val.(bool))
	case string:
		valStr = val.(string)
	case int:
		valStr = fmt.Sprintf("%d", val.(int))
	case int8:
		valStr = fmt.Sprintf("%d", val.(int8))
	case int16:
		valStr = fmt.Sprintf("%d", val.(int16))
	case int32:
		valStr = fmt.Sprintf("%d", val.(int32))
	case int64:
		valStr = fmt.Sprintf("%d", val.(int64))
	case float32:
		valStr = fmt.Sprintf("%f", val.(float32))
	case float64:
		valStr = fmt.Sprintf("%f", val.(float64))
	case complex64:
		valStr = fmt.Sprintf("%g", val.(complex64))
	case complex128:
		valStr = fmt.Sprintf("%g", val.(complex128))
	default:
		return fmt.Errorf("The data type '%T' is not supported", t)
	}
	zr.vals.Add(key, valStr)
	return nil
}

//AddMulti will add each item in the map to the requests url Values
func (zr *Request) AddMulti(m map[string]interface{}) {
	if len(m) > 0 {
		for k, v := range m {
			zr.Add(k, v)
		}
	}
}

//Request is the structure used prepare an http.Request
type Request struct {
	Resource string
	Method   string
	Map      map[string]interface{}
	vals     url.Values
	request  http.Request
}
