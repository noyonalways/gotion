package ui

import "fmt"

func (m Model) View() string {
	welcome := WelcomeStyle.Render("Welcome to Gotion! 🧠")

	helpText := "Ctrl+N: new file | Ctrl+L: list | Esc: back | Ctrl+S: save | Ctrl+E: export | Ctrl+C/Q: quit"
	if m.showingList {
		helpText = "Enter: open | Ctrl+D: delete | Esc: back | Ctrl+N: new | Ctrl+E: export | Ctrl+C/Q: quit"
	}
	helpKeys := HelpKeysStyle.Render(helpText)

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

	exportMsg := ""
	if m.exportMessage != "" {
		exportMsg = fmt.Sprintf("\n%s", m.exportMessage)
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s%s", welcome, view, helpKeys, exportMsg)
}
