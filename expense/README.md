[![](https://godoc.org/github.com/schmorrison/Zoho/expense?status.svg)](http://godoc.org/github.com/schmorrison/Zoho/expense)
# Zoho EXPENSE V1 API

NOTE: Not finished and probably unstable. PRs welcome.

This API wrapper should provide access to Zoho EXPENSE. Because some fields exist only in add-ons for EXPENSE, or are custom fields, which cannot be easily differentiated, all fields that are recieved in a record which have no direct corresponding field in the defined struct will be available in a `map[string]interface{}` field. This may extend to all fields being accessible in a field called "RAW" or some-such, which can then be manually type-asserted against.

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
            zoho.BuildScope(zoho.Expense, zoho.FullAccessScope, zoho.AllMethod, zoho.NoOp),
        }
        if err := z.AuthorizationCodeRequest("yourClientID", "yourClientSecret", scopes, "http://localhost:8080/oauthredirect"); err != nil {
            log.Fatal(err)
        }

        // Create a new Expense object and provide the zoho master struct
        c := expense.New(z)

        // While untested, getting data should work like so
        data := expense.ExpenseReportResponse{}
        _, err := c.GetExpenseReports(&data, "yourorganizationid", nil)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(data)
    }

## TODO

- [ ] Write a TODO list
- [ ] Comment code with full details
- [ ] Add page context values to returned data, or methods to interact with it via module
- [ ] Create godoc for expense APIs
- [ ] Identify Expense APIs which are used most

