package shifts

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetAllTimeoffRequests returns a list of all employee timeoff requests
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#get-all-time-off-requests
func (s *API) GetAllTimeoffRequests(
	params map[string]zoho.Parameter,
) (data GetTimeoffsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetAllTimeoffRequests",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &GetTimeoffsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"schedules":   "",
			"status":      "", // pending, approved, denied, cancelled
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
		return GetTimeoffsResponse{}, fmt.Errorf("failed to retrieve timeoff requests: %s", err)
	}
	if v, ok := endpoint.ResponseData.(*GetTimeoffsResponse); ok {
		return *v, nil
	}
	return GetTimeoffsResponse{}, fmt.Errorf("data retrieved was not 'GetTimeoffsResponse'")
}

type GetTimeoffsResponse struct {
	TimeOffRequests []struct {
		ID            string  `json:"id,omitempty"`
		StartDate     *Time   `json:"start_date,omitempty"`
		EndDate       *Time   `json:"end_date,omitempty"`
		EmployeeID    string  `json:"employee_id,omitempty"`
		Employee      string  `json:"employee,omitempty"`
		RequestedByID string  `json:"requested_by_id,omitempty"`
		RequestedBy   string  `json:"requested_by,omitempty"`
		TypeID        string  `json:"type_id,omitempty"`
		Type          string  `json:"type,omitempty"`
		DayType       string  `json:"day_type,omitempty"`
		Duration      float64 `json:"duration,omitempty"`
		Status        string  `json:"status,omitempty"`
		CreatedAt     *Time   `json:"created_at,omitempty"`
	} `json:"time_off_requests,omitempty"`
	Meta struct {
		Count int `json:"count,omitempty"`
		Limit int `json:"limit,omitempty"`
		Page  int `json:"page,omitempty"`
	} `json:"meta,omitempty"`
}

// CreateTimeoffRequest adds a new record to the list of employee timeoff requests
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#create-a-time-off-request
func (s *API) CreateTimeoffRequest(
	request CreateTimeoffRequest,
) (data CreateTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "CreateTimeoffRequest",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateTimeoffResponse{},
		RequestBody:  request,
	}

	if request.StartDate.IsZero() || request.EndDate.IsZero() || request.EmployeeID == "" ||
		request.TypeID == "" ||
		request.DayType == "" {
		return CreateTimeoffResponse{}, fmt.Errorf(
			"failed to create employee: start_date, end_date, employee_id, type_id, and day_type are required fields",
		)
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CreateTimeoffResponse{}, fmt.Errorf("failed to create timeoff request: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateTimeoffResponse); ok {
		return *v, nil
	}

	return CreateTimeoffResponse{}, fmt.Errorf("data retrieved was not 'CreateTimeoffResponse'")
}

type CreateTimeoffRequest struct {
	EmployeeID string `json:"employee_id,omitempty"`
	StartDate  *Time  `json:"start_date"` // required
	EndDate    *Time  `json:"end_date"`   // required
	TypeID     string `json:"type_id"`    // required
	DayType    string `json:"day_type"`   // required: all_day, partial
	Reason     string `json:"reason,omitempty"`
}

type CreateTimeoffResponse struct {
	ID             string  `json:"id,omitempty"`
	StartDate      *Time   `json:"start_date,omitempty"`
	EndDate        *Time   `json:"end_date,omitempty"`
	EmployeeID     string  `json:"employee_id,omitempty"`
	Employee       string  `json:"employee,omitempty"`
	RequestedByID  string  `json:"requested_by_id,omitempty"`
	RequestedBy    string  `json:"requested_by,omitempty"`
	TypeID         string  `json:"type_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	DayType        string  `json:"day_type,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Status         string  `json:"status,omitempty"`
	CreatedAt      *Time   `json:"created_at,omitempty"`
	IsPaid         bool    `json:"is_paid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	ApproverID     string  `json:"approver_id,omitempty"`
	Approver       string  `json:"approver,omitempty"`
	ApprovalRespAt *Time   `json:"approval_resp_at,omitempty"`
	CancelledAt    *Time   `json:"cancelled_at,omitempty"`
	UpdatedAt      *Time   `json:"updated_at,omitempty"`
	Comments       []struct {
		CommentID   string `json:"comment_id,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CommenterID string `json:"commenter_id,omitempty"`
		Commenter   string `json:"commenter,omitempty"`
		CreatedAt   *Time  `json:"created_at,omitempty"`
	} `json:"comments,omitempty"`
}

