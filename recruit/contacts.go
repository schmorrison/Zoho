package recruit

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetContactsRecords returns a list of all records
// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
// https://recruit.zoho.eu/recruit/v2/Contacts
func (c *API) GetContactsRecords(params map[string]zoho.Parameter) (data ContactsRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetContactsRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, ContactsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ContactsRecordsResponse{},
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
		return ContactsRecordsResponse{}, fmt.Errorf("failed to retrieve Contacts: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ContactsRecordsResponse); ok {
		return *v, nil
	}

	return ContactsRecordsResponse{}, fmt.Errorf("data returned was not 'ContactsRecordsResponse'")
}

// GetContactsRecord returns the record specified by ID
// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
// https://recruit.zoho.eu/recruit/v2/Contacts/{id}
func (c *API) GetContactsRecordById(id string) (data ContactsRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetContactsRecordById",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s", c.ZohoTLD, ContactsModule, id),
		Method:       zoho.HTTPGet,
		ResponseData: &ContactsRecordsResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ContactsRecordsResponse{}, fmt.Errorf("failed to retrieve JobOpening with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ContactsRecordsResponse); ok {
		return *v, nil
	}

	return ContactsRecordsResponse{}, fmt.Errorf("data returned was not 'ContactsRecordsResponse'")
}

// ContactsRecordsResponse is the data returned by GetContactsRecords & GetContactsRecordById
type ContactsRecordsResponse struct {
	Data []struct {
		Salary     string `json:"Salary,omitempty"`
		ClientName struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Client_Name,omitempty"`
		Email            string `json:"Email,omitempty"`
		CurrencySymbol   string `json:"$currency_symbol,omitempty"`
		MailingZip       string `json:"Mailing_Zip,omitempty"`
		IsPrimaryContact bool   `json:"Is_primary_contact,omitempty"`
		MailingState     string `json:"Mailing_State,omitempty"`
		Twitter          string `json:"Twitter,omitempty"`
		OtherZip         string `json:"Other_Zip,omitempty"`
		MailingStreet    string `json:"Mailing_Street,omitempty"`
		OtherState       string `json:"Other_State,omitempty"`
		Salutation       string `json:"Salutation,omitempty"`
		OtherCountry     string `json:"Other_Country,omitempty"`
		Source           string `json:"Source,omitempty"`
		LastActivityTime Time   `json:"Last_Activity_Time,omitempty"`
		FirstName        string `json:"First_Name,omitempty"`
		FullName         string `json:"Full_Name,omitempty"`
		Department       string `json:"Department,omitempty"`
		ModifiedBy       struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		SkypeID        string `json:"Skype_ID,omitempty"`
		ProcessFlow    bool   `json:"$process_flow,omitempty"`
		WorkPhone      string `json:"Work_Phone,omitempty"`
		MailingCountry string `json:"Mailing_Country,omitempty"`
		ID             string `json:"id,omitempty"`
		EmailOptOut    bool   `json:"Email_Opt_Out,omitempty"`
		Approved       bool   `json:"$approved,omitempty"`
		Approval       struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		IsStatusSplitDone      bool       `json:"isStatusSplitDone,omitempty"`
		ModifiedTime           Time       `json:"Modified_Time,omitempty"`
		MailingCity            string     `json:"Mailing_City,omitempty"`
		LastMailedTime         string     `json:"Last_Mailed_Time,omitempty"`
		OtherCity              string     `json:"Other_City,omitempty"`
		CreatedTime            Time       `json:"Created_Time,omitempty"`
		ClientPortalUserStatus string     `json:"Client_Portal_User_Status,omitempty"`
		Followed               bool       `json:"$followed,omitempty"`
		Editable               bool       `json:"$editable,omitempty"`
		OtherStreet            string     `json:"Other_Street,omitempty"`
		JobTitle               string     `json:"Job_Title,omitempty"`
		Mobile                 string     `json:"Mobile,omitempty"`
		AssociatedTags         []struct{} `json:"Associated_Tags,omitempty"`
		ContactOwner           struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Contact_Owner,omitempty"`
		LastName                    string `json:"Last_Name,omitempty"`
		AssociatedAnySocialProfiles bool   `json:"Associated_any_Social_Profiles,omitempty"`
		CreatedBy                   struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		Fax                 string `json:"Fax,omitempty"`
		SecondaryEmail      string `json:"Secondary_Email,omitempty"`
		IsAttachmentPresent bool   `json:"Is_Attachment_Present,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}
