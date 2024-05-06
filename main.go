// AULA 7 - Mútiplos retornos | GO - Orientado a Objetos

package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroDaConta int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {

		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {

		c.saldo += valorDoDeposito

		return "Deposito realizado com sucesso! Valor atual do saldo: ", c.saldo

	} else {

		return "Deposito menor que zero. Valor atual do saldo: ", c.saldo

	}

}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(800.0))
	fmt.Println(contaDaSilvia.Sacar(-100.0))
	fmt.Println(contaDaSilvia.Sacar(200.0))
	fmt.Println(contaDaSilvia.Sacar(400.0))
	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Depositar(2000))
	fmt.Println(contaDaSilvia.Depositar(-2000))
	fmt.Println(contaDaSilvia.saldo)

	status, valor := contaDaSilvia.Depositar(30)
	fmt.Println("Status: ", status, " | Valor: ", valor)

}
