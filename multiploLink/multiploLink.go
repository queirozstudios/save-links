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

	_, err := os.Stat("links.json")
	var data LinksData

	if err == nil {

		file, _ := ioutil.ReadFile("links.json")
		json.Unmarshal(file, &data)
	}

	fmt.Println("Digite os links (um por linha). Após terminar, pressione Enter duas vezes:")
	scanner := bufio.NewScanner(os.Stdin)
	var links []string
	for {
		scanner.Scan()
		link := scanner.Text()
		if link == "" {
			break
		}
		links = append(links, link)
	}

	for _, link := range links {
		fmt.Printf("Digite a descrição para o link '%s': ", link)
		scanner.Scan()
		descricao := scanner.Text()

		novaLinks := Link{Nome: descricao, Link: link}

		if !verificarLinkExistente(data.Links, novaLinks.Nome, novaLinks.Link) {

			data.Links = append(data.Links, novaLinks)
			fmt.Println("✅ Link adicionado com sucesso!")
		} else {
			fmt.Println("❌ Link já existe.")
		}
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")

	err = ioutil.WriteFile("links.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
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
