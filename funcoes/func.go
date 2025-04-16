package funcoes

import (
	"fmt"
	"strconv"
)

//chamando funções e printando elas no console

func Funcoes() {
	fmt.Println(soma(10, 5), subtracao(10, 5), divisao(10, 2), multiplicacao(9, 2))
	fmt.Println(nomeNaPrint("Rússia"))
	fmt.Println(transformarStringEmInt("10", 10))
	fmt.Println(maisTipos("Maria"))
	fmt.Println(nomeCompleto("Troika", "Idiomas"))
}

// Exemplos de funções com operações matemáticas e return
func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func divisao(x int, y int) int {
	return x / y
}

func multiplicacao(x int, y int) int {
	return x * y
}

// Exemplos de funções com strings
func nomeNaPrint(nome string) string {
	return nome
}

//trocando os tipos

func transformarStringEmInt(numeroEscrito string, p int) int {
	conversor, _ := strconv.Atoi(numeroEscrito) //convertendo
	return conversor + p
}

func maisTipos(primeiroTipo string) (string, string) {
	nome := "golang"
	return primeiroTipo, nome
}

func nomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}
