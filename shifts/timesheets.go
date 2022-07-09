package shifts

import (
	"fmt"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// GetAllTimesheets returns a list of all employee timesheets
// https://www.zoho.com/shifts/api/v1/timesheets-api/#get-all-time-entries
func (s *API) GetAllTimesheets(params map[string]zoho.Parameter) (data GetTimesheetsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllTimesheets",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s", s.ZohoTLD, s.OrganizationID, TimesheetsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &GetTimesheetsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"schedules":   "",
			"status":      "", // pending, approved
			"employee_id": "",
			"start_date":  "", // yyyy-mm-dd
			"end_date":    "", // yyyy-mm-dd
			"limit":       "50",
			"page":        "1",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetTimesheetsResponse{}, fmt.Errorf("failed to retrieve timesheets: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetTimesheetsResponse); ok {
		return *v, nil
	}
	return GetTimesheetsResponse{}, fmt.Errorf("data retrieved was not 'GetTimesheetsResponse'")
}

type GetTimesheetsResponse struct {
	TimeEntries []struct {
		ID               string  `json:"id,omitempty"`
		StartTime        Time    `json:"start_time,omitempty"`
		EndTime          Time    `json:"end_time,omitempty"`
		EmployeeID       string  `json:"employee_id,omitempty"`
		Employee         string  `json:"employee,omitempty"`
		ScheduleID       string  `json:"schedule_id,omitempty"`
		Schedule         string  `json:"schedule,omitempty"`
		PositionID       string  `json:"position_id,omitempty"`
		Position         string  `json:"position,omitempty"`
		JobSiteID        string  `json:"job_site_id,omitempty"`
		JobSite          string  `json:"job_site,omitempty"`
		Type             string  `json:"type,omitempty"`
		Duration         float64 `json:"duration,string,omitempty"`
		BreakDuration    float64 `json:"break_duration,string,omitempty"`
		Notes            string  `json:"notes,omitempty"`
		Status           string  `json:"status,omitempty"`
		TimeoffRequestID string  `json:"timeoff_request_id,omitempty"`
		TimeoffTypeID    string  `json:"timeoff_type_id,omitempty"`
		TimeoffType      string  `json:"timeoff_type,omitempty"`
		InLat            float64 `json:"in_lat,omitempty"`
		InLng            float64 `json:"in_lng,omitempty"`
		OutLat           float64 `json:"out_lat,omitempty"`
		OutLng           float64 `json:"out_lng,omitempty"`
		Latitude         float64 `json:"latitude,string,omitempty"`
		Longitude        float64 `json:"longitude,string,omitempty"`
		ShiftID          string  `json:"shift_id,omitempty"`
		Breaks           []struct {
			BreakID      string `json:"break_id,omitempty"`
			BreakName    string `json:"break_name,omitempty"`
			DurationMins int    `json:"duration_mins,omitempty"`
			StartTime    Time   `json:"start_time,omitempty"`
			EndTime      Time   `json:"end_time,omitempty"`
			IsPaid       bool   `json:"is_paid,omitempty"`
		} `json:"breaks,omitempty"`
	} `json:"time_entries,omitempty"`
	Meta struct {
		Count int `json:"count,omitempty"`
		Limit int `json:"limit,omitempty"`
		Page  int `json:"page,omitempty"`
	} `json:"meta,omitempty"`
}

// CreateTimesheet adds a new record to the list of employee timesheets
// https://www.zoho.com/shifts/api/v1/timesheets-api/#create-a-time-entry
func (s *API) CreateTimesheet(request CreateTimesheetRequest) (data CreateTimesheetResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreateTimesheet",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s", s.ZohoTLD, s.OrganizationID, TimesheetsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateTimesheetResponse{},
		RequestBody:  request,
	}

	if request.StartTime.IsZero() || request.EmployeeID == "" || request.ScheduleID == "" || request.PositionID == "" {
		return CreateTimesheetResponse{}, fmt.Errorf("failed to create timesheet: start_time, employee_id, schedule_id, and position_id are required fields")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateTimesheetResponse{}, fmt.Errorf("failed to create timesheet: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateTimesheetResponse); ok {
		return *v, nil
	}

	return CreateTimesheetResponse{}, fmt.Errorf("data retrieved was not 'CreateTimesheetResponse'")
}

type CreateTimesheetRequest struct {
	StartTime  time.Time `json:"start_time"` // required
	EndTime    time.Time `json:"end_time,omitempty"`
	EmployeeID string    `json:"employee_id"` // required
	ScheduleID string    `json:"schedule_id"` // required
	PositionID string    `json:"position_id"` // required
	JobSiteID  string    `json:"job_site_id,omitempty"`
	Notes      string    `json:"notes,omitempty"`
	ShiftID    string    `json:"shift_id,omitempty"`
	Breaks     []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
	} `json:"breaks,omitempty"`
}

