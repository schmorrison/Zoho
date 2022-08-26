package recruit

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
func (c *API) GetJobOpenings(
	params map[string]zoho.Parameter,
) (data JobOpeningsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetJobOpenings",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/v2/%s",
			c.ZohoTLD,
			JobOpeningsModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &JobOpeningsResponse{},
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
		return JobOpeningsResponse{}, fmt.Errorf("failed to retrieve JobOpenings: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*JobOpeningsResponse); ok {
		return *v, nil
	}

	return JobOpeningsResponse{}, fmt.Errorf("no 'JobOpeningsResponse' returned")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
func (c *API) GetJobOpeningsById(id string) (data JobOpeningsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetJobOpeningsById",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/v2/%s/%s",
			c.ZohoTLD,
			JobOpeningsModule,
			id,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &JobOpeningsResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return JobOpeningsResponse{}, fmt.Errorf("failed to retrieve JobOpening with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*JobOpeningsResponse); ok {
		return *v, nil
	}

	return JobOpeningsResponse{}, fmt.Errorf("no 'JobOpeningsResponse' returned")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/search-records.html
func (c *API) SearchJobOpenings(
	params map[string]zoho.Parameter,
) (data JobOpeningsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "SearchJobOpenings",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/v2/%s/search",
			c.ZohoTLD,
			JobOpeningsModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &JobOpeningsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"per_page": "200",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	// log.Printf("%+v\n", endpoint)

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return JobOpeningsResponse{}, fmt.Errorf(
			"failed to retrieve searched %s: %s",
			JobOpeningsModule,
			err.Error(),
		)
	}

	if v, ok := endpoint.ResponseData.(*JobOpeningsResponse); ok {
		return *v, nil
	}

	return JobOpeningsResponse{}, fmt.Errorf("no 'JobOpeningsResponse' returned")
}

type JobOpening struct {
	ClientName struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"Client_Name,omitempty"`
	Salary         interface{} `json:"Salary,omitempty"`
	CurrencySymbol string      `json:"$currency_symbol,omitempty"`
	AccountManager struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"Account_Manager,omitempty"`
	NoOfCandidatesHired int       `json:"No_of_Candidates_Hired,omitempty"`
	TargetDate          string    `json:"Target_Date,omitempty"`
	LastActivityTime    time.Time `json:"Last_Activity_Time,omitempty"`
	Industry            string    `json:"Industry,omitempty"`
	ModifiedBy          struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"Modified_By,omitempty"`
	ProcessFlow     bool   `json:"$process_flow,omitempty"`
	ExpectedRevenue int    `json:"Expected_Revenue,omitempty"`
	IsHotJobOpening bool   `json:"Is_Hot_Job_Opening,omitempty"`
	ZipCode         string `json:"Zip_Code,omitempty"`
	ID              string `json:"id,omitempty"`
	Approved        bool   `json:"$approved,omitempty"`
	Publish         bool   `json:"Publish,omitempty"`
	DateOpened      string `json:"Date_Opened,omitempty"`
	Approval        struct {
		Delegate bool `json:"delegate,omitempty"`
		Approve  bool `json:"approve,omitempty"`
		Reject   bool `json:"reject,omitempty"`
		Resubmit bool `json:"resubmit,omitempty"`
	} `json:"$approval,omitempty"`
	ModifiedTime             time.Time   `json:"Modified_Time,omitempty"`
	ActualRevenue            interface{} `json:"Actual_Revenue,omitempty"`
	RemoteJob                bool        `json:"Remote_Job,omitempty"`
	CreatedTime              time.Time   `json:"Created_Time,omitempty"`
	Followed                 bool        `json:"$followed,omitempty"`
	NoOfCandidatesAssociated int         `json:"No_of_Candidates_Associated,omitempty"`
	Editable                 bool        `json:"$editable,omitempty"`
	IsLocked                 bool        `json:"Is_Locked,omitempty"`
	City                     string      `json:"City,omitempty"`
	JobOpeningStatus         string      `json:"Job_Opening_Status,omitempty"`
	RevenuePerPosition       int         `json:"Revenue_per_Position,omitempty"`
	ContactName              struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"Contact_Name,omitempty"`
	AssociatedTags []struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"Associated_Tags,omitempty"`
	AssignedRecruiter []struct {
		Name     string `json:"name,omitempty"`
		ID       string `json:"id,omitempty"`
		Email    string `json:"email,omitempty"`
		PhotoSrc string `json:"photoSrc,omitempty"`
	} `json:"Assigned_Recruiter,omitempty"`
	MissedRevenue     interface{} `json:"Missed_Revenue,omitempty"`
	JobOpeningID      string      `json:"Job_Opening_ID,omitempty"`
	JobDescription    string      `json:"Job_Description,omitempty"`
	WorkExperience    string      `json:"Work_Experience,omitempty"`
	JobType           string      `json:"Job_Type,omitempty"`
	JobOpeningName    string      `json:"Job_Opening_Name,omitempty"`
	NumberOfPositions string      `json:"Number_of_Positions,omitempty"`
	State             string      `json:"State,omitempty"`
	Country           string      `json:"Country,omitempty"`
	CreatedBy         struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"Created_By,omitempty"`
	IsAttachmentPresent bool `json:"Is_Attachment_Present,omitempty"`
}

type JobOpeningsResponse struct {
	Data []JobOpening `json:"data,omitempty"`
	Info PageInfo     `json:"info,omitempty"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-associated-records.html
func (c *API) GetAssociatedCandidates(
	recordId string,
) (data AssociatedCandidatesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "GetAssociatedCandidates",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/v2/%s/%s/associate",
			c.ZohoTLD,
			Job_OpeningsModule,
			recordId,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &AssociatedCandidatesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AssociatedCandidatesResponse{}, fmt.Errorf(
			"failed to get associated candidates of %s: %s",
			JobOpeningsModule,
			err.Error(),
		)
	}

	if v, ok := endpoint.ResponseData.(*AssociatedCandidatesResponse); ok {
		return *v, nil
	}

	return AssociatedCandidatesResponse{}, fmt.Errorf("no 'AssociatedCandidatesResponse' returned")
}

type AssociatedCandidatesResponse struct {
	Data []struct {
		Origin                   string      `json:"Origin"`
		Email                    string      `json:"Email"`
		CurrencySymbol           string      `json:"$currency_symbol"`
		LastActivityTime         time.Time   `json:"Last_Activity_Time"`
		HighestQualificationHeld interface{} `json:"Highest_Qualification_Held"`
		SkillSet                 interface{} `json:"Skill_Set"`
		Converted                bool        `json:"$converted"`
		ProcessFlow              bool        `json:"$process_flow"`
		UpdatedOn                time.Time   `json:"Updated_On"`
		CurrentEmployer          string      `json:"Current_Employer"`
		Street                   interface{} `json:"Street"`
		ZipCode                  interface{} `json:"Zip_Code"`
		ID                       string      `json:"id"`
		ExperienceInYears        int         `json:"Experience_in_Years"`
		Approved                 bool        `json:"$approved"`
		Approval                 struct {
			Delegate bool `json:"delegate"`
			Approve  bool `json:"approve"`
			Reject   bool `json:"reject"`
			Resubmit bool `json:"resubmit"`
		} `json:"$approval"`
		CandidateStatus string      `json:"Candidate_Status"`
		CandidateID     string      `json:"Candidate_ID"`
		LastMailedTime  interface{} `json:"Last_Mailed_Time"`
		CreatedTime     time.Time   `json:"Created_Time"`
		Followed        bool        `json:"$followed"`
		CandidateOwner  struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Candidate_Owner"`
		Editable       bool          `json:"$editable"`
		IsLocked       bool          `json:"Is_Locked"`
		City           string        `json:"City"`
		IsUnqualified  bool          `json:"Is_Unqualified"`
		AssociatedTags []interface{} `json:"Associated_Tags"`
		AdditionalInfo interface{}   `json:"Additional_Info"`
		State          interface{}   `json:"State"`
		Country        interface{}   `json:"Country"`
		CreatedBy      struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Created_By"`
		SecondaryEmail      interface{} `json:"Secondary_Email"`
		IsAttachmentPresent bool        `json:"Is_Attachment_Present"`
		Rating              interface{} `json:"Rating"`
		AppliedWithLinkedin interface{} `json:"$applied_with_linkedin"`
		Website             interface{} `json:"Website"`
		Twitter             interface{} `json:"Twitter"`
		CurrentJobTitle     string      `json:"Current_Job_Title"`
		Salutation          interface{} `json:"Salutation"`
		Source              string      `json:"Source"`
		FirstName           string      `json:"First_Name"`
		FullName            string      `json:"Full_Name"`
		ModifiedBy          struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Modified_By"`
		SkypeID         interface{} `json:"Skype_ID"`
		Phone           interface{} `json:"Phone"`
		EmailOptOut     bool        `json:"Email_Opt_Out"`
		ConvertedDetail struct {
		} `json:"$converted_detail"`
		CareerPageInviteStatus      string      `json:"Career_Page_Invite_Status"`
		Mobile                      string      `json:"Mobile"`
		LastName                    string      `json:"Last_Name"`
		CurrentSalary               interface{} `json:"Current_Salary"`
		AssociatedAnySocialProfiles interface{} `json:"Associated_any_Social_Profiles"`
		Fax                         interface{} `json:"Fax"`
		ExpectedSalary              interface{} `json:"Expected_Salary"`
	} `json:"data"`
	Info PageInfo `json:"info,omitempty"`
}

