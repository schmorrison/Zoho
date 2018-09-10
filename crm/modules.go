package crm

import (
	".."
)

var ModulesEndpoint = zoho.Endpoint{
	Name:         "modules",
	URL:          "https://www.zohoapis.com/crm/v2/settings/modules",
	Methods:      []zoho.HttpMethod{zoho.HTTPGet},
	ResponseData: ModulesResponse{},
}

type ModulesResponse struct {
	Modules []struct {
		Convertable   bool      `json:"convertable,omitempty"`
		Editable      bool      `json:"editable,omitempty"`
		Deletable     bool      `json:"deletable,omitempty"`
		WebLink       string    `json:"web_link,omitempty"`
		SingularLabel string    `json:"singular_label,omitempty"`
		ModifiedTime  zoho.Time `json:"modified_time,omitempty"`
		Viewable      bool      `json:"viewable,omitempty"`
		APISupported  bool      `json:"api_supported,omitempty"`
		Createable    bool      `json:"createable,omitempty"`
		PluralLabel   string    `json:"plural_label,omitempty"`
		APIName       string    `json:"api_name,omitempty"`
		ModifiedBy    string    `json:"modified_by,omitempty"`
		GeneratedType string    `json:"generated_type,omitempty"`
		ID            string    `json:"id,omitempty"`
		ModuleName    string    `json:"module_name,omitempty"`
	} `json:"modules,omitempty"`
}
