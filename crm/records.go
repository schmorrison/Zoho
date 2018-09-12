package crm

import (
	"fmt"

	"github.com/schmorrison/Zoho"
)

func (c *API) ListRecords(request interface{}, module crmModule, params map[string]zoho.Parameter) (data interface{}, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/%s", module),
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

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve blueprint: %s", err)
	}

	if endpoint.ResponseData != nil {
		return endpoint.ResponseData, nil
	}

	return nil, fmt.Errorf("Data returned was not 'BlueprintResponse'")
}

func (c *API) GetRecord(request interface{}, module crmModule, ID string, params map[string]zoho.Parameter) (data interface{}, err error) {
	endpoint := zoho.Endpoint{
		Name:         "records",
		URL:          fmt.Sprintf("https://www.zohoapis.com/crm/v2/%s/%s", module, ID),
		Method:       zoho.HTTPGet,
		ResponseData: request,
	}

	err = c.Zoho.HttpRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve blueprint: %s", err)
	}

	if endpoint.ResponseData != nil {
		return endpoint.ResponseData, nil
	}

	return nil, fmt.Errorf("Data returned was not 'BlueprintResponse'")
}
