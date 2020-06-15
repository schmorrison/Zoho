// Wrapper for Zoho Expense API from https://www.zoho.com/expense/api/v1/

package expense

import (
	zoho "github.com/schmorrison/Zoho"
	"math/rand"
)

// Change here only if these values changes over time
const (
	ExpenseAPIEndpoint       string = "https://expense.zoho.com/api/v1/"
	ExpenseAPIEndpointHeader string = "X-com-zoho-expense-organizationid"
	OrganizationsModule      string = "organizations"
	ExpenseReportModule      string = "expensereports"
	ExpensesModule           string = "expenses"
	TripsModule              string = "trips"
	ExpenseCategoiesModule   string = "expensecategories"
	UsersModule              string = "users"
	CustomersModule          string = "contacts"
	ProjectsModule           string = "projects"
	CurrenciesModule         string = "settings/currencies"
	TaxesModule              string = "settings/taxes"
)

// API is used for interacting with the Zoho expense API
// the exposed methods are primarily access to expense modules which provide access to expense Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *expense.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			id = append(id, keyspace[rand.Intn(len(keyspace))])
		}
		return string(id)
	}()
	API := &API{
		Zoho: z,
		id:   id,
	}
	return API
}
