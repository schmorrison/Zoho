package main

var ModulesSchema = EndpointSchema{
	Name: "modules",
	EndpointList: []EndpointSchemaBase{
		EndpointSchemaBase{
			Comments: []string{
				"GetModules returns the list of modules available in the CRM account",
				"https://www.zoho.com/crm/help/api/v2/#Modules-APIs",
			},
			Endpoint: &Endpoint{
				Name:   "GetModules",
				Method: "GET",
				URI:    "/crm/v2/settings/modules",
			},
			Response: "ModulesResponse",
		},
	},
}
