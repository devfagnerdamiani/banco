// AULA 13 - Conta Poupanca | GO - Orientado a Objetos

// go mod init // ########## CÓDIGO PARA CRIAR MÓDULOS NA PASTA RAÍZ ( Onde esta o arquivo main.go )

package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaPoupanca{}

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaPati)

}
