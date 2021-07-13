package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchServices(workspacesID zoho.Parameter, serviceID zoho.Parameter, staffID zoho.Parameter) (data ServiceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchServicesModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s", c.ZohoTLD, FetchServicesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ServiceResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}
	if workspacesID == ""{
		return ServiceResponse{}, fmt.Errorf("Failed to execute FetchServices due to non-availability of workspace_id")
	}
	endpoint.URLParameters["workspace_id"] = workspacesID
	if serviceID != "" {
		endpoint.URLParameters["service_id"] = serviceID
	}
	if staffID != "" {
		endpoint.URLParameters["staff_id"] = staffID
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
	Response struct {
		ReturnValue struct {
			Data []struct {
				Duration string `json:"duration"`
				Buffertime string `json:"buffertime"`
				Price int `json:"price"`
				Name string `json:"name"`
				Currency string `json:"currency"`
				Id string `json:"id"`
			} `json:"data"`
		} `json:"returnvalue"`
		Status string `json:"status"`
	} `json:"response"`
}