// GetTimeoffRequest retrieves the timeoff request record with the given ID
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#get-a-time-off-request
func (s *API) GetTimeoffRequest(id string) (data GetTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetTimeoffRequest",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests/%s",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
			id,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &GetTimeoffResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return GetTimeoffResponse{}, fmt.Errorf(
			"failed to retrieve timeoff request with id: %s",
			err,
		)
	}

	if v, ok := endpoint.ResponseData.(*GetTimeoffResponse); ok {
		return *v, nil
	}

	return GetTimeoffResponse{}, fmt.Errorf("data returned was not 'GetTimeoffResponse'")
}

type GetTimeoffResponse struct {
	ID             string  `json:"id,omitempty"`
	StartDate      *Time   `json:"start_date,omitempty"`
	EndDate        *Time   `json:"end_date,omitempty"`
	EmployeeID     string  `json:"employee_id,omitempty"`
	Employee       string  `json:"employee,omitempty"`
	RequestedByID  string  `json:"requested_by_id,omitempty"`
	RequestedBy    string  `json:"requested_by,omitempty"`
	TypeID         string  `json:"type_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	DayType        string  `json:"day_type,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Status         string  `json:"status,omitempty"`
	CreatedAt      *Time   `json:"created_at,omitempty"`
	IsPaid         bool    `json:"is_paid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	ApproverID     string  `json:"approver_id,omitempty"`
	Approver       string  `json:"approver,omitempty"`
	ApprovalRespAt *Time   `json:"approval_resp_at,omitempty"`
	CancelledAt    *Time   `json:"cancelled_at,omitempty"`
	UpdatedAt      *Time   `json:"updated_at,omitempty"`
	Comments       []struct {
		CommentID   string `json:"comment_id,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CommenterID string `json:"commenter_id,omitempty"`
		Commenter   string `json:"commenter,omitempty"`
		CreatedAt   *Time  `json:"created_at,omitempty"`
	} `json:"comments,omitempty"`
}

// UpdateTimeoffRequest modifies the timeoff request with the given ID
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#update-a-time-off-request
func (s *API) UpdateTimeoff(
	id string,
	request UpdateTimeoffRequest,
) (data UpdateTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "UpdateTimeoff",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests/%s",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
			id,
		),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateTimeoffResponse{},
		RequestBody:  request,
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateTimeoffResponse{}, fmt.Errorf("failed to update timeoff request: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateTimeoffResponse); ok {
		return *v, nil
	}

	return UpdateTimeoffResponse{}, fmt.Errorf("data retrieved was not 'UpdateTimeoffResponse'")
}

type UpdateTimeoffRequest struct {
	StartDate *Time  `json:"start_date,omitempty"`
	EndDate   *Time  `json:"end_date,omitempty"`
	TypeID    string `json:"type_id,omitempty"`
	DayType   string `json:"day_type,omitempty"` // all_day, partial
	Reason    string `json:"reason,omitempty"`
}

