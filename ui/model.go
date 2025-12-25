package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/noyonalways/gotion/storage"
)

type Model struct {
	newFileInput           textinput.Model
	noteTextArea           textarea.Model
	createFileInputVisible bool
	currentFile            *os.File
	list                   list.Model
	showingList            bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func NewModel() Model {
	// initialize nw file input
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = CursorStyle
	ti.PromptStyle = PromptStyle
	ti.TextStyle = TextStyle

	// initialize note text area
	ta := textarea.New()
	ta.Placeholder = "Start writing your note..."
	ta.Focus()
	ta.FocusedStyle.Prompt = PromptStyle
	ta.ShowLineNumbers = true

	// initialize list
	noteList := getNoteList()
	finalList := list.New(noteList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All notes 📝"
	finalList.Styles.Title = ListTitleStyle

	return Model{
		newFileInput:           ti,
		noteTextArea:           ta,
		createFileInputVisible: false,
		list:                   finalList,
		showingList:            false,
	}
}

func getNoteList() []list.Item {
	items := make([]list.Item, 0)

	files, err := storage.ListFiles()
	if err != nil {
		return items
	}

	for _, file := range files {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				continue
			}

			modTime := info.ModTime().Format("2006-01-02 15:04:05")
			items = append(items, item{
				title: file.Name(),
				desc:  fmt.Sprintf("Modified: %s", modTime),
			})
		}
	}

	return items
}
