package maps

import "fmt"

//Chamando funções

func Mapa1() {
	mapaInfo()
}

//sintaxe básica do map sendo usada para revelar o valor que as strings estão armazenando

func mapaInfo() {
	anoNascimento := map[string]int{
		"Steve": 1999,
		"Alex":  2000,
	}

	fmt.Println(anoNascimento)
	fmt.Println(anoNascimento["Steve"])
	fmt.Println(anoNascimento["Alex"])

	//adicionado uma nova informação diferente de append

	anoNascimento["Herobrine"] = 0

	fmt.Println(anoNascimento["Herobrine"])
}
