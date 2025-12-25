package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/noyonalways/gotion/storage"
	"github.com/noyonalways/gotion/ui"
)

const Version = "1.0.0" // Increment this when you update

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("gotion version: %s\n", Version)
		return
	}

	// Initialize storage (creates vault directory if needed)
	storage.Init()

	p := tea.NewProgram(ui.NewModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
