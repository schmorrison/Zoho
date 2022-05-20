package recruit

import (
	"fmt"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.eu/recruit/developer-guide/apiv2/insert-records.html
func (c *API) InsertCandidates(request InsertCandidateRequest) (data InsertCandidateResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "InsertCandidates",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, CandidatesModule),
		Method:       zoho.HTTPPost,
		ResponseData: &InsertCandidateResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InsertCandidateResponse{}, fmt.Errorf("failed to insert Candidate(s): %s", err.Error())
	}

	if v, ok := endpoint.ResponseData.(*InsertCandidateResponse); ok {
		for _, resp := range v.Data {
			if resp.Code != "SUCCESS" {
				return InsertCandidateResponse{}, fmt.Errorf("failed to insert Candidate(s): %s: %s", resp.Code, resp.Details)
			}
		}
		return *v, nil
	}

	return InsertCandidateResponse{}, fmt.Errorf("no 'InsertCandidateResponse' returned")
}

type CandidateRequestData struct {
	CandidateId       string `json:"CandidateId,omitempty"`
	ApplicationDate   string `json:"ApplicationDate,omitempty"`
	FirstName         string `json:"First_Name,omitempty"`
	LastName          string `json:"Last_Name,omitempty"`
	City              string `json:"City,omitempty"`
	Email             string `json:"Email,omitempty"`
	Mobile            string `json:"Mobile,omitempty"`
	Street            string `json:"Street,omitempty"`
	CurrentEmployer   string `json:"Current_Employer,omitempty"`
	CurrentJobTitle   string `json:"Current_Job_Title,omitempty"`
	ExpectedSalary    string `json:"Expected_Salary,omitempty"`
	CurrencySymbol    string `json:"$currency_symbol,omitempty"`
	ExperienceInYears string `json:"Experience_in_Years,omitempty"`
	CandidateStatus   string `json:"Candidate_Status,omitempty"`
	Source            string `json:"Source,omitempty"`
}

type InsertCandidateRequest struct {
	Data    []CandidateRequestData `json:"data"`
	Trigger []string               `json:"trigger"`
}

type InsertCandidateResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime time.Time `json:"Created_Time"`
			ID          string    `json:"id"`
			CreatedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

// Upsert = Insert or Update
// https://www.zoho.com/recruit/developer-guide/apiv2/upsert-records.html
func (c *API) UpsertCandidates(request UpsertCandidateRequest) (data UpsertCandidateResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpsertCandidates",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/upsert", c.ZohoTLD, CandidatesModule),
		Method:       zoho.HTTPPost,
		ResponseData: &UpsertCandidateResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpsertCandidateResponse{}, fmt.Errorf("failed to upsert Candidate(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpsertCandidateResponse); ok {
		return *v, nil
	}

	return UpsertCandidateResponse{}, fmt.Errorf("no 'UpsertCandidateResponse' returned")
}

