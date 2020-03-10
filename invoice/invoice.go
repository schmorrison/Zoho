package invoice

import (
	"github.com/schmorrison/Zoho"
	"math/rand"
)

const (
	InvoiceAPIEndPoint      string = "https://invoice.zoho.com/api/v3/"
	ContactsModule          string = "contacts"
	ContactsPersonSubModule string = "contactpersons"
	InvoicesModule          string = "invoices"
	ItemsModule             string = "items"
	RecurringInvoicesModule string = "recurringinvoices"
	CustomerPaymentsModule  string = "customerpayments"
)

type CustomFieldRequest struct {
	CustomfieldID string `json:"customfield_id,omitempty"`
	Label         string `json:"label"`
	Value         string `json:"value,omitempty"`
}

// ZohoInvoiceAPI is used for interacting with the Zoho expense ZohoInvoiceAPI
// the exposed methods are primarily access to expense modules which provide access to expense Methods
type ZohoInvoiceAPI struct {
	*zoho.Zoho
	id string
	organisationId string
}

// New returns a *expense.ZohoInvoiceAPI with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho, organisationId string) *ZohoInvoiceAPI {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			id = append(id, keyspace[rand.Intn(len(keyspace))])
		}
		return string(id)
	}()

	zohoInvoiceAPI := ZohoInvoiceAPI{
		Zoho: z,
		id:   id,
		organisationId: organisationId,
	}
	zohoInvoiceAPI.SetOrganizationID(organisationId)
	return &zohoInvoiceAPI
}
