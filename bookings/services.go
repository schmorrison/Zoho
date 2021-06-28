package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchServices(request interface{}, params map[string]zoho.Parameter) (data ServiceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchServicesModule,
		URL:          fmt.Sprintf(BookingsAPIEndpoint+"%s", FetchServicesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ServiceResponse{},
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
		return ServiceResponse{}, fmt.Errorf("Failed to retrieve services: %s", err)
	}

	if v,ok := endpoint.ResponseData.(*ServiceResponse); ok {
		return *v, nil
	}
	return ServiceResponse{}, fmt.Errorf("Data retrieved was not 'Service Response'")
}

type ServiceResponse struct {
	response struct {
		returnValue struct {
			data []struct {
				duration string `json:"duration"`
				buffertime string `json:"buffertime"`
				price string `json:"price"`
				name string `json:"name"`
				currency string `json:"currency"`
				id string `json:"id"`
			} `json:"data"`
		} `json:"returnvalue"`
		status string `json:"status"`
	} `json:"response"`
}