type UpdateTimeoffResponse struct {
	ID             string  `json:"id,omitempty"`
	StartDate      *Time   `json:"start_date,omitempty"`
	EndDate        *Time   `json:"end_date,omitempty"`
	EmployeeID     string  `json:"employee_id,omitempty"`
	Employee       string  `json:"employee,omitempty"`
	RequestedByID  string  `json:"requested_by_id,omitempty"`
	RequestedBy    string  `json:"requested_by,omitempty"`
	TypeID         string  `json:"type_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	DayType        string  `json:"day_type,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Status         string  `json:"status,omitempty"`
	CreatedAt      *Time   `json:"created_at,omitempty"`
	IsPaid         bool    `json:"is_paid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	ApproverID     string  `json:"approver_id,omitempty"`
	Approver       string  `json:"approver,omitempty"`
	ApprovalRespAt *Time   `json:"approval_resp_at,omitempty"`
	CancelledAt    *Time   `json:"cancelled_at,omitempty"`
	UpdatedAt      *Time   `json:"updated_at,omitempty"`
	Comments       []struct {
		CommentID   string `json:"comment_id,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CommenterID string `json:"commenter_id,omitempty"`
		Commenter   string `json:"commenter,omitempty"`
		CreatedAt   *Time  `json:"created_at,omitempty"`
	} `json:"comments,omitempty"`
}

// DeleteTimeoffRequest deletes the timeoff request record with the given ID
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#delete-a-time-off-request
func (s *API) DeleteTimeoffRequest(id string) (data DeleteTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "DeleteTimeoffRequest",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests/%s",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
			id,
		),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteTimeoffResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteTimeoffResponse{}, fmt.Errorf(
			"failed to delete timeoff request with id: %s",
			err,
		)
	}

	if v, ok := endpoint.ResponseData.(*DeleteTimeoffResponse); ok {
		return *v, nil
	}

	return DeleteTimeoffResponse{}, fmt.Errorf("data returned was not 'DeleteTimeoffResponse'")
}

type DeleteTimeoffResponse struct {
	Message string `json:"message,omitempty"`
}

// CancelTimeoffRequest cancels the timeoff request with the given id
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#cancel-a-time-off-request
func (s *API) CancelTimeoffRequest(id string) (data CancelTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "CancelTimeoffRequest",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests/%s/cancel",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
			id,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &CancelTimeoffResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CancelTimeoffResponse{}, fmt.Errorf("failed to cancel timeoff request: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CancelTimeoffResponse); ok {
		return *v, nil
	}

	return CancelTimeoffResponse{}, fmt.Errorf("data retrieved was not 'CancelTimeoffResponse'")
}

type CancelTimeoffResponse struct {
	ID             string  `json:"id,omitempty"`
	StartDate      *Time   `json:"start_date,omitempty"`
	EndDate        *Time   `json:"end_date,omitempty"`
	EmployeeID     string  `json:"employee_id,omitempty"`
	Employee       string  `json:"employee,omitempty"`
	RequestedByID  string  `json:"requested_by_id,omitempty"`
	RequestedBy    string  `json:"requested_by,omitempty"`
	TypeID         string  `json:"type_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	DayType        string  `json:"day_type,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Status         string  `json:"status,omitempty"`
	CreatedAt      *Time   `json:"created_at,omitempty"`
	IsPaid         bool    `json:"is_paid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	ApproverID     string  `json:"approver_id,omitempty"`
	Approver       string  `json:"approver,omitempty"`
	ApprovalRespAt *Time   `json:"approval_resp_at,omitempty"`
	CancelledAt    *Time   `json:"cancelled_at,omitempty"`
	UpdatedAt      *Time   `json:"updated_at,omitempty"`
	Comments       []struct {
		CommentID   string `json:"comment_id,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CommenterID string `json:"commenter_id,omitempty"`
		Commenter   string `json:"commenter,omitempty"`
		CreatedAt   *Time  `json:"created_at,omitempty"`
	} `json:"comments,omitempty"`
}

// ApproveTimeoffRequest approves the timeoff request with the given id
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#approve-a-time-off-request
func (s *API) ApproveTimeoffRequest(id string) (data ApproveTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "ApproveTimeoffRequest",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests/%s/approve",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
			id,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &ApproveTimeoffResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ApproveTimeoffResponse{}, fmt.Errorf("failed to approve timeoff request: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ApproveTimeoffResponse); ok {
		return *v, nil
	}

	return ApproveTimeoffResponse{}, fmt.Errorf("data retrieved was not 'ApproveTimeoffResponse'")
}

