package shifts

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetAllEmployees returns a list of all employees
// https://www.zoho.com/shifts/api/v1/employees-api/#get-all-employees
func (s *API) GetAllEmployees(
	params map[string]zoho.Parameter,
) (data GetEmployeesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetAllEmployees",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &GetEmployeesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"schedules":     "",
			"status":        "", // active, inactive
			"invite_status": "", // not-sent, sent, accepted
			"limit":         "50",
			"page":          "1",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetEmployeesResponse{}, fmt.Errorf("failed to retrieve empmloyees: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetEmployeesResponse); ok {
		return *v, nil
	}
	return GetEmployeesResponse{}, fmt.Errorf("data retrieved was not 'GetEmployeesResponse'")
}

type GetEmployeesResponse struct {
	Employees []struct {
		ID                string `json:"id,omitempty"`
		FirstName         string `json:"first_name,omitempty"`
		LastName          string `json:"last_name,omitempty"`
		WorkEmail         string `json:"work_email,omitempty"`
		Mobile            string `json:"mobile,omitempty"`
		MobileCountryCode string `json:"mobile_country_code,omitempty"`
		AccessLevelID     string `json:"access_level_id,omitempty"`
		Status            string `json:"status,omitempty"`
		InviteStatus      string `json:"invite_status,omitempty"`
		Schedules         []struct {
			ID string `json:"id,omitempty"`
		} `json:"schedules,omitempty"`
		Positions []struct {
			ID string `json:"id,omitempty"`
		} `json:"positions,omitempty"`
	} `json:"employees,omitempty"`
	Meta struct {
		Count int `json:"count,omitempty"`
		Limit int `json:"limit,omitempty"`
		Page  int `json:"page,omitempty"`
	} `json:"meta,omitempty"`
}

// CreateEmployee adds a new record to the list of employees
// https://www.zoho.com/shifts/api/v1/employees-api/#create-an-employee
func (s *API) CreateEmployee(
	request CreateEmployeeRequest,
) (data CreateEmployeeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "CreateEmployee",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateEmployeeResponse{},
		RequestBody:  request,
	}

	if request.FirstName == "" || len(request.Schedules) == 0 || request.Timezone == "" {
		return CreateEmployeeResponse{}, fmt.Errorf(
			"failed to create employee: first_name, schedules, and timezone are required fields",
		)
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateEmployeeResponse{}, fmt.Errorf("failed to create employee: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateEmployeeResponse); ok {
		return *v, nil
	}

	return CreateEmployeeResponse{}, fmt.Errorf("data retrieved was not 'CreateEmployeeResponse'")
}

type CreateEmployeeRequest struct {
	FirstName         string `json:"first_name"` // required
	LastName          string `json:"last_name,omitempty"`
	WorkEmail         string `json:"work_email,omitempty"`
	Mobile            string `json:"mobile,omitempty"`
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	Schedules         []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules"` // required
	Positions []struct {
		ID string `json:"id,omitempty"`
	} `json:"positions,omitempty"`
	Timezone           string `json:"timezone"` // required
	AccessLevelID      string `json:"access_level_id,omitempty"`
	SendInvitation     bool   `json:"send_invitation,omitempty"`
	HourlyRate         int    `json:"hourly_rate,omitempty"`
	HireDate           *Date  `json:"hire_date,omitempty"`
	ExternalEmployeeID string `json:"external_employee_id,omitempty"`
	HideFromSchedule   bool   `json:"hide_from_schedule,omitempty"`
	OvertimeRuleID     string `json:"overtime_rule_id,omitempty"`
	EmploymentType     string `json:"employment_type,omitempty"`
	MaxHrsWeek         int    `json:"max_hrs_week,omitempty"`
	MinHrsWeek         int    `json:"min_hrs_week,omitempty"`
	MaxHrsDay          int    `json:"max_hrs_day,omitempty"`
	MaxDaysWeek        int    `json:"max_days_week,omitempty"`
	MaxShiftsDay       int    `json:"max_shifts_day,omitempty"`
}

