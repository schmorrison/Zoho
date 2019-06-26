package crm

import (
	"fmt"
	"time"

	"github.com/schmorrison/Zoho"
)

// ListRecords will return a list of the records provided in the request field, and specified by the module
// https://www.zoho.com/crm/help/api/v2/#record-api
func (c *API) ListRecords(request interface{}, module crmModule, params map[string]zoho.Parameter) (data interface{}, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPGet,
		ResponseData: request,
		URLParameters: map[string]zoho.Parameter{
			"fields":     "",
			"sort_order": "",
			"sort_by":    "",
			"converted":  "",
			"approved":   "",
			"page":       "",
			"per_page":   "200",
			"cvid":       "",
		},
	}

	for k, v := range params {
		endpoint.URLParameters[k] = v
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve records of %s: %s", module, err)
	}

	if endpoint.ResponseData != nil {
		return endpoint.ResponseData, nil
	}

	return nil, fmt.Errorf("Data returned was nil")
}

// InsertRecords will add records in request to the specified module
// https://www.zoho.com/crm/help/api/v2/#ra-insert-records
func (c *API) InsertRecords(request InsertRecordsData, module crmModule) (data InsertRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &InsertRecordsResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InsertRecordsResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*InsertRecordsResponse); ok {
		return *v, nil
	}

	return InsertRecordsResponse{}, fmt.Errorf("Data returned was nil")
}

//UpdateRecordsResponseData is the data provided to UpdateRecords
type UpdateRecordsResponseData struct {
	Message string `json:"message,omitempty"`
	Details struct {
		ExpectedDataType string `json:"expected_data_type,omitempty"`
		APIName          string `json:"api_name"`
	} `json:"details,omitempty"`
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
}

// InsertRecordsData is the data provided to InsertRecords
type InsertRecordsData struct {
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

// InsertRecordsResponse is the data returned by InsertRecords
type InsertRecordsResponse struct {
	Data []InsertRecordsResponseData `json:"data,omitempty"`
}

// UpdateRecords will modify records by the data provided to request in the specified module
// https://www.zoho.com/crm/help/api/v2/#ra-update-records
//
// When performing an update, because the natural state of the records fields in this package is to 'omitempty',
// if you want to empty the fields contents you will need to embed the records type in a struct in your own package,
// and override the field with a field that has a json tag that does not contain 'omitempty'.
// eg.
//    type struct Account {
//        crm.Account
//        CustomField string `json:"Custom_Field"`
//     }
func (c *API) UpdateRecords(request UpdateRecordsData, module crmModule) (data UpdateRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateRecordsResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateRecordsResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateRecordsResponse); ok {
		return *v, nil
	}

	return UpdateRecordsResponse{}, fmt.Errorf("Data returned was nil")
}

// UpdateRecordsData is the data provided to UpdateRecords
type UpdateRecordsData = InsertRecordsData

// UpdateRecordsResponse is the data returned by UpdateRecords
type UpdateRecordsResponse struct {
	Data []UpdateRecordsResponseData `json:"data,omitempty"`
}

// UpsertRecords will insert the provided records in the request, if they already exist it will be updated
// https://www.zoho.com/crm/help/api/v2/#ra-insert-or-update
//
// When performing an upsert, because the natural state of the records fields in this package is to 'omitempty' when encoding json,
// if you want to empty the fields contents in zoho you will need to embed the records type in a struct in your own package,
// and override the field with a field that has a json tag that does not contain 'omitempty'.
// eg.
//    type struct Account {
//        crm.Account
//        CustomField string `json:"Custom_Field"`
//     }
func (c *API) UpsertRecords(request UpsertRecordsData, module crmModule, duplicateFieldsCheck []string) (data UpsertRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/upsert", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &UpsertRecordsResponse{},
		RequestBody:  request,
		URLParameters: map[string]zoho.Parameter{
			"duplicate_field_check": func() zoho.Parameter {
				if len(duplicateFieldsCheck) > 0 {
					fields := ""
					for i, a := range duplicateFieldsCheck {
						fields += a
						if i < len(duplicateFieldsCheck)-1 {
							fields += ","
						}
					}
					return zoho.Parameter(fields)
				}
				return ""
			}(),
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpsertRecordsResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*UpsertRecordsResponse); ok {
		return *v, nil
	}

	return UpsertRecordsResponse{}, fmt.Errorf("Data returned was nil")
}

// UpsertRecordsData is the data provided to UpsertRecords
type UpsertRecordsData struct {
	Data    interface{} `json:"data,omitempty"`
	Trigger []string    `json:"trigger,omitempty"`
}

