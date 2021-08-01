package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) GetAppointment(bookingID zoho.Parameter) (data AppointmentResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         GetAppointmentModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s",c.ZohoTLD, GetAppointmentModule),
		Method:       zoho.HTTPGet,
		ResponseData: &AppointmentResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},

	}
	if bookingID == "" {
		return AppointmentResponse{}, fmt.Errorf("Failed to get appointment due to non-availability of booking_id")
	}
	endpoint.URLParameters["booking_id"] = bookingID

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AppointmentResponse{}, fmt.Errorf("Failed to retrieve appointments: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*AppointmentResponse); ok {
		return *v, nil
	}
	return AppointmentResponse{}, fmt.Errorf("Data retrieved was not 'AppointmentResponse'")
}

func (c *API) BookAppointment(request BookAppointmentData) (data AppointmentResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         BookAppointmentModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s",c.ZohoTLD,BookAppointmentModule),
		Method:       zoho.HTTPPost,
		ResponseData: &AppointmentResponse{},
		RequestBody: request,
		BodyFormat: zoho.URL,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AppointmentResponse{}, fmt.Errorf("Failed to book appointment: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*AppointmentResponse); ok {

		return *v, nil
	}
	return AppointmentResponse{}, fmt.Errorf("Data retrieved was not 'AppointmentResponse'")
}

func (c *API) UpdateAppointment(request UpdateAppointmentData) (data AppointmentResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         UpdateAppointmentModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s",c.ZohoTLD,UpdateAppointmentModule),
		Method:       zoho.HTTPPost,
		ResponseData: &AppointmentResponse{},
		RequestBody: request,
		BodyFormat: zoho.URL,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AppointmentResponse{}, fmt.Errorf("Failed to update appointments: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*AppointmentResponse); ok {

		return *v, nil
	}
	return AppointmentResponse{}, fmt.Errorf("Data retrieved was not 'AppointmentResponse'")
}

func (c *API) RescheduleAppointment(request RescheduleAppointmentData) (data AppointmentResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         RescheduleAppointmentModule,
		URL:          fmt.Sprintf("https://www.zohoapis.%s/bookings/v1/json/%s",c.ZohoTLD,RescheduleAppointmentModule),
		Method:       zoho.HTTPPost,
		ResponseData: &AppointmentResponse{},
		RequestBody: request,
		BodyFormat: zoho.URL,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AppointmentResponse{}, fmt.Errorf("Failed to update appointments: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*AppointmentResponse); ok {

		return *v, nil
	}
	return AppointmentResponse{}, fmt.Errorf("Data retrieved was not 'AppointmentResponse'")
}

type UpdateAppointmentData struct {
	BookingID string `url:"booking_id"`
	Action string `url:"action"`
}

type RescheduleAppointmentData struct {
	BookingID string `url:"booking_id"`
	StaffId string `url:"staff_id,omitempty"`
	StartTime string `url:"start_time,omitempty"`
}

type CustomerDetails struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`

}

type BookAppointmentData struct {
	ServiceId string `url:"service_id"`
	StaffId string `url:"staff_id,omitempty"`
	ResourceId string `url:"resource_id,omitempty"`
	FromTime string `url:"from_time"`
	TimeZone string `url:"time_zone,omitempty"`
	Customer_Details CustomerDetails `url:"customer_details,json,omitempty"` // Note the option `json` before `omitempty`, the order shouldn't matter
}

//AppointmentResponse is the data returned by GetAppointment
type AppointmentResponse struct {
	Response struct {
		ErrorMessage string `json:"errormessage,omitempty"`
		Status string `json:"status"`
		LogMessage []string `json:"logMessage"`
		ReturnValue struct {
			StaffName string `json:"staff_name"`
			CustomerMoreInfo struct{} `json:"customer_more_info"`
			CustomerBookingStartTime string `json:"customer_booking_start_time"`
			CustomerContactNo string `json:"customer_contact_no"`
			BookedOn string `json:"booked_on"`
			BookingID string `json:"booking_id"`
			WorkspaceId string `json:"workspace_id"`
			Duration string `json:"duration"`
			ServiceId string `json:"service_id"`
			StaffId string `json:"staff_id"`
			CostPaid string `json:"cost_paid"`
			Currency string `json:"currency"`
			WorkspaceName string `json:"workspace_name"`
			Cost string `json:"cost"`
			ServiceName string `json:"service_name"`
			TimeZone string `json:"time_zone"`
			StartTime string `json:"start_time"`
			Due string `json:"due"`
			CustomerEmail string `json:"customer_email"`
			BookingType string `json:"booking_type"`
			CustomerName string `json:"customer_name"`
			SummaryUrl string `json:"summary_url"`
			CustomerBookingTimeZone string `json:"customer_booking_time_zone"`
			Status string `json:status"`
		} `json:"returnvalue"`
	} `json:"response"`
}


