package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	verificarResultados := func (t *testing.T, repeticoes, esperado string) {
		t.Helper()
		
		if repeticoes != esperado {
			t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
		}
	}

	t.Run("Repetir 'a' 5 vezes", func(t *testing.T) {
		esperado := "aaaaa"
		repeticoes := Repetir("a", 5)

		verificarResultados(t, repeticoes, esperado)
	})

	t.Run("Repetir 'bla' 3 vezes", func(t *testing.T) {
		esperado := "blablabla"
		repeticoes := Repetir("bla", 3)

		verificarResultados(t, repeticoes, esperado)
	})

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}