package entity

const (
	REAL    = "BRL"
	DOLAR   = "USD"
	BITCOIN = "BTC"
	EURO    = "EUR"
)

type Converter struct {
	ID uint
	Amount float64
	FromCurrency string
	

}