type ApproveTimeoffResponse struct {
	ID             string  `json:"id,omitempty"`
	StartDate      *Time   `json:"start_date,omitempty"`
	EndDate        *Time   `json:"end_date,omitempty"`
	EmployeeID     string  `json:"employee_id,omitempty"`
	Employee       string  `json:"employee,omitempty"`
	RequestedByID  string  `json:"requested_by_id,omitempty"`
	RequestedBy    string  `json:"requested_by,omitempty"`
	TypeID         string  `json:"type_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	DayType        string  `json:"day_type,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Status         string  `json:"status,omitempty"`
	CreatedAt      *Time   `json:"created_at,omitempty"`
	IsPaid         bool    `json:"is_paid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	ApproverID     string  `json:"approver_id,omitempty"`
	Approver       string  `json:"approver,omitempty"`
	ApprovalRespAt *Time   `json:"approval_resp_at,omitempty"`
	CancelledAt    *Time   `json:"cancelled_at,omitempty"`
	UpdatedAt      *Time   `json:"updated_at,omitempty"`
	Comments       []struct {
		CommentID   string `json:"comment_id,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CommenterID string `json:"commenter_id,omitempty"`
		Commenter   string `json:"commenter,omitempty"`
		CreatedAt   *Time  `json:"created_at,omitempty"`
	} `json:"comments,omitempty"`
}

// DenyTimeoffRequest denies the timeoff request with the given id
// https://www.zoho.com/shifts/api/v1/time-off-requests-api/#deny-a-time-off-request
func (s *API) DenyTimeoffRequest(id string) (data DenyTimeoffResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "DenyTimeoffRequest",
		URL: fmt.Sprintf(
			"https://shifts.zoho.%s/api/v1/%s/%s/requests/%s/deny",
			s.ZohoTLD,
			s.OrganizationID,
			TimeoffModule,
			id,
		),
		Method:       zoho.HTTPPost,
		ResponseData: &DenyTimeoffResponse{},
	}

	err = s.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DenyTimeoffResponse{}, fmt.Errorf("failed to deny timeoff request: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DenyTimeoffResponse); ok {
		return *v, nil
	}

	return DenyTimeoffResponse{}, fmt.Errorf("data retrieved was not 'DenyTimeoffResponse'")
}

type DenyTimeoffResponse struct {
	ID             string  `json:"id,omitempty"`
	StartDate      *Time   `json:"start_date,omitempty"`
	EndDate        *Time   `json:"end_date,omitempty"`
	EmployeeID     string  `json:"employee_id,omitempty"`
	Employee       string  `json:"employee,omitempty"`
	RequestedByID  string  `json:"requested_by_id,omitempty"`
	RequestedBy    string  `json:"requested_by,omitempty"`
	TypeID         string  `json:"type_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	DayType        string  `json:"day_type,omitempty"`
	Duration       float64 `json:"duration,omitempty"`
	Status         string  `json:"status,omitempty"`
	CreatedAt      *Time   `json:"created_at,omitempty"`
	IsPaid         bool    `json:"is_paid,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	ApproverID     string  `json:"approver_id,omitempty"`
	Approver       string  `json:"approver,omitempty"`
	ApprovalRespAt *Time   `json:"approval_resp_at,omitempty"`
	CancelledAt    *Time   `json:"cancelled_at,omitempty"`
	UpdatedAt      *Time   `json:"updated_at,omitempty"`
	Comments       []struct {
		CommentID   string `json:"comment_id,omitempty"`
		Comment     string `json:"comment,omitempty"`
		CommenterID string `json:"commenter_id,omitempty"`
		Commenter   string `json:"commenter,omitempty"`
		CreatedAt   *Time  `json:"created_at,omitempty"`
	} `json:"comments,omitempty"`
}
