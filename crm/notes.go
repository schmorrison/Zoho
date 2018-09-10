package crm

import (
	".."
)

var NotesEndpoint = zoho.Endpoint{
	Name:         "notes",
	URL:          "https://www.zohoapis.com/crm/v2/${module}/${id}/Notes",
	Methods:      []zoho.HttpMethod{zoho.HTTPGet},
	ResponseData: NotesResponse{},
	URLParameters: map[string]zoho.Parameter{
		"page":     "",
		"per_page": "200",
	},
	OptionalSegments: map[string]string{
		"module": "",
		"id":     "",
	},
}

type NotesModules string

const (
	Leads          NotesModules = "Leads"
	Accounts       NotesModules = "Accounts"
	Contacts       NotesModules = "Contacts"
	Deals          NotesModules = "Deals"
	Campaigns      NotesModules = "Campaigns"
	Tasks          NotesModules = "Tasks"
	Events         NotesModules = "Events"
	Calls          NotesModules = "Calls"
	Cases          NotesModules = "Cases"
	Solutions      NotesModules = "Solutions"
	Products       NotesModules = "Products"
	Vendors        NotesModules = "Vendors"
	Pricebooks     NotesModules = "Pricebooks"
	Quotes         NotesModules = "Quotes"
	SalesOrders    NotesModules = "SalesOrders"
	PurchaseOrders NotesModules = "PurchaseOrders"
	Invoices       NotesModules = "Invoices"
	Custom         NotesModules = "Custom"
)

type NotesResponse struct {
	Data []struct {
		Owner struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Owner,omitempty"`
		SeModule string `json:"$se_module,omitempty"`
		Approval struct {
			Delegate bool `json:"delegate,omitempty"`
			Approve  bool `json:"approve,omitempty"`
			Reject   bool `json:"reject,omitempty"`
		} `json:"$approval,omitempty"`
		ModifiedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Modified_By,omitempty"`
		ModifiedTime zoho.Time `json:"Modified_Time,omitempty"`
		CreatedTime  zoho.Time `json:"Created_Time,omitempty"`
		Followed     bool      `json:"$followed,omitempty"`
		ParentID     struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Parent_Id,omitempty"`
		ID        string `json:"id,omitempty"`
		CreatedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"Created_By,omitempty"`
		NoteTitle   string `json:"Note_Title,omitempty"`
		NoteContent string `json:"Note_Content,omitempty"`
	} `json:"data,omitempty"`
	Info struct {
		PerPage     int  `json:"per_page,omitempty"`
		Count       int  `json:"count,omitempty"`
		Page        int  `json:"page,omitempty"`
		MoreRecords bool `json:"more_records,omitempty"`
	} `json:"info,omitempty"`
}
