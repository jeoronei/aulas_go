package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Ney")
	esperado := "Olá, Ney"

	if resultado != esperado {
		t.Errorf("resultado: %s, esperado: %s", resultado, esperado)
	}
}