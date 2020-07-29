package main

import "bytes"
import "testing"

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Ney")

	resultado := buffer.String()
	esperado := "Olá, Ney"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}