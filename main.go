// AULA 12 - Alterando a visibilidade | GO - Orientado a Objetos

// go mod init // ########## CÓDIGO PARA CRIAR MÓDULOS NA PASTA RAÍZ ( Onde esta o arquivo main.go )

package main

import (
	"banco/contas"
	"fmt"
)

// Pacote personalizado
// c	"banco/contas" // Pacote personalizado com apelido "c"

func main() {
	contaDeExemplo := contas.ContaCorrente{}
	contaDeExemplo.Depositar(-100) // não permitir que possa ser setado o valor de saldo diretamente

	fmt.Println(contaDeExemplo.ObterSaldo())
}
