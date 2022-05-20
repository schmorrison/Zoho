package recruit

import (
	"fmt"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// https://www.zoho.com/recruit/developer-guide/apiv2/search-records.html
func (c *API) SearchRecords(request interface{}, module Module, params map[string]zoho.Parameter) (data interface{}, err error) {
	endpoint := zoho.Endpoint{
		Name:         "SearchRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/search", c.ZohoTLD, module),
		Method:       zoho.HTTPGet,
		ResponseData: request,
		URLParameters: map[string]zoho.Parameter{
			"criteria":  "",    // optional
			"email":     "",    // optional
			"phone":     "",    // optional
			"word":      "",    // optional
			"converted": "",    // optional
			"approved":  "",    // optional
			"page":      "1",   // optional
			"per_page":  "200", // optional
		},
	}

	for k, v := range params {
		endpoint.URLParameters[k] = v
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records of %s: %s", module, err)
	}

	if endpoint.ResponseData != nil {
		return endpoint.ResponseData, nil
	}

	return nil, fmt.Errorf("data returned was nil")
}

// https://www.zoho.com/recruit/developer-guide/apiv2/insert-records.html
func (c *API) InsertRecords(request InsertRecords, module Module) (data InsertRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "InsertRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &InsertRecordsResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InsertRecordsResponse{}, fmt.Errorf("failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*InsertRecordsResponse); ok {
		return *v, nil
	}

	return InsertRecordsResponse{}, fmt.Errorf("data returned was nil")
}

type UpdateRecordsResponseData struct {
	Message string `json:"message,omitempty"`
	Details struct {
		ExpectedDataType string `json:"expected_data_type,omitempty"`
		APIName          string `json:"api_name"`
	} `json:"details,omitempty"`
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
}

type InsertRecords struct {
	Data    interface{} `json:"data,omitempty"`
	Trigger []string    `json:"trigger,omitempty"`
}
type InsertRecordsResponseData struct {
	Message string `json:"message,omitempty"`
	Details struct {
		CreatedBy struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"created_by,omitempty"`
		ID         string `json:"id,omitempty"`
		ModifiedBy struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"modified_by,omitempty"`
		ModifiedTime time.Time `json:"modified_time,omitempty"`
		CreatedTime  time.Time `json:"created_time,omitempty"`
	} `json:"details,omitempty"`
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
}

type InsertRecordsResponse struct {
	Data []InsertRecordsResponseData `json:"data,omitempty"`
}

// UpsertRecords will insert the provided records in the request, if they already exist it will be updated
// https://www.zoho.com/recruit/developer-guide/apiv2/upsert-records.html
//
// When performing an upsert, because the natural state of the records fields in this package is to 'omitempty' when encoding json,
// if you want to empty the fields contents in zoho you will need to embed the records type in a struct in your own package,
// and override the field with a field that has a json tag that does not contain 'omitempty'.
// eg.
//    type struct Candidate {
//        zohorecruit.Candidate
//        CustomField string `json:"Custom_Field"`
//     }
func (c *API) UpsertRecords(request UpsertRecords, module Module) (data InsertRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "UpsertRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/upsert", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &InsertRecordsResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InsertRecordsResponse{}, fmt.Errorf("failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*InsertRecordsResponse); ok {
		return *v, nil
	}

	return InsertRecordsResponse{}, fmt.Errorf("data returned was nil")
}

type UpsertRecords struct {
	Data                 interface{} `json:"data,omitempty"`
	DuplicateCheckFields []string    `json:"duplicate_check_fields,omitempty"`
	Trigger              []string    `json:"trigger,omitempty"`
}

// https://www.zoho.com/recruit/developer-guide/apiv2/get-associated-records.html
func (c *API) GetAssociatedRecords(module Module, recordId string) (data AssociateRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAssociatedRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s/associate", c.ZohoTLD, module, recordId),
		Method:       zoho.HTTPGet,
		ResponseData: &AssociateRecordsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"posting_title":      "",    // optional
			"candidate_statuses": "",    // optional
			"page":               "1",   // optional
			"per_page":           "200", // optional
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AssociateRecordsResponse{}, fmt.Errorf("failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*AssociateRecordsResponse); ok {
		return *v, nil
	}

	return AssociateRecordsResponse{}, fmt.Errorf("data returned was nil")
}

type AssociateRecordsResponse struct {
	Data []struct {
		Salary         interface{} `json:"Salary"`
		CurrencySymbol string      `json:"$currency_symbol"`
		AccountManager struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Account_Manager"`
		NoOfCandidatesHired int       `json:"No_of_Candidates_Hired"`
		TargetDate          string    `json:"Target_Date"`
		LastActivityTime    time.Time `json:"Last_Activity_Time"`
		Industry            string    `json:"Industry"`
		ModifiedBy          struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Modified_By"`
		ProcessFlow     bool        `json:"$process_flow"`
		ExpectedRevenue interface{} `json:"Expected_Revenue"`
		IsHotJobOpening bool        `json:"Is_Hot_Job_Opening"`
		AccountName     struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Account_Name"`
		ZipCode       interface{} `json:"Zip_Code"`
		ID            string      `json:"id"`
		Skillset      string      `json:"Skillset"`
		Approved      bool        `json:"$approved"`
		Publish       bool        `json:"Publish"`
		DateOpened    string      `json:"Date_Opened"`
		PotentialName string      `json:"Potential_Name"`
		Approval      struct {
			Delegate bool `json:"delegate"`
			Approve  bool `json:"approve"`
			Reject   bool `json:"reject"`
			Resubmit bool `json:"resubmit"`
		} `json:"$approval"`
		ModifiedTime             time.Time     `json:"Modified_Time"`
		ActualRevenue            interface{}   `json:"Actual_Revenue"`
		CreatedTime              time.Time     `json:"Created_Time"`
		Followed                 bool          `json:"$followed"`
		NoOfCandidatesAssociated int           `json:"No_of_Candidates_Associated"`
		Editable                 bool          `json:"$editable"`
		IsLocked                 bool          `json:"Is_Locked"`
		City                     interface{}   `json:"City"`
		JobOpeningStatus         string        `json:"Job_Opening_Status"`
		RevenuePerPosition       int           `json:"Revenue_per_Position"`
		ContactName              interface{}   `json:"Contact_Name"`
		AssociatedTags           []interface{} `json:"Associated_Tags"`
		AssignedRecruiter        []interface{} `json:"Assigned_Recruiter"`
		MissedRevenue            interface{}   `json:"Missed_Revenue"`
		JobOpeningID             string        `json:"Job_Opening_ID"`
		JobDescription           string        `json:"Job_Description"`
		WorkExperience           interface{}   `json:"Work_Experience"`
		JobType                  string        `json:"Job_Type"`
		State                    interface{}   `json:"State"`
		NumberOfPositions        string        `json:"Number_of_Positions"`
		Country                  interface{}   `json:"Country"`
		CreatedBy                struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"Created_By"`
		IsAttachmentPresent bool `json:"Is_Attachment_Present"`
	} `json:"data"`
	Info PageInfo `json:"info,omitempty"`
}
