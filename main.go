package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vaultDir    string
	cursorStyle = lipgloss.NewStyle().Foreground((lipgloss.Color("45")))
)

// go lang init file
func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory:", err)
		os.Exit(1)
	}
	vaultDir = fmt.Sprintf("%s/.gotion", homeDir)
}

type model struct {
	newFileInput           textinput.Model
	noteTextArea           textarea.Model
	createFileInputVisible bool
	currentFile            *os.File
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			// handle new file input
			m.createFileInputVisible = true
			return m, nil

		case "ctrl+s":
			// text are value -> write in it that file descriptor and close it
			if m.currentFile == nil {
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("can not save the file :(")
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("can not save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.WriteString(m.noteTextArea.Value()); err != nil {
				fmt.Println("can not save the file :(")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("can not close the file :(")
			}

			m.currentFile = nil
			m.noteTextArea.SetValue("")

			return m, nil

		case "enter":
			// todo : create file
			filename := m.newFileInput.Value()
			if filename != "" {
				filepath := fmt.Sprintf("%s/%s.md", vaultDir, filename)

				// check if file already exists
				if _, err := os.Stat(filepath); err == nil {
					return m, nil
				}

				// create file
				f, err := os.Create(filepath)
				if err != nil {
					log.Fatalf("Error creating file %s: %v", filepath, err)
					os.Exit(1)
				}
				defer f.Close()

				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")
			}
		}
	}

	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	if m.currentFile != nil {
		m.noteTextArea, cmd = m.noteTextArea.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var welcomeStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("45")).
		Padding(0, 2)

	var helpKeysStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("250"))

	welcome := welcomeStyle.Render("Welcome to Gotion! 🧠")
	helpKeys := helpKeysStyle.Render("Ctrl+N: new file | Ctrl+L: list | Esc: back/save | Ctrl+S: save | Ctrl+C/Q: quit")

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.noteTextArea.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, helpKeys)
}

func initializeModel() model {

	err := os.MkdirAll(vaultDir, 0750)
	if err != nil {
		log.Fatal("Error creating vault directory:", err)
		os.Exit(1)
	}

	// initialize nw file input
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = cursorStyle

	// initialize note text area
	ta := textarea.New()
	ta.Placeholder = "Start writing your note..."
	ta.Focus()
	ta.ShowLineNumbers = false

	return model{
		newFileInput:           ti,
		noteTextArea:           ta,
		createFileInputVisible: false,
	}
}

func main() {
	p := tea.NewProgram(initializeModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
