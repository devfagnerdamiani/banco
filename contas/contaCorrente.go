package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular                      clientes.Titular
	NumeroAgencia, NumeroDaConta int
	saldo                        float64 // Visualização de saldo alterado apenas para leitura interna deste arquivo
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {

		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente!"
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

func (c *ContaCorrente) ObterSaldo() float64 { // Inserir com Letra maiuscula para dar acesso a outros arquivos
	return c.saldo
}