// UpsertRecordsResponse is the data returned by UpsertRecords
type UpsertRecordsResponse struct {
	Data []struct {
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
		Status         string `json:"status,omitempty"`
		DuplicateField string `json:"duplicate_field,omitempty"`
		Action         string `json:"action,omitempty"`
		Code           string `json:"code,omitempty"`
	} `json:"data,omitempty"`
}

// DeleteRecords will delete the records in the ids in the specified module
// https://www.zoho.com/crm/help/api/v2/#delete-bulk-records
func (c *API) DeleteRecords(module crmModule, ids []string) (data DeleteRecordsResponse, err error) {
	if len(ids) == 0 {
		return DeleteRecordsResponse{}, fmt.Errorf("Failed to delete records, must provide at least 1 ID")
	}

	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteRecordsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"ids": func() zoho.Parameter {
				idStr := ""
				for i, a := range ids {
					idStr += a
					if i < len(ids)-1 {
						idStr += ","
					}
				}
				return zoho.Parameter(idStr)
			}(),
		},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteRecordsResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteRecordsResponse); ok {
		return *v, nil
	}

	return DeleteRecordsResponse{}, fmt.Errorf("Data returned was nil")
}

// DeleteRecordsResponse is the data returned by DeleteRecords
type DeleteRecordsResponse struct {
	Data []struct {
		Code    string `json:"code,omitempty"`
		Details struct {
			ID string `json:"id,omitempty"`
		} `json:"details,omitempty"`
		Message string `json:"message,omitempty"`
		Status  string `json:"status,omitempty"`
	} `json:"data,omitempty"`
}

// ListDeletedRecords will return a list of all records that have been deleted in the specified module. The records can be filtered by the kind parameter.
// https://www.zoho.com/crm/help/api/v2/#ra-deleted-records
func (c *API) ListDeletedRecords(module crmModule, kind DeletedRecordsType, params map[string]zoho.Parameter) (data ListDeletedRecordsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/deleted", c.ZohoTLD, module),
		Method:       zoho.HTTPGet,
		ResponseData: &ListDeletedRecordsResponse{},
		URLParameters: map[string]zoho.Parameter{
			"type":     zoho.Parameter(kind),
			"page":     "",
			"per_page": "200",
		},
	}

	for k, v := range params {
		endpoint.URLParameters[k] = v
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ListDeletedRecordsResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*ListDeletedRecordsResponse); ok {
		return *v, nil
	}

	return ListDeletedRecordsResponse{}, fmt.Errorf("Data returned was nil")
}

// DeletedRecordsType is a type for filtering deleted records
type DeletedRecordsType string

const (
	// AllDeleted - To get the list of all deleted records
	AllDeleted DeletedRecordsType = "All"
	// RecycledDeleted - To get the list of deleted records from recycle bin
	RecycledDeleted DeletedRecordsType = "Recycle"
	// PermanentDeleted - To get the list of permanently deleted records
	PermanentDeleted DeletedRecordsType = "Permanent"
)

