package entity

import "errors"

var typeSimbol = []string{"BRL", "USD", "BTC", "EUR"}

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
		return errors.New("value must be grater than zero")
	}

	if value := c.validFromCurrency(); value == 0 {
		return errors.New("it is necessary to inform the correct currency")
	}

	if value := c.validFromCurrency(); value == 0 {
		return errors.New("it is necessary to inform the correct currency")
	}

	return nil
}

func (c *Transaction) validFromCurrency() int {
	value := 0
	for index, ts := range typeSimbol {
		if string(ts[index]) == c.FromCurrency {
			value++
		}
	}

	return value
}

func (c *Transaction) validToCurrency() int {
	value := 0
	for index, ts := range typeSimbol {
		if string(ts[index]) == c.FromCurrency {
			value++
		}
	}

	return value
}