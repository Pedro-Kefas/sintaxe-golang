package arrys

import "fmt"

//LISTAS em array fixo

// chamando as funções criadas

func Array() {
	fmt.Println(arrayTamanhoFixo())
	arrayTamanhoFixo2()
	arrayTamanhoFixo3()
	arrayTamanhoFixo4()
	arrayTamanhoFixo5()
	arrayInt1()
	arrayInt2()
}

// array tamanho fixo
func arrayTamanhoFixo() [2]string {
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	return array
}

// array tamanho fixo 2
func arrayTamanhoFixo2() [3]string {
	var array2 [3]string
	array2[0] = "Olá"
	array2[1] = "Mundo"
	array2[2] = "123"

	//esse método string só pode colocar 1 argumento, mas posso concatenar usando + para colocar um segundo argumento ou mais
	//Seguindo essa sintaxe eu posso imprimir vários Indices e seus Indices
	//Vem printado dentro da lista
	primeiraLetra := string(array2[0][1]) + string(array2[1][0]) + string(array2[2][2])
	fmt.Println("Primeira letra dos Indices da Array: ", primeiraLetra)

	return array2
}

// array tamanho fixo 3
func arrayTamanhoFixo3() [2]string {
	var array3 [2]string
	array3[0] = "Peter"
	array3[1] = "Pan"
	arrayResultado := array3

	//Esse método é o normal onde vou printar somente os Indices da Array, sem querer puxar também as Letras dos indices
	//Vem printado dentro da lista
	fmt.Println("Aqui está a Array:", arrayResultado)

	return array3
}

// array tamanho fixo 4
func arrayTamanhoFixo4() [2]string {
	var array4 [2]string
	array4[0] = "Sonic"
	array4[1] = "Ouriço"

	fmt.Printf("Aqui vai retornar as strings sem Array/lista o nome %s O %s\n", array4[0], array4[1])

	return array4
}

// array tamanho fixo 5
func arrayTamanhoFixo5() [2]string {
	var array5 [2]string
	array5[0] = "Com"
	array5[1] = "Lista"

	fmt.Println("Aqui vai retornar com lista/array:", array5)
	return array5
}

// array tamanho fixo em int
func arrayInt1() [5]int {
	//Também posso atribuir passando numa única linha sem precisar fazer array[] = ""
	arrayNumber := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Atribuição direta em uma linha:", arrayNumber)
	fmt.Println(arrayNumber[1:4])
	//: 1 pega tudo que está antes 1: pega tudo que está depois

	return arrayNumber
}

// array tamanho fixo em int 2
func arrayInt2() [2]int {
	var arrayNumber2 = [2]int{10, 20}
	fmt.Println("Essa lista nova é:", arrayNumber2[0], arrayNumber2[1])
	fmt.Printf("A soma de %d e %d é: %d\n", arrayNumber2[0], arrayNumber2[1], arrayNumber2[0]+arrayNumber2[1])
	return arrayNumber2
}
