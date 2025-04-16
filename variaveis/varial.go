package variaveis

import "fmt"

// Exemplos de variáveis

func Variaveis() {

	//Variáveis tipo var
	//var + NOME DA VARIÁVEL + TIPO | Usado dentro de funções e fora do escopo delas a nível de arquivo

	var nome string = "Elliot"
	var idade int = 27
	fmt.Printf("O Hacker %s tem %d anos de idade.\n", nome, idade)

	//Posso reatribuir um no VALUE para minhas variáveis usando "="

	nome = "Pedro"
	idade = 22
	fmt.Printf("O hacker %s tem %d anos de idade.\n", nome, idade)
}

func Variaveis2() {

	//Variáveis tipo idiomático. Usado apenas dentro de funções
	// Go entende que "nome" é do tipo string ou int, ou qualquer outro tipo

	nomeDaSerie := "Mr.Robot"
	fmt.Println(nomeDaSerie)
	fmt.Printf("Type: %T - Value: %v\n", nomeDaSerie, nomeDaSerie)

	//Outro exemplo com int

	cursoRaiamSantos := "Nômade Digital"
	criador := "Raimam Santos"
	idade := 22
	fmt.Printf("Esse curso de marketing digital se chama: %s\nCriado por %s aos %d anos.\n", cursoRaiamSantos, criador, idade)
}
