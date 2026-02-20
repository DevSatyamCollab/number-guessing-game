package tui

import (
	"number-guessing-game/internal/controller"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "down", "j": // down
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "up", "k": // up
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			if !m.isGameOn {
				m.isGameOn = true
				m.controller = *controller.NewController(m.cursor)
				return m, textinput.Blink
			}
		}
	}

	return m, nil
}
