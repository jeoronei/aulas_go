package main

import (
	"io"
	"fmt"
	"os"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"

type Sleeper interface {
	Pausa()
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Pausa() {
	s.Chamadas++
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa func(time.Duration)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
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
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}