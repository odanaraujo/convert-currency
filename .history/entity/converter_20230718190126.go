package entity

const (
	REAL    = "BRL"
	DOLAR   = "USD"
	BITCOIN = "BTC"
	EURO    = "EUR"
)

type Transaction struct {
	ID           uint
	FromCurrency string
	ToCurrency   string
	Amount       float64
}

func (c *Transaction) NewTransaction() *Transaction {
	return &Transaction{
		ID:           c.ID,
		FromCurrency: c.FromCurrency,
		ToCurrency:   c.ToCurrency,
		Amount:       c.Amount,
	}
}

func (c *Transaction) Valid()
