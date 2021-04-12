package types

type Money int64

type Currency string

// Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        string
	Balance    Money // использовали Money
	MinBalance Money // использовали Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type PaymentSource struct {
	Type    string // 'card'
	Number  string // Номер вида '5058 xxxx xxxx 8888'
	Balance int    // баланс в дирамах
}
