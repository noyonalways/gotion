package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/noyonalways/gotion/storage"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-5)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}
			if m.currentFile != nil {
				m.noteTextArea.SetValue("")
				m.currentFile = nil
			}

			if m.showingList {
				if m.list.FilterState() == list.Filtering {
					break
				}
				m.showingList = false
			}
			return m, nil
		case "ctrl+l":
			noteList := getNoteList()
			m.list.SetItems(noteList)
			m.showingList = true
			return m, nil
		case "ctrl+n":
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
			if m.currentFile != nil {
				break
			}

			if m.showingList {
				selectedItem, ok := m.list.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", storage.GetVaultDir(), selectedItem.title)
					content, err := os.ReadFile(filepath)
					if err != nil {
						log.Printf("Error reading file %s: %v", filepath, err)
						return m, nil
					}
					m.noteTextArea.SetValue(string(content))

					f, rr := os.OpenFile(filepath, os.O_RDWR, 0644)
					if rr != nil {
						log.Printf("Error opening file %s: %v", filepath, rr)
						return m, nil
					}
					m.currentFile = f
					m.showingList = false
				}

				return m, nil
			}

			// todo : create file
			filename := m.newFileInput.Value()
			if filename != "" {
				filepath := fmt.Sprintf("%s/%s.md", storage.GetVaultDir(), filename)

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

	if m.showingList {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}
