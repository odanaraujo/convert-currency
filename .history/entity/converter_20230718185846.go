package entity

const (
	REAL    = "BRL"
	DOLAR   = "USD"
	BITCOIN = "BTC"
	EURO    = "EUR"
)

type Transaction struct {
	FromCurrency string
	ToCurrency   string
	ID           uint
	Amount       float64
}

func (c *Transaction) NewTransaction() *Transaction {
	return &Transaction{
		FromCurrency: c.FromCurrency,
		ToCurrency: c.ToCurrency,
		
	}
}
