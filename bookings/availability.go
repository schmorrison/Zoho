package bookings

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchAvailability(
	serviceID zoho.Parameter,
	staffID zoho.Parameter,
	resourceID zoho.Parameter,
	date zoho.Parameter,
) (data AvailabilityResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: FetchServicesModule,
		URL: fmt.Sprintf(
			"https://www.zohoapis.%s/bookings/v1/json/%s",
			c.ZohoTLD,
			GetAvailabilityModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &AvailabilityResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}
	if serviceID == "" {
		return AvailabilityResponse{}, fmt.Errorf(
			"Failed to execute FetchAvailability due to non-availability of service_id",
		)
	}
	endpoint.URLParameters["service_id"] = serviceID

	if staffID == "" && resourceID == "" {
		return AvailabilityResponse{}, fmt.Errorf(
			"Failed to execute FetchAvailability due to non-availability of both staff_id and resource_id(atleast one is required)",
		)
	}
	if resourceID != "" {
		endpoint.URLParameters["resource_id"] = resourceID
	}
	if staffID != "" {
		endpoint.URLParameters["staff_id"] = staffID
	}

	if date == "" {
		return AvailabilityResponse{}, fmt.Errorf(
			"Failed to execute FetchAvailability due to non-availability of date",
		)
	}
	endpoint.URLParameters["selected_date"] = date

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AvailabilityResponse{}, fmt.Errorf("Failed to retrieve services: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*AvailabilityResponse); ok {
		return *v, nil
	}
	return AvailabilityResponse{}, fmt.Errorf("Data retrieved was not 'Service Response'")
}

type AvailabilityResponse struct {
	Response struct {
		ReturnValue struct {
			Response bool     `json:"response"`
			Data     []string `json:"data"`
			TimeZone string   `json:"time_zone"`
		} `json:"returnvalue"`
		Status string `json:"status"`
	} `json:"response"`
}
