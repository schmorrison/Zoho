# Zoho CRM V2 API

NOTE: Not finished and probably unstable. PRs welcome.

This API wrapper should provide access to Zoho CRM. Because some fields exist only in add-ons for CRM, or are custom fields, which cannot be easily differentiated, all fields that are recieved in a record which have no direct corresponding field in the defined struct will be available in a `map[string]interface{}` field. This may extend to all fields being accessible in a field called "RAW" or some-such, which can then be manually type-asserted against.

Brainstorm: CRM defines many special types for their internal fields (SingleLine, MultiLine, Picklist, MultiSelect, Date, Time, etc.) that must be parsed from the JSON, and provided in a type safe way to this library and back to Zoho. If the records struct fields are made up of mainly these specialized field types (types.go), and each field type has a MarshalJSON/UnmarshalJSON function then parsing may be a little more staight forward. These MarshalJSON/UnmarshalJSON field types will need to handle 'null' cases in the JSON.

Additional Brainstorm: Because of the custom field problem, I believe we can define the definite fields that CRM will return in structs. Then when users that are using this client library need to have custom fields parsed, they can embed the parent struct we defined, into a struct of their own making which defines the custom fields they need.

## Usage
    import (
        "log"
        "fmt"
        "github.com/schmorrison/Zoho"
    )

    func main() {
        // get access/refresh tokens
        z := zoho.New()
        scopes := []zoho.ScopeString{
            zoho.BuildScope(zoho.Crm, zoho.ModulesScope, zoho.AllMethod, zoho.NoOp),
        }
        if err := z.AuthorizationCodeRequest("yourClientID", "yourClientSecret", scopes, "http://localhost:8080/oauthredirect"); err != nil {
            log.Fatal(err)
        }

        // Create a new CRM object and provide the zoho master struct
        c := crm.New(z)

        // While untested, getting data should work like so
        notes, err := c.GetNotes(nil)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(notes)

        // The API for getting module records is bound to change once the returned data types can be defined.
        // The returned JSON values are subject to change given that custom fields are an instrinsic part of zoho. (see brainstorm above)
        data := map[string]interface{}
        _, err := c.ListRecords(&data, crm.AccountsModule, nil)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(data)
    }

## TODO

- [ ] Write a TODO list
- [ ] Comment code with full details
- [ ] Add page context values to returned data, or methods to interact with it via module
- [ ] 
