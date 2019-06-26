[![](https://godoc.org/github.com/schmorrison/Zoho?status.svg)](http://godoc.org/github.com/schmorrison/Zoho)

# Golang API Wrapper for Zoho Services

This repo is an attempt to build a comprehensive API wrapper for Zoho Services.

This will be a long project, with alot of boilerplate code that may benefit from code generation. Pull requests would be appreciated.

- [ ] [Books](https://github.com/schmorrison/Zoho/tree/master/books)
- [ ] Campaigns
- [ ] Cliq
- [ ] Creator
- [ ] [CRM](https://github.com/schmorrison/Zoho/tree/master/crm)
- [ ] Desk
- [ ] Docs
- [ ] Inventory
- [ ] Invoice
- [ ] Mail
- [ ] Meeting
- [ ] People
- [ ] Recruit
- [ ] Reports
- [ ] Subscriptions

The API's should ideally be useful and obvious. However, as it stands, the Zoho CRM API returns alot of dynamically typed fields which became incredibly difficult to parse, which eventually resulted in an implementation using reflect and a type switch to cast/convert the value from Zoho into the expected value for the struct. I expect this to be the case for alot of Zoho services.

I will try to comment the code religously, and will read up on Go Doc so the generated documentation is useful for users.
- [ ] Write a TODO list
- [ ] Read up on writing Go Doc comments
- [ ] Comment code religously
- [ ] Write extensive unit tests
- [ ] Start versioning commits to prevent major breaks

## A special thanks to Contributors
- Thanks to @ashishsnigam for pull request #7.
- Thanks to @beatscode for pull requests #10, #11, #12, & #13.
- Thanks to @meyskens for pull request #15.

## Usage

It is reasonable to assume that each API may provide different implementation, however they should all use the common methods available in Zoho.

### Getting the master struct and starting oAuth2 flow

    import (
        "github.com/schmorrison/Zoho"
        "log"
    )

    func main() {
        z := zoho.New()

        // to start oAuth2 flow
        scopes := []zoho.ScopeString{
            zoho.BuildScope(zoho.Crm, zoho.ModulesScope, zoho.AllMethod, zoho.NoOp),
        }

        // The authorization request will provide a link that must be clicked on or pasted into a browser.
        // Sometimes it will show the consent screen, upon consenting it will redirect to the redirectURL (currently the server doesn't return a value to the browser once getting the code)
        // The redirectURL provided here must match the URL provided when generating the clientID/secret
        // if the provided redirectURL is a localhost domain, the function will create a server on that port (use non-privileged port), and wait for the redirect to occur.
        // if the redirect provides the authorization code in the URL parameter "code", then the server catches it and provides it to the function for generating AccessToken and RefreshToken

        if err := z.AuthorizationCodeRequest("yourClientID", "yourClientSecret", scopes, "http://localhost:8080/oauthredirect"); err != nil {
            log.Fatal(err)
        }
    }

Alternatively, you may not want to have to click on the link. Perhaps you are running a script on cron, or otherwise. In these case you will want to generate the authorization code manually. This can be done by going to the zoho accounts developer console, and clicking the kebab icon (3 vertical dots) beside the specified token. Click on the 'Self-Client' option, it will prompt you to enter your scopes, and an expiry time. Then it will show you your authorization code.

That code can be used to request Access and Request tokens as so.

    import (
        "log"
        "github.com/schmorrison/Zoho"
    )

    func main() {
        z := zoho.New()

        // to start oAuth2 flow
        scopes := []zoho.ScopeString{
            zoho.BuildScope(zoho.Crm, zoho.ModulesScope, zoho.AllMethod, zoho.NoOp),
        }

        if err := z.GenerateTokenRequest("yourClientID", "yourClientSecret", "authorizationCode", "redirectURL"); err != nil {
            log.Fatal(err)
        }

    }

Your Zoho master struct now has the oAuth token for that service/scope combination.

Check the Readme in each services directory for information about using that service
