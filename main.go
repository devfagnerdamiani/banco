// AULA 2 - STRUCTS | GO - Orientado a Objetos

package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroDaConta int
	Saldo         float64
}

func main() {

	fmt.Println(ContaCorrente{})

}
