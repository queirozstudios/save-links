package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Link struct {
	Nome string `json:"nome"`
	Link string `json:"link"`
}

type LinksData struct {
	Links []Link `json:"Links"`
}

func main() {
	// Verificar se o arquivo existe
	_, err := os.Stat("links.json")
	var data LinksData

	if err == nil {
		// Se o arquivo existir, ler as informações existentes
		file, _ := ioutil.ReadFile("links.json")
		json.Unmarshal(file, &data)
	}

	// Obter informações da nova Link
	var novaLinks Link
	fmt.Print("Descrição do Link: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	novaLinks.Nome = scanner.Text()

	fmt.Print("Digite o link: ")
	scanner.Scan()
	novaLinks.Link = scanner.Text()

	// Verificar se já existe o nome ou link
	if !verificarLinkExistente(data.Links, novaLinks.Nome, novaLinks.Link) {
		// Adicionar nova Link à lista
		data.Links = append(data.Links, novaLinks)

		// Converter a estrutura para JSON
		jsonData, _ := json.MarshalIndent(data, "", "  ")

		// Escrever o JSON no arquivo
		err = ioutil.WriteFile("links.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}

		fmt.Println("✅ Link adicionado com sucesso!")
	} else {
		fmt.Println("❌ Link já existe.")
	}
}

func verificarLinkExistente(Links []Link, nome, link string) bool {
	for _, emp := range Links {
		if emp.Nome == nome || emp.Link == link {
			return true
		}
	}
	return false
}