type UpsertCandidateRequest struct {
	Data                 []CandidateRequestData `json:"data"`
	DuplicateCheckFields []string               `json:"duplicate_check_fields"`
	Trigger              []string               `json:"trigger"`
}
type UpsertCandidateResponse struct {
	Data []struct {
		Code           string `json:"code"`
		DuplicateField string `json:"duplicate_field"`
		Action         string `json:"action"`
		Details        struct {
			ModifiedTime time.Time `json:"Modified_Time"`
			ModifiedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Modified_By"`
			CreatedTime time.Time `json:"Created_Time"`
			ID          string    `json:"id"`
			CreatedBy   struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"Created_By"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
func (c *API) GetCandidates(params map[string]zoho.Parameter) (data CandidatesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetCandidates",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, CandidatesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &CandidatesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"fields":        "",
			"sort_order":    "",
			"sort_by":       "",
			"converted":     "false",
			"approved":      "true",
			"page":          "1",
			"per_page":      "200",
			"cvid":          "",
			"territory_id":  "",
			"include_child": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CandidatesResponse{}, fmt.Errorf("failed to retrieve Candidates: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CandidatesResponse); ok {
		return *v, nil
	}

	return CandidatesResponse{}, fmt.Errorf("no 'CandidatesResponse' returned")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
func (c *API) GetCandidateById(id string) (data CandidatesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetCandidateById",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s", c.ZohoTLD, CandidatesModule, id),
		Method:       zoho.HTTPGet,
		ResponseData: &CandidatesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CandidatesResponse{}, fmt.Errorf("failed to retrieve Candidate with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CandidatesResponse); ok {
		return *v, nil
	}

	return CandidatesResponse{}, fmt.Errorf("no 'CandidatesResponse' returned")
}

type CandidatesResponse struct {
	Data []struct {
		Origin                   string `json:"Origin,omitempty"`
		Email                    string `json:"Email,omitempty"`
		CurrencySymbol           string `json:"$currency_symbol,omitempty"`
		Last_Activity_Time       Time   `json:"Last_Activity_Time,omitempty"`
		HighestQualificationHeld string `json:"Highest_Qualification_Held,omitempty"`
		SkillSet                 string `json:"Skill_Set,omitempty"`
		Converted                bool   `json:"$converted,omitempty"`
		ProcessFlow              bool   `json:"$process_flow,omitempty"`
		Updated_On               Time   `json:"Updated_On,omitempty"`
		CurrentEmployer          string `json:"Current_Employer,omitempty"`
		Street                   string `json:"Street,omitempty"`
		ZipCode                  string `json:"Zip_Code,omitempty"`
		ID                       string `json:"id,omitempty"`
		ExperienceInYears        int    `json:"Experience_in_Years,omitempty"`
		Approved                 bool   `json:"$approved,omitempty"`
		Approval                 struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		CandidateStatus string `json:"Candidate_Status,omitempty"`
		CandidateID     string `json:"Candidate_ID,omitempty"`
		LastMailedTime  Time   `json:"Last_Mailed_Time,omitempty"`
		CreatedTime     string `json:"Created_Time,omitempty"`
		Followed        string `json:"followed,omitempty"`
		CandidateOwner  struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Candidate_Owner,omitempty"`
		Editable       bool       `json:"$editable,omitempty"`
		IsLocked       bool       `json:"Is_Locked,omitempty"`
		City           string     `json:"City,omitempty"`
		IsUnqualified  bool       `json:"Is_Unqualified,omitempty"`
		AssociatedTags []struct{} `json:"Associated_Tags,omitempty"`
		AdditionalInfo string     `json:"Additional_Info,omitempty"`
		State          string     `json:"State,omitempty"`
		Country        string     `json:"Country,omitempty"`
		CreatedBy      struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		SecondaryEmail      string `json:"Secondary_Email,omitempty"`
		IsAttachmentPresent bool   `json:"Is_Attachment_Present,omitempty"`
		Rating              int    `json:"Rating,omitempty"`
		AppliedWithLinkedin string `json:"$applied_with_linkedin,omitempty"`
		Website             string `json:"Website,omitempty"`
		Twitter             string `json:"Twitter,omitempty"`
		CurrentJobTitle     string `json:"Current_Job_Title,omitempty"`
		Salutation          string `json:"Salutation,omitempty"`
		Source              string `json:"Source,omitempty"`
		FirstName           string `json:"First_Name,omitempty"`
		FullName            string `json:"Full_Name,omitempty"`
		ModifiedBy          struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		SkypeID                     string   `json:"Skype_ID,omitempty"`
		Phone                       string   `json:"Phone,omitempty"`
		Email_Opt_Out               bool     `json:"Email_Opt_Out,omitempty"`
		IsStatusSplitDone           bool     `json:"isStatusSplitDone,omitempty"`
		ConvertedDetail             struct{} `json:"#converted_detail,omitempty"`
		CareerPageInviteStatus      string   `json:"Career_Page_Invite_Status,omitempty"`
		Mobile                      string   `json:"Mobile,omitempty"`
		LastName                    string   `json:"Last_Name,omitempty"`
		CurrentSalary               string   `json:"Current_Salary,omitempty"`
		AssociatedAnySocialProfiles bool     `json:"Associated_any_Social_Profiles,omitempty"`
		Fax                         string   `json:"Fax,omitempty"`
		ExpectedSalary              string   `json:"Expected_Salary,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-related-records.html
func (c *API) GetCandidateRelatedRecords(params map[string]zoho.Parameter, candidateId string, record RelatedRecord) (data CandidateRelatedRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetCandidateRelatedRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s/%s", c.ZohoTLD, CandidatesModule, candidateId, record),
		Method:       zoho.HTTPGet,
		ResponseData: &CandidateRelatedRecordsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"ids":        "",
			"sort_order": "",
			"sort_by":    "",
			"page":       "1",
			"per_page":   "200",
			"fields":     "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CandidateRelatedRecordsResponse{}, fmt.Errorf("failed to retrieve Candidates: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CandidateRelatedRecordsResponse); ok {
		return *v, nil
	}

	return CandidateRelatedRecordsResponse{}, fmt.Errorf("no 'CandidateRelatedRecordsResponse' returned")
}

type CandidateRelatedRecordsResponse struct {
	Data []struct {
		AttachType struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"$attach_type"`
		ModifiedTime    time.Time `json:"Modified_Time"`
		FileName        string    `json:"File_Name"`
		Size            string    `json:"Size"`
		CreatedTime     time.Time `json:"Created_Time"`
		LinkDocs        int       `json:"$link_docs"`
		ParentID        string    `json:"Parent_Id"`
		Editable        bool      `json:"$editable"`
		AttachmentOwner struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Attachment_Owner"`
		FileID     string `json:"$file_id"`
		Type       string `json:"$type"`
		SeModule   string `json:"$se_module"`
		ModifiedBy struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Modified_By"`
		ID        string `json:"id"`
		CreatedBy struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Created_By"`
		LinkURL interface{} `json:"$link_url"`
	} `json:"data"`
	Info PageInfo `json:"info,omitempty"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/delete-records.html
func (c *API) DeleteCandidateById(ID string) (data DeleteCandidateResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "DeleteCandidateById",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s", c.ZohoTLD, CandidatesModule, ID),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteCandidateResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteCandidateResponse{}, fmt.Errorf("failed to delete Candidate: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteCandidateResponse); ok {
		return *v, nil
	}

	return DeleteCandidateResponse{}, fmt.Errorf("no 'DeleteCandidateResponse' returned")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/delete-records.html
func (c *API) DeleteCandidatesByIds(IDs ...string) (data DeleteCandidateResponse, err error) {
	if len(IDs) == 0 {
		return DeleteCandidateResponse{}, fmt.Errorf("failed to delete Candidates, must provide at least 1 ID")
	}

	idStr := ""
	for i, a := range IDs {
		idStr += a
		if i < len(IDs)-1 {
			idStr += ","
		}
	}

	endpoint := zoho.Endpoint{
		Name:         "DeleteCandidatesByIds",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, CandidatesModule),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteCandidateResponse{},
		URLParameters: map[string]zoho.Parameter{
			"ids": zoho.Parameter(idStr),
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteCandidateResponse{}, fmt.Errorf("failed to delete Candidate(s): %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteCandidateResponse); ok {
		return *v, nil
	}

	return DeleteCandidateResponse{}, fmt.Errorf("no 'DeleteCandidateResponse' returned")
}

type DeleteCandidateResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			ID string `json:"id"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-deleted-records.html
func (c *API) ListDeletedCandidates(params map[string]zoho.Parameter) (data DeletedCandidatesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "ListDeletedCandidates",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/deleted", c.ZohoTLD, CandidatesModule),
		Method:       zoho.HTTPGet,
		ResponseData: &DeletedCandidatesResponse{},
		URLParameters: map[string]zoho.Parameter{
			"type":     "",
			"page":     "1",
			"per_page": "200",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeletedCandidatesResponse{}, fmt.Errorf("failed to retrieve Deleted Candidates: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*DeletedCandidatesResponse); ok {
		return *v, nil
	}

	return DeletedCandidatesResponse{}, fmt.Errorf("no 'DeletedCandidatesResponse' returned")
}

type DeletedCandidatesResponse struct {
	Data []struct {
		DeletedBy struct {
			Name string `json:"id,omitempty"`
			ID   string `json:"id2,omitempty"`
		} `json:"deleted_by,omitempty"`
		ID          string `json:"id,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
		Type        string `json:"type,omitempty"`
		CreatedBy   struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"created_by,omitempty"`
		DeletedTime Time `json:"deleted_time,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/associate-candidate.html
func (c *API) AssociateCandidates(request AssociateCandidatesRequest) (data AssociateCandidatesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "AssociateCandidates",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/actions/associate", c.ZohoTLD, CandidatesModule),
		Method:       zoho.HTTPPut,
		ResponseData: &AssociateCandidatesResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AssociateCandidatesResponse{}, fmt.Errorf("failed to associate candidate: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*AssociateCandidatesResponse); ok {
		return *v, nil
	}

	return AssociateCandidatesResponse{}, fmt.Errorf("no 'AssociateCandidatesResponse' returned")
}

type AssociateCandidatesResponseData struct {
	Jobids   []string `json:"jobids"`
	Ids      []string `json:"ids"`
	Comments string   `json:"comments"`
}
type AssociateCandidatesRequest struct {
	Data []AssociateCandidatesResponseData `json:"data"`
}
type AssociateCandidatesResponse struct {
	Data []struct {
		Code    string `json:"code"`
		Details struct {
			Jobid string `json:"jobid"`
			Ids   string `json:"ids"`
			Error []struct {
				Code    string `json:"code"`
				Details struct {
					ID string `json:"id"`
				} `json:"details"`
				Message string `json:"message"`
				Status  string `json:"status"`
			} `json:"error"`
		} `json:"details"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"data"`
}
