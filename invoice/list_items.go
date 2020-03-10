package invoice

import (
    "fmt"
    zoho "github.com/schmorrison/Zoho"
)

func (c *ZohoInvoiceAPI) ListItems() (data ListItemsResponse, err error) {

    // Renew token if necessary
    if c.Zoho.Token.CheckExpiry() {
        err := c.Zoho.RefreshTokenRequest()
        if err != nil {
            return ListItemsResponse{}, err
        }
    }

    endpoint := zoho.Endpoint{
        Name:         ItemsModule,
        URL:          fmt.Sprintf(InvoiceAPIEndPoint + "%s", ItemsModule),
        Method:       zoho.HTTPGet,
        ResponseData: &ListItemsResponse{},
        URLParameters: map[string]zoho.Parameter{
            "filter_by": "",
        },
    }

    /*for k, v := range params {
    	endpoint.URLParameters[k] = v
    }
    */

    err = c.Zoho.HTTPRequest(&endpoint)
    if err != nil {
        return ListItemsResponse{}, fmt.Errorf("Failed to retrieve expense reports: %s", err)
    }

    if v, ok := endpoint.ResponseData.(*ListItemsResponse); ok {
        // Check if the request succeeded
        if v.Code != 0 {
            return *v, fmt.Errorf("Failed to list items: %s", v.Message)
        }
        return *v, nil
    }
    return ListItemsResponse{}, fmt.Errorf("Data retrieved was not 'ListContactsResponse'")
}

type ListItemsResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Items   []struct {
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
    } `json:"items"`
}
