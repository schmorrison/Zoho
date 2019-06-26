package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetModules returns the list of modules available in the CRM account
// https://www.zoho.com/crm/help/api/v2/#Modules-APIs
func (c *API) GetModules() (data ModulesResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "modules",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/settings/modules", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &ModulesResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ModulesResponse{}, fmt.Errorf("Failed to retrieve modules: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ModulesResponse); ok {
		return *v, nil
	}

	return ModulesResponse{}, fmt.Errorf("Data retrieved was not 'ModuleResponse'")
}

// ModulesResponse is the data returned by GetModules
type ModulesResponse struct {
	Modules []struct {
		Convertable   bool   `json:"convertable,omitempty"`
		Editable      bool   `json:"editable,omitempty"`
		Deletable     bool   `json:"deletable,omitempty"`
		WebLink       string `json:"web_link,omitempty"`
		SingularLabel string `json:"singular_label,omitempty"`
		ModifiedTime  Time   `json:"modified_time,omitempty"`
		Viewable      bool   `json:"viewable,omitempty"`
		APISupported  bool   `json:"api_supported,omitempty"`
		Createable    bool   `json:"createable,omitempty"`
		PluralLabel   string `json:"plural_label,omitempty"`
		APIName       string `json:"api_name,omitempty"`
		ModifiedBy    string `json:"modified_by,omitempty"`
		GeneratedType string `json:"generated_type,omitempty"`
		ID            string `json:"id,omitempty"`
		ModuleName    string `json:"module_name,omitempty"`
	} `json:"modules,omitempty"`
}
