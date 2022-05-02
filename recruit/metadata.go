package recruit

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

// GetAllMetadata returns the metadata for fields, layouts, and related lists for the specified module.
// It lists the entire fields available and related list for that module.
// https://www.zoho.com/recruit/developer-guide/apiv2/module-meta.html
// https://recruit.zoho.%s/v2/settings/modules
func (c *API) GetAllMetadata() (data AllMetadataResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetAllMetadata",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/v2/settings/modules", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &AllMetadataResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return AllMetadataResponse{}, fmt.Errorf("failed to retrieve modules: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*AllMetadataResponse); ok {
		return *v, nil
	}

	return AllMetadataResponse{}, fmt.Errorf("data retrieved was not 'ModuleResponse'")
}

// AllMetadataResponse is the data returned by GetAllMetadata

type AllMetadataResponse struct {
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

// GetModuleMetadata returns the metadata for a specific module.
// https://www.zoho.com/recruit/developer-guide/apiv2/module-meta.html
// https://recruit.zoho.eu/recruit/v2/settings/modules/{module}
func (c *API) GetModuleMetadata(module string) (data ModuleMetadataResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetModuleMetadata",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/modules/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPGet,
		ResponseData: &ModuleMetadataResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ModuleMetadataResponse{}, fmt.Errorf("failed to retrieve metadata module: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ModuleMetadataResponse); ok {
		return *v, nil
	}

	return ModuleMetadataResponse{}, fmt.Errorf("data retrieved was not 'ModuleResponse'")
}

// ModuleMetadataResponse is the data returned by GetModuleMetadata
type ModuleMetadataResponse struct {
	Modules []struct {
		ImportMy              bool        `json:"Import_My"`
		GlobalSearchSupported bool        `json:"global_search_supported"`
		Deletable             bool        `json:"deletable"`
		Creatable             bool        `json:"creatable"`
		ModifiedTime          interface{} `json:"modified_time"`
		PluralLabel           string      `json:"plural_label"`
		PresenceSubMenu       bool        `json:"presence_sub_menu"`
		Export                bool        `json:"Export"`
		ID                    string      `json:"id"`
		RelatedListProperties struct {
			SortBy    interface{} `json:"sort_by"`
			Fields    []string    `json:"fields"`
			SortOrder interface{} `json:"sort_order"`
		} `json:"related_list_properties"`
		Properties           []string `json:"$properties"`
		PerPage              int      `json:"per_page"`
		Convertable          bool     `json:"convertable"`
		Editable             bool     `json:"editable"`
		EmailTemplateSupport bool     `json:"emailTemplate_support"`
		Profiles             []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"profiles"`
		FilterSupported bool `json:"filter_supported"`
		DisplayField    struct {
			APIName string `json:"api_name"`
			ID      string `json:"id"`
		} `json:"display_field"`
		WebLink          interface{} `json:"web_link"`
		SequenceNumber   int         `json:"sequence_number"`
		SingularLabel    string      `json:"singular_label"`
		Viewable         bool        `json:"viewable"`
		APISupported     bool        `json:"api_supported"`
		APIName          string      `json:"api_name"`
		QuickCreate      bool        `json:"quick_create"`
		ModifiedBy       interface{} `json:"modified_by"`
		GeneratedType    string      `json:"generated_type"`
		ScoringSupported bool        `json:"scoring_supported"`
		ModuleName       string      `json:"module_name"`
		Attachmenttypes  []struct {
			SequenceNumber int    `json:"sequence_number"`
			CustomField    bool   `json:"custom_field"`
			APIName        string `json:"api_name"`
			Isauto         bool   `json:"isauto"`
			FieldLabel     string `json:"field_label"`
			Publish        bool   `json:"publish"`
			ViewType       struct {
				View     bool `json:"view"`
				Download bool `json:"download"`
				Edit     bool `json:"edit"`
				Create   bool `json:"create"`
				Delete   bool `json:"delete"`
			} `json:"view_type"`
			SePresence bool   `json:"se_presence"`
			ID         string `json:"id"`
			Isdefault  bool   `json:"isdefault"`
			Bulk       bool   `json:"bulk"`
			Required   bool   `json:"required"`
		} `json:"attachmenttypes"`
		ImportMyOrg bool `json:"Import_MyOrg"`
		CustomView  struct {
			DisplayValue  string      `json:"display_value"`
			SharedType    interface{} `json:"shared_type"`
			Criteria      interface{} `json:"criteria"`
			SystemName    string      `json:"system_name"`
			SharedDetails interface{} `json:"shared_details"`
			SortBy        interface{} `json:"sort_by"`
			Offline       bool        `json:"offline"`
			Default       bool        `json:"default"`
			SystemDefined bool        `json:"system_defined"`
			Name          string      `json:"name"`
			ID            string      `json:"id"`
			Category      string      `json:"category"`
			Fields        []string    `json:"fields"`
			Favorite      interface{} `json:"favorite"`
			SortOrder     interface{} `json:"sort_order"`
			IsSearch      bool        `json:"is_search"`
		} `json:"custom_view"`
		ParentModule struct {
		} `json:"parent_module"`
	} `json:"modules"`
}

// GetFieldsMetadata returns field metadata for the specified module.
// https://www.zoho.com/recruit/developer-guide/apiv2/field-meta.html
// https://recruit.zoho.eu/recruit/v2/settings/fields
func (c *API) GetFieldsMetadata(params map[string]zoho.Parameter) (data FieldsMetadataResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetFieldsMetadata",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/fields", c.ZohoTLD),
		Method:       zoho.HTTPGet,
		ResponseData: &FieldsMetadataResponse{},
		URLParameters: map[string]zoho.Parameter{
			"module": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return FieldsMetadataResponse{}, fmt.Errorf("failed to retrieve fields: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*FieldsMetadataResponse); ok {
		return *v, nil
	}

	return FieldsMetadataResponse{}, fmt.Errorf("data retrieved was not 'ModuleResponse'")
}

type FieldsMetadataResponse struct {
	Fields []struct {
		SystemMandatory       bool        `json:"system_mandatory"`
		Webhook               bool        `json:"webhook"`
		JSONType              string      `json:"json_type"`
		FieldLabel            string      `json:"field_label"`
		Tooltip               interface{} `json:"tooltip"`
		CreatedSource         string      `json:"created_source"`
		FieldReadOnly         bool        `json:"field_read_only"`
		DisplayLabel          string      `json:"display_label"`
		ReadOnly              bool        `json:"read_only"`
		BusinesscardSupported bool        `json:"businesscard_supported"`
		Currency              struct {
		} `json:"currency"`
		ID          string `json:"id"`
		CustomField bool   `json:"custom_field"`
		Lookup      struct {
		} `json:"lookup"`
		ConvertMapping struct {
			Potentials interface{} `json:"Potentials"`
			Contacts   interface{} `json:"Contacts"`
			Accounts   interface{} `json:"Accounts"`
		} `json:"convert_mapping"`
		Visible  bool `json:"visible"`
		Length   int  `json:"length"`
		ViewType struct {
			View        bool `json:"view"`
			Edit        bool `json:"edit"`
			QuickCreate bool `json:"quick_create"`
			Create      bool `json:"create"`
		} `json:"view_type"`
		APIName string `json:"api_name"`
		Unique  struct {
		} `json:"unique"`
		DataType string `json:"data_type"`
		Formula  struct {
		} `json:"formula"`
		DecimalPlace       interface{}   `json:"decimal_place"`
		MassUpdate         bool          `json:"mass_update"`
		BlueprintSupported bool          `json:"blueprint_supported"`
		PickListValues     []interface{} `json:"pick_list_values"`
		AutoNumber         struct {
		} `json:"auto_number"`
	} `json:"fields"`
}

// GetCustomViewsMetadata returns the custom views data of a particular module.
// https://www.zoho.com/recruit/developer-guide/apiv2/custom-view-meta.html
// https://recruit.zoho.eu/recruit/v2/settings/custom_views/{module_id}?module={module}
func (c *API) GetCustomViewsMetadata(moduleId string, params map[string]zoho.Parameter) (data CustomViewsMetadataResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "GetCustomViewsMetadata",
		URL:          fmt.Sprintf("https://recruit.zoho.%s/recruit/v2/settings/custom_views/%s", c.ZohoTLD, moduleId),
		Method:       zoho.HTTPGet,
		ResponseData: &CustomViewsMetadataResponse{},
		URLParameters: map[string]zoho.Parameter{
			"module": "",
		},
	}

	if len(params) > 0 {
		for k, v := range params {
			endpoint.URLParameters[k] = v
		}
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return CustomViewsMetadataResponse{}, fmt.Errorf("failed to retrieve custom views: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CustomViewsMetadataResponse); ok {
		return *v, nil
	}

	return CustomViewsMetadataResponse{}, fmt.Errorf("data retrieved was not 'ModuleResponse'")
}

type CustomViewsMetadataResponse struct {
	CustomViews []struct {
		DisplayValue  string      `json:"display_value"`
		SharedType    interface{} `json:"shared_type"`
		Criteria      interface{} `json:"criteria"`
		SystemName    string      `json:"system_name"`
		SharedDetails interface{} `json:"shared_details"`
		SortBy        string      `json:"sort_by"`
		Offline       bool        `json:"offline"`
		Default       bool        `json:"default"`
		SystemDefined bool        `json:"system_defined"`
		Name          string      `json:"name"`
		ID            string      `json:"id"`
		Category      string      `json:"category"`
		Fields        []string    `json:"fields"`
		Favorite      interface{} `json:"favorite"`
		SortOrder     string      `json:"sort_order"`
		IsSearch      bool        `json:"is_search"`
	} `json:"custom_views"`
}
