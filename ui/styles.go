package ui

import "github.com/charmbracelet/lipgloss"

var (
	TextStyle   = lipgloss.NewStyle().Foreground((lipgloss.Color("45")))
	CursorStyle = lipgloss.NewStyle().Foreground((lipgloss.Color("45")))
	PromptStyle = lipgloss.NewStyle().Foreground((lipgloss.Color("45")))
	DocStyle    = lipgloss.NewStyle().Margin(1, 2)

	WelcomeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("16")).
			Background(lipgloss.Color("45")).
			Padding(0, 2)

	HelpKeysStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250"))

	ListTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("16")).
			Background(lipgloss.Color("45")).
			Padding(0, 1)
)
