package crm

import (
	"math/rand"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

type Module string

// Proper names for CRM modules
const (
	AccountsModule       Module = "Accounts"
	CallsModule          Module = "Calls"
	CampaignsModule      Module = "Campaigns"
	CasesModule          Module = "Cases"
	ContactsModule       Module = "Contacts"
	CustomModule         Module = "Custom"
	DealsModule          Module = "Deals"
	EventsModule         Module = "Events"
	InvoicesModule       Module = "Invoices"
	LeadsModule          Module = "Leads"
	PotentialsModule     Module = "Potentials"
	PriceBooksModule     Module = "PriceBooks"
	ProductsModule       Module = "Products"
	PurchaseOrdersModule Module = "PurchaseOrders"
	QuotesModule         Module = "Quotes"
	SalesOrdersModule    Module = "SalesOrders"
	SolutionsModule      Module = "Solutions"
	TasksModule          Module = "Tasks"
	VendorsModule        Module = "Vendors"
)

// API is used for interacting with the Zoho CRM API
// the exposed methods are primarily access to CRM modules which provide access to CRM Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *crm.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			source := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(source)
			id = append(id, keyspace[rnd.Intn(len(keyspace))])
		}
		return string(id)
	}()

	return &API{
		Zoho: z,
		id:   id,
	}
}
