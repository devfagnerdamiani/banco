package contas

import (
	"banco/clientes"
)

type ContaPoupanca struct {
	Titular                                clientes.Titular
	NumeroAgencia, NumeroDaConta, Operacao int
	saldo                                  float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {

		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {

		c.saldo += valorDoDeposito

		return "Deposito realizado com sucesso! Valor atual do saldo: ", c.saldo

	} else {

		return "Deposito menor que zero. Valor atual do saldo: ", c.saldo

	}

}

func (c *ContaPoupanca) ObterSaldo() float64 { // Inserir com Letra maiuscula para dar acesso a outros arquivos
	return c.saldo
}
