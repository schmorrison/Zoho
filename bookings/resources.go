package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchResources(resourceID zoho.Parameter, serviceID zoho.Parameter) (data ResourceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchResourceModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s", c.ZohoTLD, FetchResourceModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ResourceResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}

	if resourceID != "" {
		endpoint.URLParameters["resource_id"] = resourceID
	}
	if serviceID != "" {
		endpoint.URLParameters["service_id"] = serviceID
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
	Response struct {
		ReturnValue struct {
			Data []struct {
				Name string `json:"name"`
				Id string `json:"id"`
			} `json:"data"`
		} `json:"returnvalue"`
		Status string `json:"status"`
	} `json:"response"`
}
