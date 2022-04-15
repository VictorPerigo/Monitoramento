package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibirMenu()
}

var urls []string = []string{
	"https://www.google.com.br/notfound",
	"https://www.youtube.com",
}

func menuFuncionamento() {
	switch inputOpcao() {
	case 1: // Iniciar monitoramento
		monitoramento()
	case 2: // Exibir os logs
		fmt.Println("Exibindo os logs...")
	case 3: // Urls
		menuUrls()
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
		os.Exit(-1)
	}
}

func inputOpcao() int {
	opcao := 5
	// variavel opcao recebe o valor digitado pelo usuario, sem expecificar o tipo do input, o proprio go ira inferir o tipo apartir da variavel opcao
	fmt.Scan(&opcao)

	return opcao
}

func monitoramento() {
	fmt.Print("\033[H\033[2J")

	for _, url := range urls {
		response, err := http.Get(url)

		if err != nil {
			fmt.Println("Erro ao acessar a url:", err)
		} else if response.StatusCode == 200 {
			fmt.Println("status code:", response.StatusCode, "OK  Site:", url)
			defer response.Body.Close()
		} else {
			fmt.Println("status code:", response.StatusCode, "ERR Site:", url)
			defer response.Body.Close()
		}
	}

}

func menuUrls() {
	exibirUrls()

	exibirMenuUrl()

	switch inputOpcao() {
	case 1: // Adicionar url
		adicionarUrl()
	case 2: // Remover url
		removerUrl()
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		fmt.Println("\033[H\033[2J")
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
	}
}

func adicionarUrl() {
	exibirUrls()
	fmt.Println("Digite a url:")
	var url string
	fmt.Scan(&url)
	urls = append(urls, url)
	fmt.Println("Url adicionada com sucesso!")
}

func removerUrl() {
	exibirUrls()
	fmt.Println("\nDigite o index da a url que deseja remover:")
	var indexEscolhido int
	fmt.Scan(&indexEscolhido)
	for index := range urls {
		if index == indexEscolhido {
			urls = append(urls[:index], urls[index+1:]...)
			fmt.Println("Url removida com sucesso!")
			break
		}
	}
}

// exibicoes:

func exibirUrls() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Urls:")
	for index, url := range urls {
		fmt.Println("index:", index, ", url:", url)
	}
}

func exibirMenuUrl() {
	fmt.Println("\nOque deseja fazer?")
	fmt.Println("1 - Adicionar url")
	fmt.Println("2 - Remover url")
	fmt.Println("0 - Voltar ao menu:")
}

func exibirMenu() {
	for {
		fmt.Println("\nOque deseja fazer?")
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Exibir os logs")
		fmt.Println("3 - Urls")
		fmt.Println("0 - Sair do programa:")
		menuFuncionamento()
	}
}