// ============================== ABANDONED DUE TO ZOHO API v1 not properly working ==============================
// https://help.zoho.com/portal/en/kb/recruit/developer-guide/api-methods/articles/getsearchrecords
func (c *API) XMLSearchJobOpenings(
	params map[string]zoho.Parameter,
) (data JobOpeningsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "XMLSearchJobOpenings",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/private/xml/%s/getSearchRecords",
			c.ZohoTLD,
			JobOpeningsModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &XMLJobOpeningsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"fromIndex":       "1",     // Integer | Default value - 1
			"toIndex":         "200",   // Integer | Default value - 20 | Maximum value - 200
			"version":         "2",     // This will fetch responses based on the latest API implementation.
			"newFormat":       "1",     // 1 - To exclude fields with "null" values while inserting data from your Recruit account.
			"selectColumns":   "(ALL)", // mandatory eg: (Job Description) | Module(Job Description)
			"searchCondition": "",      // mandatory eg: (Associated Tags|=|JobIsPublished)
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	// log.Printf("%s\n", endpoint)

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return JobOpeningsResponse{}, fmt.Errorf(
			"failed to retrieve XML searched %s: %s",
			JobOpeningsModule,
			err.Error(),
		)
	}

	if v, ok := endpoint.ResponseData.(*XMLJobOpeningsResponse); ok {
		// Convert to Json
		total, _ := strconv.Atoi(v.Result.JobOpenings.Row.No)
		dataXML := v.Result.JobOpenings.Row.FL

		res := &JobOpeningsResponse{}
		items := make([]JobOpening, 0)

		for i := 0; i <= total; i++ {
			switch dataXML[i].Val {
			case "JOBOPENINGID":
				items = append(items, JobOpening{ID: dataXML[i].Text})
			case "CLIENTID":
				items = append(items, JobOpening{})
			case "Client Name":
				items = append(items, JobOpening{})
			case "Posting Title":
				items = append(items, JobOpening{JobOpeningName: dataXML[i].Text})
			case "Job Opening ID":
				items = append(items, JobOpening{JobOpeningID: dataXML[i].Text})
			case "Job Description":
				items = append(items, JobOpening{JobDescription: dataXML[i].Text})
			}
		}

		res.Data = items

		return *res, nil
	}

	return JobOpeningsResponse{}, fmt.Errorf("no 'JobOpeningsResponse' returned")
}

