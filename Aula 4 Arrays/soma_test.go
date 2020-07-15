package arrays

import "testing"
import "reflect"

func TestSoma(t *testing.T) {
	t.Run("coleção com 5 valores", func(t *testing.T) {
		numeros := []int {1,2,3,4,5}
		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado '%d', esperado '%d', dados '%v'", resultado, esperado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int {1,2}, []int {0,9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(esperado, resultado) {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {
	validaResultado := func(t *testing.T, esperado, resultado []int) {
		t.Helper()
		
		if !reflect.DeepEqual(esperado, resultado) {
			t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
		}
	}

	t.Run("realiza soma com alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int {1,2}, []int {0,9})
		esperado := []int{2, 9}
		validaResultado(t, esperado, resultado)
	})

	t.Run("realiza a soma com slice vazio", func(t *testing.T) {
		resultado := SomaTodoOResto([]int {}, []int {3,4,5})
		esperado := []int{0, 9}
		validaResultado(t, esperado, resultado)
	})
}