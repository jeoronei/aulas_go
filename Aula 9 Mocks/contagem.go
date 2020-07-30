package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3

type Sleeper interface {
	Pausa()
}

type SleeperPadrao struct {}

func(d *SleeperPadrao) Pausa() {
	time.Sleep(1 * time.Second)
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Pausa() {
	s.Chamadas++
}

func Contagem(saida io.Writer, sleep Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleep.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleep.Pausa()
	fmt.Fprintf(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}