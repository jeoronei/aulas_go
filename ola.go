package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) string {
	switch idioma {
	case frances:
		return prefixoOlaFrances
	case espanhol:
		return prefixoOlaEspanhol
	default:
		return prefixoOlaPortugues
	}
}

func main() {
	fmt.Println(Ola("Mundo", ""))
}