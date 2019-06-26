package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

// GetBlueprint retrieves a blueprint record specified by the ID parameter from the module specified
// https://www.zoho.com/crm/help/api/v2/#blueprint-api
func (c *API) GetBlueprint(module crmModule, id string) (data BlueprintResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "blueprints",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/actions/blueprint", c.ZohoTLD, module, id),
		Method:       zoho.HTTPGet,
		ResponseData: &BlueprintResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return BlueprintResponse{}, fmt.Errorf("Failed to retrieve blueprint: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*BlueprintResponse); ok {
		return *v, nil
	}

	return BlueprintResponse{}, fmt.Errorf("Data returned was not 'BlueprintResponse'")
}

// BlueprintResponse is the data returned by the GetBlueprint endpoint
type BlueprintResponse struct {
	Blueprint struct {
		ProcessInfo struct {
			FieldID      string `json:"field_id"`
			IsContinuous bool   `json:"is_continuous"`
			APIName      string `json:"api_name"`
			Continuous   bool   `json:"continuous"`
			FieldLabel   string `json:"field_label"`
			Name         string `json:"name"`
			ColumnName   string `json:"column_name"`
			FieldValue   string `json:"field_value"`
			ID           string `json:"id"`
			FieldName    string `json:"field_name"`
		} `json:"process_info"`
		Transitions []struct {
			NextTransitions    []string `json:"next_transitions"`
			PercentPartialSave float64  `json:"percent_partial_save"`
			Data               struct {
				Attachments string `json:"Attachments"`
			} `json:"data"`
			NextFieldValue  string `json:"next_field_value"`
			Name            string `json:"name"`
			CriteriaMatched bool   `json:"criteria_matched"`
			ID              string `json:"id"`
			Fields          []struct {
				DisplayLabel       string `json:"display_label"`
				Type               string `json:"_type"`
				DataType           string `json:"data_type"`
				ColumnName         string `json:"column_name"`
				PersonalityName    string `json:"personality_name"`
				ID                 string `json:"id"`
				TransitionSequence int    `json:"transition_sequence"`
				Mandatory          bool   `json:"mandatory"`
				Layouts            string `json:"layouts"`
			} `json:"fields"`
			CriteriaMessage string `json:"criteria_message"`
		} `json:"transitions"`
	} `json:"blueprint"`
}

// UpdateBlueprint updates the blueprint specified by ID in the specified module
// https://www.zoho.com/crm/help/api/v2/#update-blueprint
func (c *API) UpdateBlueprint(request UpdateBlueprintData, module crmModule, id string) (data UpdateBlueprintResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "blueprints",
		URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/%s/%s/actions/blueprint", c.ZohoTLD, module, id),
		Method:       zoho.HTTPPost,
		ResponseData: &UpdateBlueprintResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return UpdateBlueprintResponse{}, fmt.Errorf("Failed to update blueprint: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*UpdateBlueprintResponse); ok {
		return *v, nil
	}

	return UpdateBlueprintResponse{}, fmt.Errorf("Data returned was not 'UpdateBlueprintResponse'")
}

// UpdateBlueprintData is the data that should be provided to UpdateBlueprint
type UpdateBlueprintData struct {
	Blueprint []struct {
		TransitionID string                 `json:"transition_id"`
		Data         map[string]interface{} `json:"data"`
	} `json:"blueprint"`
}

// UpdateBlueprintResponse is the data returned by UpdateBlueprint
type UpdateBlueprintResponse struct {
	Code    string `json:"code"`
	Details struct {
	} `json:"details"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
