// Wrapper for Zoho Expense API from https://www.zoho.com/expense/api/v1/

package expense

import (
	"github.com/schmorrison/Zoho"
	"math/rand"
)

type expenseModule string

// Change here only if these values changes over time
const (
	ExpenseAPIEndPoint     expenseModule = "https://expense.zoho.com/api/v1/"
	OrganizationsModule    expenseModule = "organizations"
	ExpenseReportModule    expenseModule = "expensereports"
	ExpensesModule         expenseModule = "expenses"
	TripsModule            expenseModule = "trips"
	ExpenseCategoiesModule expenseModule = "expensecategories"
	UsersModule            expenseModule = "users"
	CustomersModule        expenseModule = "contacts"
	ProjectsModule         expenseModule = "projects"
	CurrenciesModule       expenseModule = "settings/currencies"
	TaxesModule            expenseModule = "settings/taxes"
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

	return &API{
		Zoho: z,
		id:   id,
	}
}
