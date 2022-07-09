package shifts

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetAllShifts returns a list of all shifts
// https://www.zoho.com/shifts/api/v1/shifts-api/#get-all-shifts
func (s *API) GetAllShifts(params map[string]zoho.Parameter) (data GetShiftsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllShifts",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s", s.ZohoTLD, s.OrganizationID, shiftsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &GetShiftsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"start_date": "", // yyyy-mm-dd
			"end_date":   "", // yyyy-mm-dd
			"schedules":  "",
			"job_sites":  "",
			"positions":  "",
			"employees":  "",
			"status":     "", // published, unpublished
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	if endpoint.URLParameters["start_date"] == "" || endpoint.URLParameters["end_date"] == "" {
		return GetShiftsResponse{}, fmt.Errorf("failed to retrieve shifts: start_date and end_date are required search parameters")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetShiftsResponse{}, fmt.Errorf("failed to retrieve shifts: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*GetShiftsResponse); ok {
		return *v, nil
	}

	return GetShiftsResponse{}, fmt.Errorf("data retrieved was not 'GetShiftsResponse'")
}

type GetShiftsResponse struct {
	Shifts []struct {
		ID            string  `json:"id,omitempty"`
		StartTime     Time    `json:"start_time,omitempty"`
		EndTime       Time    `json:"end_time,omitempty"`
		EmployeeID    string  `json:"employee_id,omitempty"`
		Count         int     `json:"count,omitempty"`
		ScheduleID    string  `json:"schedule_id,omitempty"`
		PositionID    string  `json:"position_id,omitempty"`
		JobSiteID     string  `json:"job_site_id,omitempty"`
		Duration      float64 `json:"duration,omitempty"`
		BreakDuration float64 `json:"break_duration,omitempty"`
		Notes         string  `json:"notes,omitempty"`
		IsPublished   bool    `json:"is_published,omitempty"`
		IsConfirmed   bool    `json:"is_confirmed,omitempty"`
		Breaks        []struct {
			BreakID      string `json:"break_id,omitempty"`
			DurationMins int    `json:"duration_mins,omitempty"`
			StartTime    Time   `json:"start_time,omitempty"`
			EndTime      Time   `json:"end_time,omitempty"`
			IsPaid       bool   `json:"is_paid,omitempty"`
		} `json:"breaks,omitempty"`
	} `json:"shifts,omitempty"`
}

// CreateShift adds a new record to the list of shifts
// https://www.zoho.com/shifts/api/v1/shifts-api/#create-a-shift
func (s *API) CreateShift(request CreateShiftRequest) (data CreateShiftResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreateShift",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s", s.ZohoTLD, s.OrganizationID, shiftsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateShiftResponse{},
		RequestBody:  request,
	}

	if request.StartTime.IsZero() || request.EndTime.IsZero() || request.ScheduleID == "" || request.PositionID == "" {
		return CreateShiftResponse{}, fmt.Errorf("failed to create shift: start_time, end_time, schedule_id, and position_id are required fields")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateShiftResponse{}, fmt.Errorf("failed to create shift: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateShiftResponse); ok {
		return *v, nil
	}

	return CreateShiftResponse{}, fmt.Errorf("data retrieved was not 'CreateShiftResponse'")
}

type CreateShiftRequest struct {
	StartTime  Time   `json:"start_time"`            // required
	EndTime    Time   `json:"end_time"`              // required
	EmployeeID string `json:"employee_id,omitempty"` // empty creates an open shift
	Count      int    `json:"count,omitempty"`
	ScheduleID string `json:"schedule_id"` // required
	PositionID string `json:"position_id"` // required
	JobSiteID  string `json:"job_site_id,omitempty"`
	Notes      string `json:"notes,omitempty"`
	Breaks     []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
	} `json:"breaks,omitempty"`
}

type CreateShiftResponse struct {
	ID            string  `json:"id,omitempty"`
	StartTime     Time    `json:"start_time,omitempty"`
	EndTime       Time    `json:"end_time,omitempty"`
	EmployeeID    string  `json:"employee_id,omitempty"`
	Count         int     `json:"count,omitempty"`
	ScheduleID    string  `json:"schedule_id,omitempty"`
	PositionID    string  `json:"position_id,omitempty"`
	JobSiteID     string  `json:"job_site_id,omitempty"`
	Duration      float64 `json:"duration,omitempty"`
	BreakDuration float64 `json:"break_duration,omitempty"`
	Notes         string  `json:"notes,omitempty"`
	IsPublished   bool    `json:"is_published,omitempty"`
	IsConfirmed   bool    `json:"is_confirmed,omitempty"`
	Breaks        []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
		IsPaid       bool   `json:"is_paid,omitempty"`
	} `json:"breaks,omitempty"`
}

