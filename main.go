package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/noyonalways/gotion/storage"
	"github.com/noyonalways/gotion/ui"
)

// Changed from 'const' to 'var' so ldflags can overwrite it
var Version = "dev"

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("gotion version: %s\n", Version)
		return
	}

	storage.Init()

	p := tea.NewProgram(ui.NewModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
