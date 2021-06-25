package bookings

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

func (c *API) GetAppointment(request interface{}, params map[string]zoho.Parameter) (data AppointmentResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         GetAppointmentModule,
		URL:          fmt.Sprintf(BookingsAPIEndpoint+"%s?",GetAppointmentModule),
		Method:       zoho.HTTPGet,
		ResponseData: &AppointmentResponse{},
		URLParameters: map[string]zoho.Parameter{
			"filter_by": "",
		},
	}

	for k,v := range params {
		endpoint.URLParameters[k] = v;
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AppointmentResponse{}, fmt.Errorf("Failed to retrieve appointments: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*AppointmentResponse); ok {
		return *v, nil
	}
	return AppointmentResponse{}, fmt.Errorf("Data retrieved was not 'AppointmentResponse'")
}

//AppointmentResponse is the data returned by GetAppointment
type AppointmentResponse struct {
	Status string `json:"status"`
	LogMessage []string `json:"logMessage"`
	Appointments struct {
		staffName string `json:"staff_name"`
		customerMoreInfo struct{} `json:"customer_more_info"`
		customerBookingStartTime string `json:"customer_booking_start_time"`
		customerContactNo string `json:"customer_contact_no"`
		bookedOn string `json:"booked_on"`
		bookingID string `json:"booking_id"`
		workspaceId string `json:"workspace_id"`
		duration string `json:"duration"`
		serviceId string `json:"service_id"`
		staffId string `json:"staff_id"`
		costPaid string `json:"cost_paid"`
		currency string `json:"currency"`
		workspaceName string `json:"workspace_name"`
		cost string `json:"cost"`
		serviceName string `json:"service_name"`
		timeZone string `json:"time_zone"`
		startTime string `json:"start_time"`
		due string `json:"due"`
		customerEmail string `json:"customer_email"`
		bookingType string `json:"booking_type"`
		customerName string `json:"customer_name"`
		summaryUrl string `json:"summary_url"`
		customerBookingTimeZone string `json:"customer_booking_time_zone"`
		status string `json:status"`
	} `json:"returnvalue"`
}


