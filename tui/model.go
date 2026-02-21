package tui

import (
	"number-guessing-game/internal/controller"
	"number-guessing-game/tui/style"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	controller controller.Controller
	textinput  textinput.Model
	style      style.StyleBundle
	choices    []string
	result     string
	cursor     int
	isGameOn   bool
	isChoiceOn bool
	quit       bool
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "15"
	ti.Focus()
	ti.CharLimit = 2
	ti.Prompt = ""
	ti.Width = 5

	return model{
		textinput:  ti,
		style:      style.DefaultStyle(),
		choices:    []string{"1.Easy (10 chances)\n", "2.Medium (5 chances)\n", "3.Hard (3 chances)\n"},
		cursor:     0,
		isChoiceOn: true,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
