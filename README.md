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
- [ ] Read up on writing Go Doc comments
- [ ] Comment code religously
- [ ] Write extensive unit tests

## Usage

It is reasonable to assume that each API may provide different implementation, however they should all use the common methods available in Zoho.