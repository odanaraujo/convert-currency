package entity

import "errors"

var typeSimbol = []string  {"BRL", "USD", "BTC", "EUR"}

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

func (c *Transaction) Valid() error {
	if c.Amount == 0 {
		return errors.New("Value must be grater than zero")
	}

	for ts := range typeSimbol {
		if ts[] == c.FromCurrency {

		}
	}
}
