package crm

import (
	"math/rand"
	"time"

	"github.com/schmorrison/Zoho"
)

type crmModule string

// Proper names for CRM modules
const (
	AccountsModule       crmModule = "Accounts"
	CallsModule          crmModule = "Calls"
	CampaignsModule      crmModule = "Campaigns"
	CasesModule          crmModule = "Cases"
	ContactsModule       crmModule = "Contacts"
	CustomModule         crmModule = "Custom"
	DealsModule          crmModule = "Deals"
	EventsModule         crmModule = "Events"
	InvoicesModule       crmModule = "Invoices"
	LeadsModule          crmModule = "Leads"
	PotentialsModule     crmModule = "Potentials"
	PriceBooksModule     crmModule = "PriceBooks"
	ProductsModule       crmModule = "Products"
	PurchaseOrdersModule crmModule = "PurchaseOrders"
	QuotesModule         crmModule = "Quotes"
	SalesOrdersModule    crmModule = "SalesOrders"
	SolutionsModule      crmModule = "Solutions"
	TasksModule          crmModule = "Tasks"
	VendorsModule        crmModule = "Vendors"
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
