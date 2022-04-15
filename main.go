package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá, que deseja fazer?")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")

	var opcao int
	fmt.Scanf("%d", &opcao)

	switch opcao {
	case 1: // Iniciar monitoramento
		fmt.Println("Monitorando...")
	case 2: // Exibir os logs
		fmt.Println("Exibindo os logs...")
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
	}

}