type CreateEmployeeResponse struct {
	ID                string `json:"id,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	WorkEmail         string `json:"work_email,omitempty"`
	Mobile            string `json:"mobile,omitempty"`
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	Schedules         []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
	Positions []struct {
		ID string `json:"id,omitempty"`
	} `json:"positions,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	AccessLevelID      string `json:"access_level_id,omitempty"`
	Status             string `json:"status,omitempty"`
	InviteStatus       string `json:"invite_status,omitempty"`
	HireDate           *Date  `json:"hire_date,omitempty"`
	ExternalEmployeeID string `json:"external_employee_id,omitempty"`
	HideFromSchedule   bool   `json:"hide_from_schedule,omitempty"`
	OvertimeRuleID     string `json:"overtime_rule_id,omitempty"`
	EmploymentType     string `json:"employment_type,omitempty"`
	MaxHrsWeek         int    `json:"max_hrs_week,omitempty"`
	MinHrsWeek         int    `json:"min_hrs_week,omitempty"`
	MaxHrsDay          int    `json:"max_hrs_day,omitempty"`
	MaxDaysWeek        int    `json:"max_days_week,omitempty"`
	MaxShiftsDay       int    `json:"max_shifts_day,omitempty"`
}

// GetEmployee retrieves the employee record with the given ID
// https://www.zoho.com/shifts/api/v1/employees-api/#get-an-employee
func (s *API) GetEmployee(id string) (data GetEmployeeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetEmployee",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/%s",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
			id,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &GetEmployeeResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetEmployeeResponse{}, fmt.Errorf("failed to retrieve Employee with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*GetEmployeeResponse); ok {
		return *v, nil
	}

	return GetEmployeeResponse{}, fmt.Errorf("data returned was not 'GetEmployeeResponse'")
}

type GetEmployeeResponse struct {
	ID                string `json:"id,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	WorkEmail         string `json:"work_email,omitempty"`
	Mobile            string `json:"mobile,omitempty"`
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	AccessLevelID     string `json:"access_level_id,omitempty"`
	Status            string `json:"status,omitempty"`
	InviteStatus      string `json:"invite_status,omitempty"`
	Schedules         []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
	Positions []struct {
		ID string `json:"id,omitempty"`
	} `json:"positions,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	HireDate           *Date  `json:"hire_date,omitempty"`
	ExternalEmployeeID string `json:"external_employee_id,omitempty"`
	HideFromSchedule   bool   `json:"hide_from_schedule,omitempty"`
	OvertimeRuleID     string `json:"overtime_rule_id,omitempty"`
	EmploymentType     string `json:"employment_type,omitempty"`
	MaxHrsWeek         int    `json:"max_hrs_week,omitempty"`
	MinHrsWeek         int    `json:"min_hrs_week,omitempty"`
	MaxHrsDay          int    `json:"max_hrs_day,omitempty"`
	MaxDaysWeek        int    `json:"max_days_week,omitempty"`
	MaxShiftsDay       int    `json:"max_shifts_day,omitempty"`
}

// UpdateEmployee modifies the employee with the given ID
// https://www.zoho.com/shifts/api/v1/employees-api/#update-an-employee
func (s *API) UpdateEmployee(
	id string,
	request UpdateEmployeeRequest,
) (data UpdateEmployeeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "UpdateEmployee",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/%s",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
			id,
		),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateEmployeeResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateEmployeeResponse{}, fmt.Errorf("failed to update employee: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateEmployeeResponse); ok {
		return *v, nil
	}

	return UpdateEmployeeResponse{}, fmt.Errorf("data retrieved was not 'UpdateEmployeeResponse'")
}

