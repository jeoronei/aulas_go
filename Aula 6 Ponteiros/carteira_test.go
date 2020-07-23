package ponteiros

import "testing"

func TestCarteira(t *testing.T) {
	
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(10)
		confirmarSaldo(t, carteira, Bitcoin(10))
	})
	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo :Bitcoin(20)}
		erro := carteira.Retirar(10)
		confirmarSaldo(t, carteira, 10)
		confirmarErroInexistente(t, erro)
	})
	
	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))
		
		confirmarSaldo(t, carteira, saldoInicial)
		confirmarErro(t, erro, erroSaldoInsuficiente)
	})
}

func confirmarSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()
	if resultado != esperado {
		t.Errorf("Esperado '%s', resultado '%s'", esperado, resultado)
	}
}

func confirmarErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	
	if resultado != nil {
		t.Fatal("Esperava um erro mas nenhum ocorreu")
	}
}

func confirmarErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	
	if resultado == nil {
		t.Fatal("Esperava um erro mas nenhum ocorreu")
	}

	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}