package zerovalue

import (
	"fmt"
)

// Quando não há um valor. Sempre será retornado esses
func ValueExemplo1() {
	var i int     // 0
	var f float64 // 0
	var b bool    // false
	var s string  // ""
	fmt.Printf("Inteiro: %v\nFloat: %v\nBool: %v\nString: %q\n", i, f, b, s)
}
