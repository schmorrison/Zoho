package invoice

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

type CreateItemResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Item    struct {
		ItemID        string  `json:"item_id"`
		Name          string  `json:"name"`
		Status        string  `json:"status"`
		Description   string  `json:"description"`
		Rate          float64 `json:"rate"`
		Unit          string  `json:"unit"`
		TaxID         string  `json:"tax_id"`
		TaxName       string  `json:"tax_name"`
		TaxPercentage float64 `json:"tax_percentage"`
		TaxType       string  `json:"tax_type"`
		SKU           string  `json:"sku"`
		ProductType   string  `json:"product_type"`
	} `json:"item"`
}

type CreateItemRequest struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Rate         string        `json:"rate"`
	TaxID        string        `json:"tax_id"`
	CustomFields []interface{} `json:"custom_fields"`
	ItemType     string        `json:"item_type"`
	ProductType  string        `json:"product_type"`
	Unit         string        `json:"unit"`
}

func (c *API) CreateItem(request CreateItemRequest) (data CreateItemResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         ItemsModule,
		URL:          fmt.Sprintf("https://invoice.zoho.%s/api/v3/%s", c.ZohoTLD, ItemsModule),
		Method:       zoho.HTTPPost,
		ResponseData: &CreateItemResponse{},
		RequestBody:  request,
		BodyFormat:   zoho.JSON_STRING,
		Headers: map[string]string{
			InvoiceAPIEndpointHeader: c.OrganizationID,
		},
	}

	if err = c.Zoho.HTTPRequest(&endpoint); err != nil {
		return CreateItemResponse{}, fmt.Errorf("Failed to create item: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*CreateItemResponse); ok {
		// Check if the request succeeded
		if v.Code != 0 {
			return *v, fmt.Errorf("Failed to create item: %s", v.Message)
		}
		return *v, nil
	}

	return CreateItemResponse{}, fmt.Errorf("Data retrieved was not 'CreateItemResponse'")
}
