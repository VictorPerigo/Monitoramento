package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	exibirMenu()
}

var urls []string = getUrlsFromTxt()

// funcionamento dos menus:

func menuFuncionamento() {
	switch inputOpcao() {
	case 1: // Iniciar monitoramento
		monitoramento()
	case 2: // Exibir os logs
		fmt.Println("Exibindo os logs...")
	case 3: // Urls
		menuFuncionamentoUrls()
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
		os.Exit(-1)
	}
}

func menuFuncionamentoUrls() {
	exibirUrls()
	exibirMenuUrl()

	switch inputOpcao() {
	case 1: // Adicionar url
		_ = adicionarUrl()
	case 2: // Remover url
		removerUrl()
	case 0: // Sair do programa
		fmt.Println("Saindo do programa...")
		fmt.Println("\033[H\033[2J")
	default: // Caso o usuario digite uma opcao invalida
		fmt.Println("Opção invalida")
	}
}

// exibicoes:

func exibirUrls() {
	limparTerminal()
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

// funcionamento das urls:

func adicionarUrl() error {
	file, err := os.OpenFile("urls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return err
	}
	defer file.Close()
	if _, err = file.WriteString("jon\n"); err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return err
	}
	return nil
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

func getUrlsFromTxt() []string {
	var sites []string
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		os.Exit(-1)
	} else {
		leitor := bufio.NewReader(file)

		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.TrimSpace(linha)
			sites = append(sites, linha)
			if err == io.EOF {
				defer file.Close()
				break
			}

		}
	}

	return sites
}

// funcionamento das opcoes:

func monitoramento() {
	limparTerminal()

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

// ferramentas:

func inputOpcao() int {
	opcao := 999
	// variavel opcao recebe o valor digitado pelo usuario, sem expecificar o tipo do input, o proprio go ira inferir o tipo apartir da variavel opcao
	fmt.Scan(&opcao)

	return opcao
}

func limparTerminal() {
	fmt.Print("\033[H\033[2J")
}
