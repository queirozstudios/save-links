package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Link struct {
	ID   int    `json:"id"`
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

	// Inicializa o ID que será atribuído ao novo link
	nextID := len(data.Links) + 1

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

	// Insere os links com auto incremento de ID
	for _, link := range links {
		novaLink := Link{
			ID:   nextID,
			Link: link,
		}
		// Verifica se o link já existe
		if !verificarLinkExistente(data.Links, novaLink.Link) {
			data.Links = append(data.Links, novaLink)
			fmt.Printf("✅ Link com ID %d adicionado com sucesso!\n", novaLink.ID)
			nextID++ // Incrementa o ID para o próximo link
		} else {
			fmt.Println("❌ Link já existe.")
		}
	}

	// Escreve os dados atualizados de volta ao arquivo links.json
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	err = ioutil.WriteFile("links.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
}

func verificarLinkExistente(Links []Link, link string) bool {
	for _, emp := range Links {
		if emp.Link == link {
			return true
		}
	}
	return false
}
