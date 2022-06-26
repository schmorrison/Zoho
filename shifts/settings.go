package shifts

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetAllSchedules returns a list of all schedules
// https://www.zoho.com/shifts/api/v1/schedules-api/#get-all-schedules
func (s *API) GetAllSchedules(params map[string]zoho.Parameter) (data GetSchedulesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllSchedules",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, schedulesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &GetSchedulesResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetSchedulesResponse{}, fmt.Errorf("failed to retrieve schedules: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetSchedulesResponse); ok {
		return *v, nil
	}
	return GetSchedulesResponse{}, fmt.Errorf("data retrieved was not 'GetSchedulesResponse'")
}

type GetSchedulesResponse struct {
	Schedules []struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Timezone  string `json:"timezone,omitempty"`
		Address   string `json:"address,omitempty"`
		Latitude  string `json:"latitude,omitempty"`
		Longitude string `json:"longitude,omitempty"`
	} `json:"schedules,omitempty"`
}

// CreateSchedule adds a new record to the list of schedules
// https://www.zoho.com/shifts/api/v1/schedules-api/#create-a-schedule
func (s *API) CreateSchedule(request CreateScheduleRequest) (data CreateScheduleResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreateSchedule",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, schedulesModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateScheduleResponse{},
		RequestBody:  request,
	}

	if request.Name == "" {
		return CreateScheduleResponse{}, fmt.Errorf("failed to create schedule: name is a required field")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateScheduleResponse{}, fmt.Errorf("failed to create a schedule: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateScheduleResponse); ok {
		return *v, nil
	}

	return CreateScheduleResponse{}, fmt.Errorf("data retrieved was not 'CreateScheduleResponse'")
}

type CreateScheduleRequest struct {
	Name      string `json:"name"` // required
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type CreateScheduleResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

// UpdateSchedule modifies the schedule with the given ID
// https://www.zoho.com/shifts/api/v1/schedules-api/#update-a-schedule
func (s *API) UpdateSchedule(id string, request UpdateScheduleRequest) (data UpdateScheduleResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdateSchedule",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, schedulesModule, id),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateScheduleResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateScheduleResponse{}, fmt.Errorf("failed to update schedule: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateScheduleResponse); ok {
		return *v, nil
	}

	return UpdateScheduleResponse{}, fmt.Errorf("data retrieved was not 'UpdateScheduleResponse'")
}

type UpdateScheduleRequest struct {
	Name      string `json:"name,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type UpdateScheduleResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

// DeleteSchedule deletes the schedule record with the given ID
// https://www.zoho.com/shifts/api/v1/schedules-api/#delete-a-schedule
func (s *API) DeleteSchedule(id string) (data DeleteScheduleResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeleteSchedule",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, schedulesModule, id),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteScheduleResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteScheduleResponse{}, fmt.Errorf("failed to delete schedule with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteScheduleResponse); ok {
		return *v, nil
	}

	return DeleteScheduleResponse{}, fmt.Errorf("data returned was not 'DeleteScheduleResponse'")
}

type DeleteScheduleResponse struct {
	Message string `json:"message,omitempty"`
}

// GetAllPositions returns a list of all position
// https://www.zoho.com/shifts/api/v1/positions-api/#get-all-positions
func (s *API) GetAllPositions(params map[string]zoho.Parameter) (data GetPositionsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllPositions",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, positionsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &GetPositionsResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetPositionsResponse{}, fmt.Errorf("failed to retrieve positions: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetPositionsResponse); ok {
		return *v, nil
	}
	return GetPositionsResponse{}, fmt.Errorf("data retrieved was not 'GetPositionsResponse'")
}

type GetPositionsResponse struct {
	Positions []struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Color     string `json:"color,omitempty"`
		Schedules []struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"schedules,omitempty"`
	} `json:"positions,omitempty"`
}

// CreatePosition adds a new record to the list of positions
// https://www.zoho.com/shifts/api/v1/positions-api/#create-a-position
func (s *API) CreatePosition(request CreatePositionRequest) (data CreatePositionResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreatePosition",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, positionsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreatePositionResponse{},
		RequestBody:  request,
	}

	if request.Name == "" {
		return CreatePositionResponse{}, fmt.Errorf("failed to create position: name is a required field")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreatePositionResponse{}, fmt.Errorf("failed to create a position: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreatePositionResponse); ok {
		return *v, nil
	}

	return CreatePositionResponse{}, fmt.Errorf("data retrieved was not 'CreatePositionResponse'")
}

type CreatePositionRequest struct {
	Name      string `json:"name"`            // required
	Color     string `json:"color,omitempty"` // red, pink, magenta, purple, deep-purple, indigo, light-violet, blue, light-blue, cyan, muted-green, teal, green, light-green, lime, yellow, amber, orange, deep-orange, brown, grey, blue-grey
	Schedules []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
}

type CreatePositionResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Color     string `json:"color,omitempty"`
	Schedules []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"schedules,omitempty"`
}

// UpdatePosition modifies the position with the given ID
// https://www.zoho.com/shifts/api/v1/positions-api/#update-a-position
func (s *API) UpdatePosition(id string, request UpdatePositionRequest) (data UpdatePositionResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdatePosition",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, positionsModule, id),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdatePositionResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdatePositionResponse{}, fmt.Errorf("failed to update position: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdatePositionResponse); ok {
		return *v, nil
	}

	return UpdatePositionResponse{}, fmt.Errorf("data retrieved was not 'UpdatePositionResponse'")
}

type UpdatePositionRequest struct {
	Name      string `json:"name,omitempty"`
	Color     string `json:"color,omitempty"` // red, pink, magenta, purple, deep-purple, indigo, light-violet, blue, light-blue, cyan, muted-green, teal, green, light-green, lime, yellow, amber, orange, deep-orange, brown, grey, blue-grey
	Schedules []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
}

type UpdatePositionResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Color     string `json:"color,omitempty"`
	Schedules []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"schedules,omitempty"`
}

// DeletePosition deletes the schedule record with the given ID
// https://www.zoho.com/shifts/api/v1/positions-api/#delete-a-position
func (s *API) DeletePosition(id string) (data DeletePositionResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeletePosition",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, positionsModule, id),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeletePositionResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeletePositionResponse{}, fmt.Errorf("failed to delete position with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeletePositionResponse); ok {
		return *v, nil
	}

	return DeletePositionResponse{}, fmt.Errorf("data returned was not 'DeletePositionResponse'")
}

type DeletePositionResponse struct {
	Message string `json:"message,omitempty"`
}

// GetAllJobsites returns a list of all job sites
// https://www.zoho.com/shifts/api/v1/job-sites-api/#get-all-job-sites
func (s *API) GetAllJobsites(params map[string]zoho.Parameter) (data GetJobsitesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllJobsites",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, jobSitesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &GetJobsitesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"page":  "",
			"limit": "",
		},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetJobsitesResponse{}, fmt.Errorf("failed to retrieve job sites: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetJobsitesResponse); ok {
		return *v, nil
	}
	return GetJobsitesResponse{}, fmt.Errorf("data retrieved was not 'GetJobsitesResponse'")
}

type GetJobsitesResponse struct {
	JobSites []struct {
		ID        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
		Schedules []struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"schedules,omitempty"`
		Address   string `json:"address,omitempty"`
		Latitude  string `json:"latitude,omitempty"`
		Longitude string `json:"longitude,omitempty"`
		Notes     string `json:"notes,omitempty"`
	} `json:"job_sites,omitempty"`
	Meta struct {
		Count int `json:"count,omitempty"`
		Limit int `json:"limit,omitempty"`
		Page  int `json:"page,omitempty"`
	} `json:"meta,omitempty"`
}

// CreateJobsite adds a new record to the list of job sites
// https://www.zoho.com/shifts/api/v1/job-sites-api/#create-a-job-site
func (s *API) CreateJobsite(request CreateJobsiteRequest) (data CreateJobsiteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "CreateJobsite",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, jobSitesModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateJobsiteResponse{},
		RequestBody:  request,
	}

	if request.Name == "" {
		return CreateJobsiteResponse{}, fmt.Errorf("failed to create job site: name is a required field")
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateJobsiteResponse{}, fmt.Errorf("failed to create a job site: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateJobsiteResponse); ok {
		return *v, nil
	}

	return CreateJobsiteResponse{}, fmt.Errorf("data retrieved was not 'CreateJobsiteResponse'")
}

type CreateJobsiteRequest struct {
	Name      string `json:"name"` // required
	Schedules []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Notes     string `json:"notes,omitempty"`
}

type CreateJobsiteResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Schedules []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"schedules,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Notes     string `json:"notes,omitempty"`
}

// UpdateJobsite modifies the job site with the given ID
// https://www.zoho.com/shifts/api/v1/job-sites-api/#update-a-job-site
func (s *API) UpdateJobsite(id string, request UpdateJobsiteRequest) (data UpdateJobsiteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpdateJobsite",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, jobSitesModule, id),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateJobsiteResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateJobsiteResponse{}, fmt.Errorf("failed to update job site: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateJobsiteResponse); ok {
		return *v, nil
	}

	return UpdateJobsiteResponse{}, fmt.Errorf("data retrieved was not 'UpdateJobsiteResponse'")
}

type UpdateJobsiteRequest struct {
	Name      string `json:"name,omitempty"`
	Schedules []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Notes     string `json:"notes,omitempty"`
}

type UpdateJobsiteResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Schedules []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"schedules,omitempty"`
	Address   string `json:"address,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Notes     string `json:"notes,omitempty"`
}

// DeleteJobsite deletes the job site record with the given ID
// https://www.zoho.com/shifts/api/v1/job-sites-api/#delete-a-job-site
func (s *API) DeleteJobsite(id string) (data DeleteJobsiteResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeleteJobsite",
		URL:          fmt.Sprintf("https://shifts.zoho.%s/api/v1/%s/%s/%s/%s", s.ZohoTLD, s.OrganizationID, SettingsModule, jobSitesModule, id),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteJobsiteResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteJobsiteResponse{}, fmt.Errorf("failed to delete job site with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteJobsiteResponse); ok {
		return *v, nil
	}

	return DeleteJobsiteResponse{}, fmt.Errorf("data returned was not 'DeleteJobsiteResponse'")
}

type DeleteJobsiteResponse struct {
	Message string `json:"message,omitempty"`
}
