package styles

import "github.com/charmbracelet/lipgloss"

var StartBlock = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("228")).
	BorderBottom(false).
	BorderLeft(true).
	BorderRight(true).
	BorderTop(false).
	Foreground(lipgloss.Color("#33FF33")).
	PaddingLeft(2).
	Width(50)

var StartBlockTop = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("228")).
	BorderLeft(true).
	BorderRight(true).
	BorderTop(true).
	Foreground(lipgloss.Color("#33FF33")).
	PaddingLeft(2).
	Width(50)

var StartBlockBottom = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("228")).
	BorderBottom(true).
	BorderLeft(true).
	BorderRight(true).
	Foreground(lipgloss.Color("#33FF33")).
	PaddingLeft(2).
	Width(50)