type UpdateEmployeeRequest struct {
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	WorkEmail         string `json:"work_email,omitempty"`
	Mobile            string `json:"mobile,omitempty"`
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	Schedules         []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
	Positions []struct {
		ID string `json:"id,omitempty"`
	} `json:"positions,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	AccessLevelID      string `json:"access_level_id,omitempty"`
	HourlyRate         int    `json:"hourly_rate,omitempty"`
	HireDate           *Date  `json:"hire_date,omitempty"`
	ExternalEmployeeID string `json:"external_employee_id,omitempty"`
	HideFromSchedule   bool   `json:"hide_from_schedule,omitempty"`
	OvertimeRuleID     string `json:"overtime_rule_id,omitempty"`
	EmploymentType     string `json:"employment_type,omitempty"`
	MaxHrsWeek         int    `json:"max_hrs_week,omitempty"`
	MinHrsWeek         int    `json:"min_hrs_week,omitempty"`
	MaxHrsDay          int    `json:"max_hrs_day,omitempty"`
	MaxDaysWeek        int    `json:"max_days_week,omitempty"`
	MaxShiftsDay       int    `json:"max_shifts_day,omitempty"`
}

type UpdateEmployeeResponse struct {
	ID                string `json:"id,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	WorkEmail         string `json:"work_email,omitempty"`
	Mobile            string `json:"mobile,omitempty"`
	MobileCountryCode string `json:"mobile_country_code,omitempty"`
	AccessLevelID     string `json:"access_level_id,omitempty"`
	Status            string `json:"status,omitempty"`
	InviteStatus      string `json:"invite_status,omitempty"`
	Schedules         []struct {
		ID string `json:"id,omitempty"`
	} `json:"schedules,omitempty"`
	Positions []struct {
		ID string `json:"id,omitempty"`
	} `json:"positions,omitempty"`
	Timezone           string `json:"timezone,omitempty"`
	HireDate           *Date  `json:"hire_date,omitempty"`
	ExternalEmployeeID string `json:"external_employee_id,omitempty"`
	HideFromSchedule   bool   `json:"hide_from_schedule,omitempty"`
	OvertimeRuleID     string `json:"overtime_rule_id,omitempty"`
	EmploymentType     string `json:"employment_type,omitempty"`
	MaxHrsWeek         int    `json:"max_hrs_week,omitempty"`
	MinHrsWeek         int    `json:"min_hrs_week,omitempty"`
	MaxHrsDay          int    `json:"max_hrs_day,omitempty"`
	MaxDaysWeek        int    `json:"max_days_week,omitempty"`
	MaxShiftsDay       int    `json:"max_shifts_day,omitempty"`
}

// ActivateEmployee activates the employee accounts of the employee IDs provided
// https://www.zoho.com/shifts/api/v1/employees-api/#activate-employees
func (s *API) ActivateEmployee(
	request ActivateEmployeeRequest,
) (data ActivateEmployeeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "ActivateEmployee",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/activate",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &ActivateEmployeeResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ActivateEmployeeResponse{}, fmt.Errorf("failed to activate employees: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ActivateEmployeeResponse); ok {
		return *v, nil
	}

	return ActivateEmployeeResponse{}, fmt.Errorf(
		"data retrieved was not 'ActivateEmployeeResponse'",
	)
}

type ActivateEmployeeRequest struct {
	Employees []string `json:"employees"` // required
}

type ActivateEmployeeResponse struct {
	Message string `json:"message,omitempty"`
}

// DeactivateEmployee deactivates the employee accounts of the employee IDs provided
// https://www.zoho.com/shifts/api/v1/employees-api/#deactivate-employees
func (s *API) DeactivateEmployee(
	request DeactivateEmployeeRequest,
) (data DeactivateEmployeeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "DeactivateEmployee",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/deactivate",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &DeactivateEmployeeResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeactivateEmployeeResponse{}, fmt.Errorf("failed to deactivate employees: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeactivateEmployeeResponse); ok {
		return *v, nil
	}

	return DeactivateEmployeeResponse{}, fmt.Errorf(
		"data retrieved was not 'DeactivateEmployeeResponse'",
	)
}

type DeactivateEmployeeRequest struct {
	Employees []string `json:"employees"` // required
}

type DeactivateEmployeeResponse struct {
	Message string `json:"message,omitempty"`
}

// InviteEmployee sends an invite to the employee accounts of the employee IDs provided
// https://www.zoho.com/shifts/api/v1/employees-api/#invite-employees
func (s *API) InviteEmployee(
	request InviteEmployeeRequest,
) (data InviteEmployeeResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "InviteEmployee",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/invite",
			s.ZohoTLD,
			s.OrganizationID,
			EmployeesModule,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &InviteEmployeeResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InviteEmployeeResponse{}, fmt.Errorf("failed to invite employees: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*InviteEmployeeResponse); ok {
		return *v, nil
	}

	return InviteEmployeeResponse{}, fmt.Errorf("data retrieved was not 'InviteEmployeeResponse'")
}

type InviteEmployeeRequest struct {
	Employees     []string `json:"employees"` // required
	AccessLevelID string   `json:"access_level_id,omitempty"`
}

type InviteEmployeeResponse struct {
	Message string `json:"message,omitempty"`
}