type CreateTimesheetResponse struct {
	ID               string  `json:"id,omitempty"`
	StartTime        Time    `json:"start_time,omitempty"`
	EndTime          Time    `json:"end_time,omitempty"`
	EmployeeID       string  `json:"employee_id,omitempty"`
	Employee         string  `json:"employee,omitempty"`
	ScheduleID       string  `json:"schedule_id,omitempty"`
	Schedule         string  `json:"schedule,omitempty"`
	PositionID       string  `json:"position_id,omitempty"`
	Position         string  `json:"position,omitempty"`
	JobSiteID        string  `json:"job_site_id,omitempty"`
	JobSite          string  `json:"job_site,omitempty"`
	Type             string  `json:"type,omitempty"`
	Duration         float64 `json:"duration,string,omitempty"`
	BreakDuration    float64 `json:"break_duration,string,omitempty"`
	Notes            string  `json:"notes,omitempty"`
	Status           string  `json:"status,omitempty"`
	TimeoffRequestID string  `json:"timeoff_request_id,omitempty"`
	TimeoffTypeID    string  `json:"timeoff_type_id,omitempty"`
	TimeoffType      string  `json:"timeoff_type,omitempty"`
	InLat            float64 `json:"in_lat,omitempty"`
	InLng            float64 `json:"in_lng,omitempty"`
	OutLat           float64 `json:"out_lat,omitempty"`
	OutLng           float64 `json:"out_lng,omitempty"`
	Latitude         float64 `json:"latitude,string,omitempty"`
	Longitude        float64 `json:"longitude,string,omitempty"`
	ShiftID          string  `json:"shift_id,omitempty"`
	Breaks           []struct {
		BreakID      string `json:"break_id,omitempty"`
		BreakName    string `json:"break_name,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
		IsPaid       bool   `json:"is_paid,omitempty"`
	} `json:"breaks,omitempty"`
}

// GetTimesheet retrieves the timesheet record with the given ID
// https://www.zoho.com/shifts/api/v1/timesheets-api/#get-a-time-entry
func (s *API) GetTimesheet(id string) (data GetTimesheetResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetTimesheet",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, TimesheetsModule, id),
		Method:       zoho.HTTPGet,
		ResponseData: &GetTimesheetResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetTimesheetResponse{}, fmt.Errorf("failed to retrieve timesheet with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*GetTimesheetResponse); ok {
		return *v, nil
	}

	return GetTimesheetResponse{}, fmt.Errorf("data returned was not 'GetTimesheetResponse'")
}

type GetTimesheetResponse struct {
	ID               string  `json:"id,omitempty"`
	StartTime        Time    `json:"start_time,omitempty"`
	EndTime          Time    `json:"end_time,omitempty"`
	EmployeeID       string  `json:"employee_id,omitempty"`
	Employee         string  `json:"employee,omitempty"`
	ScheduleID       string  `json:"schedule_id,omitempty"`
	Schedule         string  `json:"schedule,omitempty"`
	PositionID       string  `json:"position_id,omitempty"`
	Position         string  `json:"position,omitempty"`
	JobSiteID        string  `json:"job_site_id,omitempty"`
	JobSite          string  `json:"job_site,omitempty"`
	Type             string  `json:"type,omitempty"`
	Duration         float64 `json:"duration,string,omitempty"`
	BreakDuration    float64 `json:"break_duration,string,omitempty"`
	Notes            string  `json:"notes,omitempty"`
	Status           string  `json:"status,omitempty"`
	TimeoffRequestID string  `json:"timeoff_request_id,omitempty"`
	TimeoffTypeID    string  `json:"timeoff_type_id,omitempty"`
	TimeoffType      string  `json:"timeoff_type,omitempty"`
	InLat            float64 `json:"in_lat,omitempty"`
	InLng            float64 `json:"in_lng,omitempty"`
	OutLat           float64 `json:"out_lat,omitempty"`
	OutLng           float64 `json:"out_lng,omitempty"`
	Latitude         float64 `json:"latitude,string,omitempty"`
	Longitude        float64 `json:"longitude,string,omitempty"`
	ShiftID          string  `json:"shift_id,omitempty"`
	Breaks           []struct {
		BreakID      string `json:"break_id,omitempty"`
		BreakName    string `json:"break_name,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
		IsPaid       bool   `json:"is_paid,omitempty"`
	} `json:"breaks,omitempty"`
}

