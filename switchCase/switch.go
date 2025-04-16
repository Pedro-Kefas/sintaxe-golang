package switchCase

import "fmt"

//Chamando as funções

func Swicaso() {
	casosSwitch()
}

// switch cases exemplos

func casosSwitch() {
	posicao := 3

	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar!")
	case 2:
		fmt.Println("Segundo lugar!")
	case 3:
		fmt.Println("Terceiro lugar!")
	}
}