// GetShift retrieves the shift record with the given ID
// https://www.zoho.com/shifts/api/v1/shifts-api/#get-a-shift
func (s *API) GetShift(id string) (data GetShiftResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetShift",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, shiftsModule, id),
		Method:       zoho.HTTPGet,
		ResponseData: &GetShiftResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetShiftResponse{}, fmt.Errorf("failed to retrieve shift with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*GetShiftResponse); ok {
		return *v, nil
	}

	return GetShiftResponse{}, fmt.Errorf("data returned was not 'GetShiftResponse'")
}

type GetShiftResponse struct {
	ID            string  `json:"id,omitempty"`
	StartTime     Time    `json:"start_time,omitempty"`
	EndTime       Time    `json:"end_time,omitempty"`
	EmployeeID    string  `json:"employee_id,omitempty"`
	Count         int     `json:"count,omitempty"`
	ScheduleID    string  `json:"schedule_id,omitempty"`
	PositionID    string  `json:"position_id,omitempty"`
	JobSiteID     string  `json:"job_site_id,omitempty"`
	Duration      float64 `json:"duration,omitempty"`
	BreakDuration float64 `json:"break_duration,omitempty"`
	Notes         string  `json:"notes,omitempty"`
	IsPublished   bool    `json:"is_published,omitempty"`
	IsConfirmed   bool    `json:"is_confirmed,omitempty"`
	Breaks        []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
		IsPaid       bool   `json:"is_paid,omitempty"`
	} `json:"breaks,omitempty"`
}

// UpdateShift modifies the shift with the given ID
// https://www.zoho.com/shifts/api/v1/shifts-api/#update-a-shift
func (s *API) UpdateShift(id string, request UpdateShiftRequest) (data UpdateShiftResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdateShift",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, shiftsModule, id),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateShiftResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateShiftResponse{}, fmt.Errorf("failed to update shift: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateShiftResponse); ok {
		return *v, nil
	}

	return UpdateShiftResponse{}, fmt.Errorf("data retrieved was not 'UpdateShiftResponse'")
}

type UpdateShiftRequest struct {
	StartTime  Time   `json:"start_time,omitempty"`
	EndTime    Time   `json:"end_time,omitempty"`
	EmployeeID string `json:"employee_id,omitempty"` // empty creates an open shift
	Count      int    `json:"count,omitempty"`
	ScheduleID string `json:"schedule_id,omitempty"`
	PositionID string `json:"position_id,omitempty"`
	JobSiteID  string `json:"job_site_id,omitempty"`
	Notes      string `json:"notes,omitempty"`
	Breaks     []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
	} `json:"breaks,omitempty"`
}

type UpdateShiftResponse struct {
	ID            string  `json:"id,omitempty"`
	StartTime     Time    `json:"start_time,omitempty"`
	EndTime       Time    `json:"end_time,omitempty"`
	EmployeeID    string  `json:"employee_id,omitempty"`
	Count         int     `json:"count,omitempty"`
	ScheduleID    string  `json:"schedule_id,omitempty"`
	PositionID    string  `json:"position_id,omitempty"`
	JobSiteID     string  `json:"job_site_id,omitempty"`
	Duration      float64 `json:"duration,omitempty"`
	BreakDuration float64 `json:"break_duration,omitempty"`
	Notes         string  `json:"notes,omitempty"`
	IsPublished   bool    `json:"is_published,omitempty"`
	IsConfirmed   bool    `json:"is_confirmed,omitempty"`
	Breaks        []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
		IsPaid       bool   `json:"is_paid,omitempty"`
	} `json:"breaks,omitempty"`
}

// DeleteShift deletes the shift record with the given ID
// https://www.zoho.com/shifts/api/v1/shifts-api/#delete-a-shift
func (s *API) DeleteShift(id string) (data DeleteShiftResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeleteShift",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, shiftsModule, id),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteShiftResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteShiftResponse{}, fmt.Errorf("failed to delete shift with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteShiftResponse); ok {
		return *v, nil
	}

	return DeleteShiftResponse{}, fmt.Errorf("data returned was not 'DeleteShiftResponse'")
}

type DeleteShiftResponse struct {
	Message string `json:"message,omitempty"`
}

