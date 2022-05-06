package recruit

import (
	"fmt"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// GetInterviewsRecords returns a list of all records
// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
// https://recruit.zoho.eu/recruit/v2/Interviews
func (c *API) GetInterviewsRecords(params map[string]zoho.Parameter) (data InterviewsRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetInterviewsRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, InterviewsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &InterviewsRecordsResponse{},
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
		return InterviewsRecordsResponse{}, fmt.Errorf("failed to retrieve Interviews: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*InterviewsRecordsResponse); ok {
		return *v, nil
	}

	return InterviewsRecordsResponse{}, fmt.Errorf("data returned was not 'InterviewsRecordsResponse'")
}

// GetInterviewsRecord returns the record specified by ID
// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
// https://recruit.zoho.eu/recruit/v2/Interviews/{id}
func (c *API) GetInterviewsRecordById(id string) (data InterviewsRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetInterviewsRecordById",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s", c.ZohoTLD, InterviewsModule, id),
		Method:       zoho.HTTPGet,
		ResponseData: &InterviewsRecordsResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InterviewsRecordsResponse{}, fmt.Errorf("failed to retrieve JobOpening with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*InterviewsRecordsResponse); ok {
		return *v, nil
	}

	return InterviewsRecordsResponse{}, fmt.Errorf("data returned was not 'InterviewsRecordsResponse'")
}

// InterviewsRecordsResponse is the data returned by GetInterviewsRecords & GetInterviewsRecordById
type InterviewsRecordsResponse struct {
	Data []struct {
		ClientName struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Client_Name,omitempty"`
		CurrencySymbol           string    `json:"$currency_symbol,omitempty"`
		VideoInterviewIsReviewed bool      `json:"$video_interview_isreviewed,omitempty"`
		StartDateTime            time.Time `json:"Start_DateTime,omitempty"`
		InterviewDuration        struct {
			Mins int `json:"mins,omitempty"`
			Hrs  int `json:"hrs,omitempty"`
			Days int `json:"days,omitempty"`
		} `json:"$interview_duration,omitempty"`
		InterviewEvaluationDone   bool `json:"$interview_evaluationdone,omitempty"`
		EndDateTime               Time `json:"End_DateTime,omitempty"`
		LastActivityTime          Time `json:"Last_Activity_Time,omitempty"`
		InterviewEnableEvaluation bool `json:"$interview_enableevaluation,omitempty"`
		ModifiedBy                struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		CandidateName struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Candidate_Name,omitempty"`
		ProcessFlow    bool `json:"$process_flow,omitempty"`
		InterviewOwner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Interview_Owner,omitempty"`
		Feedback         string `json:"Feedback,omitempty"`
		ScheduleComments string `json:"Schedule_Comments,omitempty"`
		InterviewType    int    `json:"$interview_type,omitempty"`
		ReviewedBy       string `json:"Reviewed_By,omitempty"`
		ID               string `json:"id,omitempty"`
		Approved         bool   `json:"$approved,omitempty"`
		Approval         struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		IsStatusSplitDone bool       `json:"isStatusSplitDone,omitempty"`
		ModifiedTime      Time       `json:"Modified_Time,omitempty"`
		Venue             string     `json:"Venue,omitempty"`
		CreatedTime       Time       `json:"Created_Time,omitempty"`
		Followed          bool       `json:"$followed,omitempty"`
		Taxable           bool       `json:"$taxable,omitempty"`
		Editable          bool       `json:"$editable,omitempty"`
		IsLocked          bool       `json:"Is_Locked,omitempty"`
		InterviewName     string     `json:"Interview_Name,omitempty"`
		QuestionnaireName string     `json:"Questionnaire_Name,omitempty"`
		AssociatedTags    []struct{} `json:"Associated_Tags,omitempty"`
		InterviewStatus   string     `json:"Interview_Status,omitempty"`
		JobOpeningName    struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Job_Opening_Name,omitempty"`
		ReviewedTime string `json:"Reviewed_Time,omitempty"`
		Interviewer  []struct {
			Name     string `json:"name,omitempty"`
			ID       string `json:"id,omitempty"`
			Email    string `json:"email,omitempty"`
			PhotoSrc string `json:"photoSrc,omitempty"`
		} `json:"Interviewer,omitempty"`
		Created_By struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		IsAttachmentPresent       bool     `json:"Is_Attachment_Present,omitempty"`
		InterviewTimeTillReview   struct{} `json:"$interview_timetillreview,omitempty"`
		VideoInterviewIsSubmitted bool     `json:"video_interview_issubmitted,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}