// UpdateTimesheet modifies the timesheet with the given ID
// https://www.zoho.com/shifts/api/v1/timesheets-api/#update-a-time-entry
func (s *API) UpdateTimesheet(id string, request UpdateTimesheetRequest) (data UpdateTimesheetResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdateTimesheet",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, TimesheetsModule, id),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateTimesheetResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateTimesheetResponse{}, fmt.Errorf("failed to update timesheet: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateTimesheetResponse); ok {
		return *v, nil
	}

	return UpdateTimesheetResponse{}, fmt.Errorf("data retrieved was not 'UpdateTimesheetResponse'")
}

type UpdateTimesheetRequest struct {
	StartTime  Time   `json:"start_time,omitempty"`
	EndTime    Time   `json:"end_time,omitempty"`
	EmployeeID string `json:"employee_id,omitempty"`
	ScheduleID string `json:"schedule_id,omitempty"`
	PositionID string `json:"position_id,omitempty"`
	JobSiteID  string `json:"job_site_id,omitempty"`
	Notes      string `json:"notes,omitempty"`
	ShiftID    string `json:"shift_id,omitempty"`
	Breaks     []struct {
		BreakID      string `json:"break_id,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
	} `json:"breaks,omitempty"`
}

type UpdateTimesheetResponse struct {
	ID               string  `json:"id,omitempty"`
	StartTime        Time    `json:"start_time,omitempty"`
	EndTime          Time    `json:"end_time,omitempty"`
	EmployeeID       string  `json:"employee_id,omitempty"`
	Employee         string  `json:"employee,omitempty"`
	ScheduleID       string  `json:"schedule_id,omitempty"`
	Schedule         string  `json:"schedule,omitempty"`
	PositionID       string  `json:"position_id,omitempty"`
	Position         string  `json:"position,omitempty"`
	JobSiteID        string  `json:"job_site_id,omitempty"`
	JobSite          string  `json:"job_site,omitempty"`
	Type             string  `json:"type,omitempty"`
	Duration         float64 `json:"duration,string,omitempty"`
	BreakDuration    float64 `json:"break_duration,string,omitempty"`
	Notes            string  `json:"notes,omitempty"`
	Status           string  `json:"status,omitempty"`
	TimeoffRequestID string  `json:"timeoff_request_id,omitempty"`
	TimeoffTypeID    string  `json:"timeoff_type_id,omitempty"`
	TimeoffType      string  `json:"timeoff_type,omitempty"`
	InLat            float64 `json:"in_lat,omitempty"`
	InLng            float64 `json:"in_lng,omitempty"`
	OutLat           float64 `json:"out_lat,omitempty"`
	OutLng           float64 `json:"out_lng,omitempty"`
	Latitude         float64 `json:"latitude,string,omitempty"`
	Longitude        float64 `json:"longitude,string,omitempty"`
	ShiftID          string  `json:"shift_id,omitempty"`
	Breaks           []struct {
		BreakID      string `json:"break_id,omitempty"`
		BreakName    string `json:"break_name,omitempty"`
		DurationMins int    `json:"duration_mins,omitempty"`
		StartTime    Time   `json:"start_time,omitempty"`
		EndTime      Time   `json:"end_time,omitempty"`
		IsPaid       bool   `json:"is_paid,omitempty"`
	} `json:"breaks,omitempty"`
}

// DeleteTimesheet deletes the timesheet record with the given ID
// https://www.zoho.com/shifts/api/v1/timesheets-api/#delete-a-time-entry
func (s *API) DeleteTimesheet(id string) (data DeleteTimesheetResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeleteTimesheet",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, TimesheetsModule, id),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteTimesheetResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteTimesheetResponse{}, fmt.Errorf("failed to delete timesheet with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteTimesheetResponse); ok {
		return *v, nil
	}

	return DeleteTimesheetResponse{}, fmt.Errorf("data returned was not 'DeleteTimesheetResponse'")
}

type DeleteTimesheetResponse struct {
	Message string `json:"message,omitempty"`
}