// GetAllAvailabilities returns a list of all employee availabilities
// https://www.zoho.com/shifts/api/v1/availability-api/#get-all-availabilities
func (s *API) GetAllAvailabilities(params map[string]zoho.Parameter) (data GetAvailabilitiesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllAvailabilities",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s", s.ZohoTLD, s.OrganizationID, availabilityModule),
		Method:       zoho.HTTPGet,
		ResponseData: &GetAvailabilitiesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"start_date": "", // yyyy-mm-dd
			"end_date":   "", // yyyy-mm-dd
			"employees":  "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	if endpoint.URLParameters["start_date"] == "" || endpoint.URLParameters["end_date"] == "" {
		return GetAvailabilitiesResponse{}, fmt.Errorf("failed to retreive availabilities: start_date and end_date are required fields")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetAvailabilitiesResponse{}, fmt.Errorf("failed to retrieve availabilities: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetAvailabilitiesResponse); ok {
		return *v, nil
	}
	return GetAvailabilitiesResponse{}, fmt.Errorf("data retrieved was not 'GetAvailabilitiesResponse'")
}

type GetAvailabilitiesResponse struct {
	Availabilities []struct {
		ID         string `json:"id,omitempty"`
		EmployeeID string `json:"employee_id,omitempty"`
		StartTime  Time   `json:"start_time"`
		EndTime    Time   `json:"end_time"`
		Preference string `json:"preference,omitempty"`
		Notes      string `json:"notes,omitempty"`
	} `json:"availabilities,omitempty"`
}

// CreateAvailability adds a new record to the list of available shifts
// https://www.zoho.com/shifts/api/v1/availability-api/#create-an-availability
func (s *API) CreateAvailability(request CreateAvailabilityRequest) (data CreateAvailabilityResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreateAvailability",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s", s.ZohoTLD, s.OrganizationID, availabilityModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateAvailabilityResponse{},
		RequestBody:  request,
	}

	if request.StartTime.IsZero() || request.EndTime.IsZero() || request.EmployeeID == "" || request.Preference == "" {
		return CreateAvailabilityResponse{}, fmt.Errorf("failed to create availability: start_time, end_time, employee_id, and preference are required fields")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateAvailabilityResponse{}, fmt.Errorf("failed to create an availability: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateAvailabilityResponse); ok {
		return *v, nil
	}

	return CreateAvailabilityResponse{}, fmt.Errorf("data retrieved was not 'CreateAvailabilityResponse'")
}

type CreateAvailabilityRequest struct {
	EmployeeID string `json:"employee_id"` // required
	StartTime  Time   `json:"start_time"`  // required
	EndTime    Time   `json:"end_time"`    // required
	Preference string `json:"preference"`  // required: preferred, unavailable
	Notes      string `json:"notes,omitempty"`
}

type CreateAvailabilityResponse struct {
	ID         string `json:"id,omitempty"`
	EmployeeID string `json:"employee_id,omitempty"`
	StartTime  Time   `json:"start_time,omitempty"`
	EndTime    Time   `json:"end_time,omitempty"`
	Preference string `json:"preference,omitempty"`
	Notes      string `json:"notes,omitempty"`
}

// UpdateAvailability modifies the availability with the given ID
// https://www.zoho.com/shifts/api/v1/availability-api/#update-an-availability
func (s *API) UpdateAvailability(id string, request UpdateAvailabilityRequest) (data UpdateAvailabilityResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdateAvailability",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, availabilityModule, id),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateAvailabilityResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateAvailabilityResponse{}, fmt.Errorf("failed to update availability: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateAvailabilityResponse); ok {
		return *v, nil
	}

	return UpdateAvailabilityResponse{}, fmt.Errorf("data retrieved was not 'UpdateAvailabilityResponse'")
}

type UpdateAvailabilityRequest struct {
	ID         string `json:"id,omitempty"`
	EmployeeID string `json:"employee_id,omitempty"`
	StartTime  Time   `json:"start_time,omitempty"`
	EndTime    Time   `json:"end_time,omitempty"`
	Preference string `json:"preference,omitempty"` // preferred, unavailable
	Notes      string `json:"notes,omitempty"`
}

type UpdateAvailabilityResponse struct {
	EmployeeID string `json:"employee_id,omitempty"`
	StartTime  Time   `json:"start_time,omitempty"`
	EndTime    Time   `json:"end_time,omitempty"`
	Preference string `json:"preference,omitempty"`
	Notes      string `json:"notes,omitempty"`
}

// DeleteAvailability deletes the availability record with the given ID
// https://www.zoho.com/shifts/api/v1/availability-api/#delete-an-availability
func (s *API) DeleteAvailability(id string) (data DeleteAvailabilityResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeleteAvailability",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, availabilityModule, id),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteAvailabilityResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteAvailabilityResponse{}, fmt.Errorf("failed to delete availability with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteAvailabilityResponse); ok {
		return *v, nil
	}

	return DeleteAvailabilityResponse{}, fmt.Errorf("data returned was not 'DeletAavailabilityResponse'")
}

type DeleteAvailabilityResponse struct {
	Message string `json:"message,omitempty"`
}
