package loops

import "fmt"

//loops, laços de repetição e repetir tarefas

func Lop1() {
	repeticao()
	repeticao2()
}

// estrutura

func repeticao() {
	sum := 0

	//i++ --> soma 1
	//i-- --> subtraindo 1

	//sintaxe for em golang

	for i := 0; i < 10; i++ {
		fmt.Println("Sum:", sum)
		fmt.Println("Indice:", i)
		sum += i
		//mesma coisa que sum = sum + 1
		//sum -= / sum = sum - 1
	}
	fmt.Println(sum)
}

//for range pecorre listas que vai pegar as posições e printar os valores da lista

func repeticao2() {
	frutas := []string{"Banana", "Maçã", "Morango", "Uva", "Cereja"}
	for _, fruit := range frutas {
		fmt.Println(fruit)
	}
}
