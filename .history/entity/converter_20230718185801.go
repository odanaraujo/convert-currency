package entity

const (
	REAL    = "BRL"
	DOLAR   = "USD"
	BITCOIN = "BTC"
	EURO    = "EUR"
)

type Converter struct {
	FromCurrency string
	ToCurrency   string
	ID           uint
	Amount       float64
}

func (c *Converter)