// AULA 8 - Transferencia entre contas | GO - Orientado a Objetos

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

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {

		return false

	}

}

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	fmt.Println(contaDaSilvia, contaDoGustavo)

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	//	status := contaDaSilvia.Transferir(-300, &contaDoGustavo)

	fmt.Println(status)

	fmt.Println(contaDaSilvia, contaDoGustavo)

}
