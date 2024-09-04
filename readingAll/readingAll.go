package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "os"
	"path/filepath"
	"strings"

	"github.com/nsf/termbox-go"
)

type Link struct {
	Nome string `json:"nome"`
	Link string `json:"link"`
}

type LinksData struct {
	Links []Link `json:"Links"`
}

func main() {
	// Inicializa a termbox
	err := termbox.Init()
	if err != nil {
		fmt.Println("Erro ao inicializar o termbox:", err)
		return
	}
	defer termbox.Close()

	directory := "."

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Erro ao ler o diretório:", err)
		return
	}

	var jsonFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			jsonFiles = append(jsonFiles, file.Name())
		}
	}

	if len(jsonFiles) == 0 {
		fmt.Println("Nenhum arquivo JSON encontrado no diretório.")
		return
	}

	drawList := func(selected int) {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		for i, file := range jsonFiles {
			if i == selected {
				termbox.SetCell(0, i, '>', termbox.ColorWhite, termbox.ColorBlue)
			} else {
				termbox.SetCell(0, i, ' ', termbox.ColorWhite, termbox.ColorDefault)
			}
			for j, c := range file {
				termbox.SetCell(j+1, i, c, termbox.ColorWhite, termbox.ColorDefault)
			}
		}
		termbox.Flush()
	}

	selectedIndex := 0
	drawList(selectedIndex)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				if selectedIndex > 0 {
					selectedIndex--
					drawList(selectedIndex)
				}
			case termbox.KeyArrowDown:
				if selectedIndex < len(jsonFiles)-1 {
					selectedIndex++
					drawList(selectedIndex)
				}
			case termbox.KeyEnter:
				selectedFile := jsonFiles[selectedIndex]
				fmt.Printf("Você escolheu o arquivo: %s\n", selectedFile)

				filePath := filepath.Join(directory, selectedFile)
				file, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Println("Erro ao ler o arquivo JSON:", err)
					return
				}

				var data LinksData
				err = json.Unmarshal(file, &data)
				if err != nil {
					fmt.Println("Erro ao decodificar o arquivo JSON:", err)
					return
				}

				fmt.Println("Informações dos Links:")
				for _, link := range data.Links {
					fmt.Println("Nome:", link.Nome)
					fmt.Println("Link:", link.Link)
					fmt.Println("___________________________________")
				}
				return
			case termbox.KeyEsc:
				return
			}
		}
	}
}
