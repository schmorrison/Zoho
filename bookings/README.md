# Zoho BOOKINGS V1 API

NOTE: Not finished and probably unstable. PRs welcome.

This API wrapper should provide access to Zoho BOOKINGS. For the GET api's which use query parameters, these should be passed as the type `map[string]zoho.Parameter`. For the POST API's which require RequestBody, the key value pairs should be passed on as variables of type `map[string]string`.

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
            zoho.BuildScope(zoho.Bookings, zoho.DataScope, "", zoho.Create),
        }
        if err := z.AuthorizationCodeRequest("yourClientID", "yourClientSecret", scopes, "http://localhost:8080/oauthredirect"); err != nil {
            log.Fatal(err)
        }

        // Create a new Bookings object and provide the Zoho struct
        c := bookings.New(z)

        // While untested, getting data should work like so
        data := bookings.ServiceResponse{}
        param := make(map[string]zoho.Parameter)
        param["workspace_id"] = "yourWorkspaceID"
        resp, err := c.FetchServices(&data,param)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(resp)
    }

## TODO

- [ ] Write a TODO list
- [ ] Comment code with full details
- [ ] Add page context values to returned data, or methods to interact with it via module
- [ ] Create godoc for bookings APIs
- [ ] Identify Bookings APIs which are used most
