// AULA 1 - Variáveis e Tipos | GO - Orientado a Objetos

package main

import "fmt"

func main() {

	var titular string = "Guilherme"
	var numeroAgencia int = 589
	var numeroDaConta int = 123456
	var saldo float64 = 125.5

	fmt.Println("Titular: ", titular, " | Número da Agência: ", numeroAgencia, " | Número da Conta: ", numeroDaConta, " | Saldo: ", saldo)

	var titular2 string = "Luciene"
	var numeroAgencia2 int = 555
	var numeroDaConta2 int = 111333
	var saldo2 float64 = 200.0

	fmt.Println("Titular: ", titular2, " | Número da Agência: ", numeroAgencia2, " | Número da Conta: ", numeroDaConta2, " | Saldo: ", saldo2)

}
