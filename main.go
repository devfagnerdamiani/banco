// AULA 14 - Interface | GO - Orientado a Objetos

// go mod init // ########## CÓDIGO PARA CRIAR MÓDULOS NA PASTA RAÍZ ( Onde esta o arquivo main.go )

package main

import (
	"banco/contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {

	conta.Sacar(valorDoBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDaLuisa := contas.ContaCorrente{}

	contaDoDenis.Depositar(100)
	pagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())
	//	fmt.Println(contaDoDenis)
	// contaDoDenis.Sacar(5000) // Valor de teste a não ser sacado por conta do valor a sacar é superior ao valor do saldo
	// contaDoDenis.Sacar(55)
	contaDaLuisa.Depositar(500)
	pagarBoleto(&contaDaLuisa, 400)
	fmt.Println(contaDaLuisa.ObterSaldo())

}
