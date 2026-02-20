package tui

import (
	"number-guessing-game/internal/controller"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	controller controller.Controller
	textinput  textinput.Model
	choices    []string
	cursor     int
	isGameOn   bool
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "15"
	ti.Focus()
	ti.Width = 2

	return model{
		textinput: ti,
		choices:   []string{"1.Easy (10 chances)\n", "2.Medium (5 chances)\n", "3.Hard (3 chances)\n"},
		cursor:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
