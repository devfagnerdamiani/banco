// AULA 11 - Tipos aninhados | GO - Orientado a Objetos

// go mod init // ########## CÓDIGO PARA CRIAR MÓDULOS NA PASTA RAÍZ ( Onde esta o arquivo main.go )

package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

// Pacote personalizado
// c	"banco/contas" // Pacote personalizado com apelido "c"

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroDaConta: 123456, Saldo: 100}
	fmt.Println(contaDoBruno)
}