// https://help.zoho.com/portal/en/kb/recruit/developer-guide/api-methods/articles/getrecordbyid#Purpose
func (c *API) XMLgetRecordById(params map[string]zoho.Parameter) (data JobOpening, err error) {
	endpoint := zoho.Endpoint{
		Name: "XMLgetRecordById",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/private/xml/%s/getRecordById",
			c.ZohoTLD,
			JobOpeningsModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &XMLJobOpeningsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"id":        "",  // mandatory
			"version":   "2", // This will fetch responses based on the latest API implementation.
			"newFormat": "1", // 1 - To exclude fields with "null" values while inserting data from your Recruit account.
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	// log.Printf("%s\n", endpoint)

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return JobOpening{}, fmt.Errorf(
			"failed to retrieve XML searched %s: %s",
			JobOpeningsModule,
			err.Error(),
		)
	}

	if v, ok := endpoint.ResponseData.(*XMLJobOpeningsResponse); ok {
		// Convert to Json
		xmlData := v.Result.JobOpenings.Row.FL
		res := &data

		for _, item := range xmlData {
			switch item.Val {
			case "Job Description":
				data.JobDescription = item.Text
			}
		}
		return *res, nil
	}

	return JobOpening{}, fmt.Errorf("no 'JobOpening' returned")
}

type XMLJobOpeningsResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	URI     string   `xml:"uri,attr"`
	Result  struct {
		Text        string `xml:",chardata"`
		JobOpenings struct {
			Text string `xml:",chardata"`
			Row  struct {
				Text string `xml:",chardata"`
				No   string `xml:"no,attr"`
				FL   []struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"FL"`
			} `xml:"row"`
		} `xml:"JobOpenings"`
	} `xml:"result"`
}

// https://help.zoho.com/portal/en/kb/recruit/developer-guide/api-methods/articles/getrecords#Request_Parameters
func (c *API) XMLGetRecords(
	params map[string]zoho.Parameter,
) (data XMLGetRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name: "XMLGetRecords",
		URL: fmt.Sprintf(
			"https://recruit.zoho.%s/recruit/private/xml/%s/getRecordById",
			c.ZohoTLD,
			JobOpeningsModule,
		),
		Method:       zoho.HTTPGet,
		ResponseData: &XMLGetRecordsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"version":       "2", // This will fetch responses based on the latest API implementation.
			"id":            "",  // 1 - To exclude fields with "null" values while inserting data from your Recruit account.
			"selectColumns": "",  // mandatory eg: (Job Description) | Module(Job Description)
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	// log.Printf("ENDPOINT: %s\n", endpoint)

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return XMLGetRecordsResponse{}, fmt.Errorf(
			"failed to retrieve XML get %s: %s",
			JobOpeningsModule,
			err.Error(),
		)
	}

	if v, ok := endpoint.ResponseData.(*XMLGetRecordsResponse); ok {
		return *v, nil
	}

	return XMLGetRecordsResponse{}, fmt.Errorf("no 'XMLGetRecordsResponse' returned")
}

type XMLGetRecordsResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	URI     string   `xml:"uri,attr"`
	Result  struct {
		Text        string `xml:",chardata"`
		JobOpenings struct {
			Text string `xml:",chardata"`
			Row  struct {
				Text string `xml:",chardata"`
				No   string `xml:"no,attr"`
				FL   []struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"FL"`
			} `xml:"row"`
		} `xml:"JobOpenings"`
	} `xml:"result"`
}
