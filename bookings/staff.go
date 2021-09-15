package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchStaff(serviceID zoho.Parameter, staffID zoho.Parameter) (data StaffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchStaffModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s", c.ZohoTLD,FetchStaffModule),
		Method:       zoho.HTTPGet,
		ResponseData: &StaffResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}
	if serviceID != "" {
		endpoint.URLParameters["service_id"] = serviceID
	}
	if staffID != "" {
		endpoint.URLParameters["staff_id"] = staffID
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return StaffResponse{}, fmt.Errorf("Failed to retrieve staffs: %s", err)
	}

	if v,ok := endpoint.ResponseData.(*StaffResponse); ok {
		return *v, nil
	}
	return StaffResponse{}, fmt.Errorf("Data retrieved was not 'Staff Response'")
}

type StaffResponse struct {
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
