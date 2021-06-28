package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) FetchStaff(request interface{}, params map[string]zoho.Parameter) (data StaffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         FetchStaffModule,
		URL:          fmt.Sprintf(BookingsAPIEndpoint+"%s", FetchStaffModule),
		Method:       zoho.HTTPGet,
		ResponseData: &StaffResponse{},
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
		return StaffResponse{}, fmt.Errorf("Failed to retrieve staffs: %s", err)
	}

	if v,ok := endpoint.ResponseData.(*StaffResponse); ok {
		return *v, nil
	}
	return StaffResponse{}, fmt.Errorf("Data retrieved was not 'Staff Response'")
}

type StaffResponse struct {
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
