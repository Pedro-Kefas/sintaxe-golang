package slice

import "fmt"

// chamando as funções

func Slice() {
	slice()
	sliceAppend()
}

//Slices são usados para trabalhar sem tamanho fixo diferente do array fixo

func slice() ([]string, []int) {
	//usamos make para criar slice do zero
	slice := make([]string, 5) //[] Esses colchetes é um argumento e a tipagem dele e o segundo argumento é a quantidade
	slice[0] = "Oi"
	slice[1] = "Maria"

	fmt.Println(slice[0], slice[1]) // Aqui vai printar os indices sem os colchetes (vai sobrar espaço porque eu apliquei um valor mínimo de 5)
	fmt.Println(slice)              // vai printar o slice com os colchetes (vai sobrar espaço porque eu apliquei um valor mínimo de 5)

	//exemplo com int
	sliceInt := make([]int, 5)
	sliceInt[0] = 22
	sliceInt[1] = 33

	//Usando a função len() para descobrir o tamanho do array ou slice
	fmt.Println(sliceInt[0], sliceInt[1])
	fmt.Println("Pegando o que está em 1 em diante:", sliceInt[1:])
	fmt.Println("Usando len para verificar o tamanho:", len(sliceInt))

	//Eles retornam valores vazios: No caso do int é 0. No caso da "string" é vazia
	return slice, sliceInt
}

//Como adicionar mais elementos com a função append

func sliceAppend() ([]string, []string) {
	var adicionar = make([]string, 4)
	adicionar[0] = "Maria"
	adicionar[1] = "Joseph"
	adicionar[2] = "Mario"
	adicionar[3] = "Luigi"
	fmt.Println("Lista antiga de nomes:", adicionar)

	//usando append simplesmente para adicionar mais dados na lista

	novaLista := append(adicionar, "Pedro", "Jazz", "Ghost", "Sasha")
	fmt.Println("Nova lista atualizada:", novaLista)

	return adicionar, novaLista
}
