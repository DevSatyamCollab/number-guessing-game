package style

import "github.com/charmbracelet/lipgloss"

type StyleBundle struct {
	Cursor     lipgloss.Style
	FailedRes  lipgloss.Style
	CorrectRes lipgloss.Style
	Quit       lipgloss.Style
}

func DefaultStyle() StyleBundle {
	return StyleBundle{
		Cursor: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF06B7")),

		FailedRes: lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")),

		CorrectRes: lipgloss.NewStyle().
			Foreground(lipgloss.Color("82")),

		Quit: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F5F")),
	}
}
