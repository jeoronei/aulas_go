package estruturas

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo {10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("Esperado '%.2f', resultado '%.2f'", esperado, resultado)
	}
}

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()

		resultado := forma.Area()
		if resultado != esperado {
			t.Errorf("Esperado '%.2f', resultado '%.2f'", esperado, resultado)
		}
	}

	t.Run("retângulo", func(t *testing.T) {
		retangulo := Retangulo {12.0, 6.0}
		verificaArea(t, retangulo, 72.0)
	})

	t.Run("círculo", func(t *testing.T) {
		circulo := Circulo {10.0}
		verificaArea(t, circulo, 314.1592653589793)
	})
}

func TestAreaUtilizandoTableDrivenTest(t *testing.T) {
	testesArea := []struct {
		nome string
		forma Forma
		esperado float64
	}{
		{nome: "Retângulo",forma: Retangulo {12.0, 6.0}, esperado: 72.0},
		{nome: "Cícrulo",forma: Circulo {10}, esperado: 314.1592653589793},
		{nome: "Triângulo",forma: Triangulo{12,6}, esperado: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.esperado {
				t.Errorf("%#v Esperado '%.2f', resultado '%.2f'", tt.forma, tt.esperado, resultado)
			}
		})
	}
}