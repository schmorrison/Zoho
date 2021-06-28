package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchResources(request interface{}, params map[string]zoho.Parameter) (data ResourceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchResourceModule,
		URL:          fmt.Sprintf(BookingsAPIEndpoint+"%s", FetchResourceModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ResourceResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}
	if len(params) != 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ResourceResponse{}, fmt.Errorf("Failed to retrieve resources: %s", err)
	}

	if v,ok := endpoint.ResponseData.(*ResourceResponse); ok {
		return *v, nil
	}
	return ResourceResponse{}, fmt.Errorf("Data retrieved was not 'Resource Response'")
}

type ResourceResponse struct {
	response struct {
		returnValue struct {
			data []struct {
				name string `json:"name"`
				id string `json:"id"`
			} `json:"data"`
		} `json:"returnvalue"`
		status string `json:"status"`
	} `json:"response"`
}
