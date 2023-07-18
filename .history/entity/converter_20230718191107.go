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
		return errors.New("value must be grater than zero")
	}

	shouldReturn, returnValue := c.newMethod()
	if shouldReturn {
		return returnValue
	}
}

func (c *Transaction) newMethod() (bool, error) {
	value := 0
	for index, ts := range typeSimbol {
		if string(ts[index]) == c.FromCurrency {
			value++
		}
	}
	if value == 0 {
		return true, errors.New("it is necessary to inform the correct currency")
	}
	return false, nil
}
