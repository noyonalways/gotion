package ui

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor = lipgloss.Color("51")
	BlackColor   = lipgloss.Color("16")
	WhiteColor   = lipgloss.Color("255")
	OffGrayColor = lipgloss.Color("250")

	TextStyle   = lipgloss.NewStyle().Foreground(PrimaryColor)
	CursorStyle = lipgloss.NewStyle().Foreground(PrimaryColor)
	PromptStyle = lipgloss.NewStyle().Foreground(PrimaryColor)
	DocStyle    = lipgloss.NewStyle().Margin(1, 2)

	WelcomeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(BlackColor).
			Background(PrimaryColor).
			Padding(0, 2)

	HelpKeysStyle = lipgloss.NewStyle().
			Foreground(OffGrayColor)

	ListTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(BlackColor).
			Background(PrimaryColor).
			Padding(0, 1)
)
