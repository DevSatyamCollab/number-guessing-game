package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	isGameOn bool
}

func InitialModel() model {
	return model{
		choices: []string{"1.Easy (10 chances)\n", "2.Medium (5 chances)\n", "3.Hard (3 chances)\n"},
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
