package ifelse

//importações usadas
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Chamando as funções

func IfElse() {
	condicao1()
}

// Aqui é meus códigos básicos sobre if e else em golang
// Uma estrutura de condição como outras linguagens.
func condicao1() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite um número maior ou menor que 5: ")
	scanner.Scan()
	numero, _ := strconv.Atoi(scanner.Text())

	valor := numero
	//if (condição {<ação>}
	if valor < 5 { //true
		fmt.Println("O valor é menor que 5.")
	} else if valor > 5 {
		fmt.Println("O valor é maior que 5.")
	} else {
		fmt.Println("O valor é 5")
	}
}

//operações +, -, /, *, %
