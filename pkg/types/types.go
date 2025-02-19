package types

// / Money provides a monetary amount in minimal units (cents, kopecks, dirams, etc.)
type Money int64

// / Currency represents currency code
type Currency string

// Currencies codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// / PAN represents card number
type PAN string

// / Card represents information about a payment card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money // Использовали Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money // Минимальный остаток на карте за месяц
}

// / Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string // 'card'
	Number  string // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
