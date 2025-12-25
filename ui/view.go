package ui

import "fmt"

func (m Model) View() string {
	welcome := WelcomeStyle.Render("Welcome to Gotion! 🧠")
	helpKeys := HelpKeysStyle.Render("Ctrl+N: new file | Ctrl+L: list | Esc: back | Ctrl+S: save | Ctrl+C/Q: quit")

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.noteTextArea.View()
	}

	// show list if showingList is true
	if m.showingList {
		view = m.list.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, helpKeys)
}
