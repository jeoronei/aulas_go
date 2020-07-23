package ponteiros

import "fmt"
import "errors"

var erroSaldoInsuficiente = errors.New("Saldo insuficiente!")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct { 
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if c.saldo < quantidade {
		return erroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}