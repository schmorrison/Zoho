[![](https://godoc.org/github.com/schmorrison/Zoho/recruit?status.svg)](http://godoc.org/github.com/schmorrison/Zoho/recruit)
# Zoho RECRUIT V1 API

NOTE: Preatty much stable and tested. PRs welcome.

This API wrapper should provide access to Zoho RECRUIT. Because some fields exist only in add-ons for RECRUIT, or are custom fields, which cannot be easily differentiated, all fields that are recieved in a record which have no direct corresponding field in the defined struct will be available in a `map[string]interface{}` field. This may extend to all fields being accessible in a field called "RAW" or some-such, which can then be manually type-asserted against.

Note: These APIs result entire API response in JSON format. It is expected that client application will parse the API response to get the desired fields.
 fields they need.

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
            zoho.BuildScope(zoho.Recruit, zoho.FullAccessScope, zoho.AllMethod, zoho.NoOp),
        }
        if err := z.AuthorizationCodeRequest("yourClientID", "yourClientSecret", scopes, "http://localhost:8080/oauthredirect"); err != nil {
            log.Fatal(err)
        }

        // Create a new Recruit object and provide the Zoho struct
        c := recruit.New(z)

        // While untested, getting data should work like so
        params := map[string]zoho.Parameter{}
        _, err := c.GetJobOpenings(params)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(data)
    }

## TODO

- [ ] Write a TODO list
- [ ] Comment code with full details
- [ ] Add page context values to returned data, or methods to interact with it via module
- [ ] Create godoc for recruit APIs
- [ ] Identify Recruit APIs which are used most

