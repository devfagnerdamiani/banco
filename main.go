// AULA 4 - New e Ponteiros | GO - Orientado a Objetos

package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroDaConta int
	saldo         float64
}

func main() {
	//	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroDaConta: 123456, saldo: 125.5}
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200.0}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	fmt.Println(*contaDaCris)

}
