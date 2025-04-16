package variaveis

//importações usadas
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Usando o Scanner importado para ler e depois strconv para converter string em int

func Input1() {
	scanner := bufio.NewScanner(os.Stdin)

	// Lendo nome

	fmt.Print("Digite seu nome: ")
	scanner.Scan()
	nome := scanner.Text()

	//Lendo idade

	fmt.Print("Escreva sua idade: ")
	scanner.Scan()
	idade, _ := strconv.Atoi(scanner.Text()) // Convertendo para int

	fmt.Println("Seu nome é:", nome, "\nSua idade é:", idade+10)
}

// Nota de usuário usando as importações

func InputNota() {
	notaScanInput := bufio.NewScanner(os.Stdin) // bufio.NewScanner

	//Primeira nota

	fmt.Print("Digite sua primeira nota: ")
	notaScanInput.Scan()
	nota1, _ := strconv.Atoi(notaScanInput.Text()) //Convertendo para int

	//Segunda nota

	fmt.Print("Digite sua segunda nota: ")
	notaScanInput.Scan()
	nota2, _ := strconv.Atoi(notaScanInput.Text()) //Convertendo para int

	//Resultado da soma

	fmt.Println("A primeira nota é:", nota1, "a segunda nota é:", nota2, "A soma delas é:", nota1+nota2)
}
