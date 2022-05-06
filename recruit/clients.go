package recruit

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetClientsRecords returns a list of all records
// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
// https://recruit.zoho.eu/recruit/v2/Clients
func (c *API) GetClientsRecords(params map[string]zoho.Parameter) (data ClientsRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetClientsRecords",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s", c.ZohoTLD, ClientsModule),
		Method:       zoho.HTTPGet,
		ResponseData: &ClientsRecordsResponse{},
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
		return ClientsRecordsResponse{}, fmt.Errorf("failed to retrieve Clients: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ClientsRecordsResponse); ok {
		return *v, nil
	}

	return ClientsRecordsResponse{}, fmt.Errorf("data returned was not 'ClientsRecordsResponse'")
}

// GetClientsRecord returns the record specified by ID
// https://www.zoho.com/recruit/developer-guide/apiv2/get-records.html
// https://recruit.zoho.eu/recruit/v2/Clients/{id}
func (c *API) GetClientsRecordById(id string) (data ClientsRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetClientsRecordById",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/%s/%s", c.ZohoTLD, ClientsModule, id),
		Method:       zoho.HTTPGet,
		ResponseData: &ClientsRecordsResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ClientsRecordsResponse{}, fmt.Errorf("failed to retrieve JobOpening with id: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ClientsRecordsResponse); ok {
		return *v, nil
	}

	return ClientsRecordsResponse{}, fmt.Errorf("data returned was not 'ClientsRecordsResponse'")
}

// ClientsRecordsResponse is the data returned by GetClientsRecords & GetClientsRecordById
type ClientsRecordsResponse struct {
	Data []struct {
		ClientName     string `json:"Client_Name,omitempty"`
		CurrencySymbol string `json:"$currency_symbol,omitempty"`
		ShippingState  string `json:"Shipping_State,omitempty"`
		Website        string `json:"Website,omitempty"`
		AccountManager struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Account_Manager,omitempty"`
		Source           string `json:"Source,omitempty"`
		LastActivityTime Time   `json:"Last_Activity_Time,omitempty"`
		Industry         string `json:"Industry,omitempty"`
		ModifiedBy       struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		ProcessFlow    bool   `json:"$process_flow,omitempty"`
		BillingCountry string `json:"Billing_Country,omitempty"`
		ContactNumber  string `json:"Contact_Number,omitempty"`
		ID             string `json:"id,omitempty"`
		Approved       bool   `json:"$approved,omitempty"`
		Approval       struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
			Resubmit bool `json:"resubmit,omitempty"`
		} `json:"$approval,omitempty"`
		IsStatusSplitDone bool       `json:"isStatusSplitDone,omitempty"`
		ModifiedTime      Time       `json:"Modified_Time,omitempty"`
		BillingStreet     string     `json:"Billing_Street,omitempty"`
		LastMailedTime    string     `json:"Last_Mailed_Time,omitempty"`
		CreatedTime       Time       `json:"Created_Time,omitempty"`
		Followed          bool       `json:"$followed,omitempty"`
		Editable          bool       `json:"$editable,omitempty"`
		BillingCode       string     `json:"Billing_Code,omitempty"`
		ParentAccount     string     `json:"Parent_Account,omitempty"`
		About             string     `json:"About,omitempty"`
		AssociatedTags    []struct{} `json:"Associated_Tags,omitempty"`
		ShippingCity      string     `json:"Shipping_City,omitempty"`
		ShippingCountry   string     `json:"Shipping_Country,omitempty"`
		ShippingCode      string     `json:"Shipping_Code,omitempty"`
		BillingCity       string     `json:"Billing_City,omitempty"`
		BillingState      string     `json:"Billing_State,omitempty"`
		CreatedBy         struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		Fax                 string `json:"Fax,omitempty"`
		IsAttachmentPresent bool   `json:"Is_Attachment_Present,omitempty"`
		ShippingStreet      string `json:"Shipping_Street,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}
