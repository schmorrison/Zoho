# Zoho CRM V2 API

NOTE: Not finished and probably unstable. PRs welcome.

This API wrapper should provide access to Zoho CRM. Because some fields exist only in add-ons for CRM, or are custom fields, which cannot be easily differentiated, all fields that are recieved in a record which have no direct corresponding field in the defined struct will be available in a ```map[string]interface{}``` field.

## Usage

    // Create a new Zoho object
    z := zoho.New()
	z.SetAuthToken(<API Key/Auth Token>, "ZohoCRM/crmapi")

    // Create a new CRM object
	c := crm.New(z)

    // Get all sales orders
	data, _ := c.SalesOrders().GetMyRecords(crm.GetRecordsOptions{})

    fmt.Println(data)

## TODO

- [ ] Comment code with full details
- [ ] Add page context values to returned data, or methods to interact with it via module
- [ ] 