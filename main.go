// AULA 6 - Criando o método sacar | GO - Orientado a Objetos

package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroDaConta int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {

		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}

}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.sacar(800.0))
	fmt.Println(contaDaSilvia.sacar(-100.0))
	fmt.Println(contaDaSilvia.sacar(200.0))
	fmt.Println(contaDaSilvia.sacar(400.0))
	fmt.Println(contaDaSilvia.saldo)
}
