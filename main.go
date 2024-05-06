// AULA 9 - Pacotes e Visibilidade | GO - Orientado a Objetos

package main

import (
	"banco/contas" // Pacote personalizado
	// c	"banco/contas" // Pacote personalizado com apelido "c"
	"fmt"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	fmt.Println(contaDaSilvia, contaDoGustavo)

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	//	status := contaDaSilvia.Transferir(-300, &contaDoGustavo)

	fmt.Println(status)

	fmt.Println(contaDaSilvia, contaDoGustavo)

}
