package struc

import "fmt"

// structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade

func Structs() {
	struct1()
	listaDePessoas()
}

// criando um tipo com struct

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func struct1() {
	var dadosPessoais = Pessoa{Nome: "Steve", Sobrenome: "Quadrado", Idade: 14}

	fmt.Println(Pessoa{Nome: "Alex", Sobrenome: "Quadrado", Idade: 15})
	fmt.Printf("O nome dele é %s, seu sobrenome é %s e sua idade é %d.\nO tipo é %T.\n", dadosPessoais.Nome, dadosPessoais.Sobrenome, dadosPessoais.Idade, dadosPessoais)
}

type Pessoas struct {
	Nome2  string
	Idade2 int
}

func listaDePessoas() Pessoas {
	p1 := Pessoas{Nome2: "Alex", Idade2: 33}
	p2 := Pessoas{Nome2: "Steve", Idade2: 35}
	p3 := Pessoas{Nome2: "Herobrine", Idade2: 1000}

	//go transforma essa estrutura em um tipo novo

	listaDeUsuarios := []Pessoas{p1, p2} //Aqui é uma lista slice
	listaAtualizada := append(listaDeUsuarios, p3)
	fmt.Printf("Primeiro usuário:\nNome: %s.\nIdade: %d.\n", listaAtualizada[0].Nome2, listaAtualizada[0].Idade2)
	fmt.Printf("Segundo usuário:\nNome: %s.\nIdade: %d.\n", listaAtualizada[1].Nome2, listaAtualizada[1].Idade2)
	fmt.Printf("Terceiro usuário:\nNome: %s.\nIdade: %d.\n", listaAtualizada[2].Nome2, listaAtualizada[2].Idade2)

	return Pessoas{}
}