// ListDeletedRecordsResponse is the data returned by ListDeletedRecords
type ListDeletedRecordsResponse struct {
	Data []struct {
		DeletedBy struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"deleted_by,omitempty"`
		ID          string `json:"id,omitempty"`
		DisplayName string `json:"display_name,omitempty"`
		Type        string `json:"type,omitempty"`
		CreatedBy   struct {
			Name string `json:"name,omitempty"`
			ID   string `json:"id,omitempty"`
		} `json:"created_by,omitempty"`
		DeletedTime time.Time `json:"deleted_time,omitempty"`
	} `json:"data,omitempty"`
	Info PageInfo `json:"info,omitempty"`
}

// SearchRecords is used for searching records in the specified module using the parameters.
// Parameters are 'criteria', 'email', 'phone', and 'word'
// https://www.zoho.com/crm/help/api/v2/#ra-search-records
func (c *API) SearchRecords(response interface{}, module crmModule, params map[string]zoho.Parameter) (data interface{}, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/search", c.ZohoTLD, module),
		Method:       zoho.HTTPGet,
		ResponseData: response,
		URLParameters: map[string]zoho.Parameter{
			"criteria": "",
			"email":    "",
			"phone":    "",
			"word":     "",
			"page":     "",
			"per_page": "200",
		},
	}

	for k, v := range params {
		endpoint.URLParameters[k] = v
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if endpoint.ResponseData != nil {
		return endpoint.ResponseData, nil
	}

	return nil, fmt.Errorf("Data returned was nil")
}

// GetRecord will retrieve the specified record by id in the specified module.
// https://www.zoho.com/crm/help/api/v2/#single-records
func (c *API) GetRecord(request interface{}, module crmModule, ID string) (data interface{}, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s", c.ZohoTLD, module, ID),
		Method:       zoho.HTTPGet,
		ResponseData: request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve blueprint: %s", err)
	}

	if endpoint.ResponseData != nil {
		return endpoint.ResponseData, nil
	}

	return nil, fmt.Errorf("Data returned was nil")
}

// InsertRecord will insert the specifed record in the module
// https://www.zoho.com/crm/help/api/v2/#create-specify-records
func (c *API) InsertRecord(request InsertRecordData, module crmModule) (data InsertRecordResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s", c.ZohoTLD, module),
		Method:       zoho.HTTPPost,
		ResponseData: &InsertRecordResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return InsertRecordResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*InsertRecordResponse); ok {
		return *v, nil
	}

	return InsertRecordResponse{}, fmt.Errorf("Data returned was nil")
}

// InsertRecordData is the data provided to InsertRecord
type InsertRecordData = InsertRecordsData

// InsertRecordResponse is the data returned by InsertRecord
type InsertRecordResponse = InsertRecordsResponse

// UpdateRecord will update the record specified by ID in the specified module
// https://www.zoho.com/crm/help/api/v2/#update-specify-records
func (c *API) UpdateRecord(request UpdateRecordData, module crmModule, ID string) (data UpdateRecordResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s", c.ZohoTLD, module, ID),
		Method:       zoho.HTTPPut,
		ResponseData: &UpdateRecordResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateRecordResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateRecordResponse); ok {
		return *v, nil
	}

	return UpdateRecordResponse{}, fmt.Errorf("Data returned was nil")
}

// UpdateRecordData is the data provided to UpdateRecord
type UpdateRecordData = InsertRecordsData

// UpdateRecordResponse is the data returned by UpdateRecord
type UpdateRecordResponse = UpdateRecordsResponse

// DeleteRecord will delete the record specified by the id in the specified module
// https://www.zoho.com/crm/help/api/v2/#delete-specify-records
func (c *API) DeleteRecord(module crmModule, ID string) (data DeleteRecordResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s", c.ZohoTLD, module, ID),
		Method:       zoho.HTTPDelete,
		ResponseData: &DeleteRecordResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return DeleteRecordResponse{}, fmt.Errorf("Failed to insert records of %s: %s", module, err)
	}

	if v, ok := endpoint.ResponseData.(*DeleteRecordResponse); ok {
		return *v, nil
	}

	return DeleteRecordResponse{}, fmt.Errorf("Data returned was nil")
}

// DeleteRecordResponse is the data returned by DeleteRecord
type DeleteRecordResponse = DeleteRecordsResponse

// ConvertLead will modify the Lead record specified by ID and convert it to a Contact/Potential depending on the request data
// https://www.zoho.com/crm/help/api/v2/#convert-lead
func (c *API) ConvertLead(request ConvertLeadData, ID string) (data ConvertLeadResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/actions/convert", c.ZohoTLD, LeadsModule, ID),
		Method:       zoho.HTTPPost,
		ResponseData: &ConvertLeadResponse{},
		RequestBody:  data,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return ConvertLeadResponse{}, fmt.Errorf("Failed to insert records of %s: %s", LeadsModule, err)
	}

	if v, ok := endpoint.ResponseData.(*ConvertLeadResponse); ok {
		return *v, nil
	}

	return ConvertLeadResponse{}, fmt.Errorf("Data returned was nil")
}

// ConvertLeadData is the data provided to ConvertLead
type ConvertLeadData struct {
	Data []struct {
		Overwrite            bool   `json:"overwrite,omitempty"`
		NotifyLeadOwner      bool   `json:"notify_lead_owner,omitempty"`
		NotifyNewEntityOwner bool   `json:"notify_new_entity_owner,omitempty"`
		Accounts             string `json:"Accounts,omitempty"`
		Contacts             string `json:"Contacts,omitempty"`
		AssignTo             string `json:"assign_to,omitempty"`
		Deals                struct {
			CampaignSource string  `json:"Campaign_Source,omitempty"`
			DealName       string  `json:"Deal_Name,omitempty"`
			ClosingDate    string  `json:"Closing_Date,omitempty"`
			Stage          string  `json:"Stage,omitempty"`
			Amount         float64 `json:"Amount,omitempty"`
		} `json:"Deals,omitempty"`
	} `json:"data,omitempty"`
}

// ConvertLeadResponse is the data returned by ConvertLead
type ConvertLeadResponse struct {
	Data []struct {
		Contacts string `json:"Contacts,omitempty"`
		Deals    string `json:"Deals,omitempty"`
		Accounts string `json:"Accounts,omitempty"`
	} `json:"data,omitempty"`
}
