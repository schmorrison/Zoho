package main

var BlueprintSchema = EndpointSchema{
	Name: "blueprints",
	EndpointList: []EndpointSchemaBase{
		EndpointSchemaBase{
			Comments: []string{
				"GetBlueprint retrieves a blueprint record specified by the ID parameter from the module specified",
				"https://www.zoho.com/crm/help/api/v2/#blueprint-api",
			},
			Endpoint: &Endpoint{
				Name:   "GetBlueprint",
				Method: "GET",
				URI:    "/crm/v2/%s/%s/actions/blueprint",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "module",
					Type: "crmModule",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "id",
					Type: "string",
					Func: true,
					URI:  true,
				},
			},
			Response: "BlueprintResponse",
		},
		EndpointSchemaBase{
			Comments: []string{
				"UpdateBlueprint updates the blueprint specified by ID in the specified module",
				"https://www.zoho.com/crm/help/api/v2/#update-blueprint",
			},
			Endpoint: &Endpoint{
				Name:   "UpdateBlueprint",
				Method: "POST",
				URI:    "/crm/v2/%s/%s/actions/blueprint",
			},
			Params: EndpointParams{
				EndpointParam{
					Name: "module",
					Type: "crmModule",
					Func: true,
					URI:  true,
				},
				EndpointParam{
					Name: "ID",
					Type: "string",
					Func: true,
					URI:  true,
				},
			},
			Request:  "UpdateBlueprintData",
			Response: "UpdateBlueprintResponse",
		},
	},
}
