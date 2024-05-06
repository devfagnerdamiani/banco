// AULA 5 - Comparando Tipos | GO - Orientado a Objetos

package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroDaConta int
	saldo         float64
}

func main() {
	// contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroDaConta: 123456, saldo: 125.5}
	// contaDoGuilherme2 := ContaCorrente{titular: "Guilherme", numeroAgencia: 580, numeroDaConta: 123456, saldo: 125.5}
	//	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}
	//contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200.0}
	//contaDaBruna2 := ContaCorrente{"Bruna", 222, 111111, 200.0}

	// fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDoGuilherme2)
	//fmt.Println(contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris)

	fmt.Println(&contaDaCris)
	fmt.Println(&contaDaCris)

	fmt.Println(*contaDaCris)

	//fmt.Println(contaDoGuilherme == contaDoGuilherme2)
	//fmt.Println(contaDaBruna == contaDaBruna2)

	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)

}
