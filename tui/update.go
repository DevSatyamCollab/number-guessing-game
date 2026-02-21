package tui

import (
	"fmt"
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

			if m.isGameOn {
				res, err := m.controller.HandleGuess(m.textinput.Value())
				if err != nil {
					m.result = err.Error() + "\n"
					return m, nil
				}

				currentAttempts, state := m.controller.GetProgess()

				// formated result
				switch state {
				case controller.Won:
					m.isChoiceOn = false
					m.isGameOn = false
					m.result = fmt.Sprintf(
						"Congratulations! You guessed the correct number in %d attempts.\n",
						currentAttempts,
					)

					return m, tea.Quit
				case controller.Lost:
					m.isChoiceOn = false
					m.isGameOn = false
					m.result = fmt.Sprintf(
						"Game Over! ❌ \nSorry, you've run out of attempts.\nThe secret number was %d.\n",
						m.controller.GetSecretNum(),
					)

					return m, tea.Quit
				default:
					m.result = fmt.Sprintf(
						"Incorrect! The number is %s than %s\n",
						res, m.textinput.Value(),
					)
				}

				m.textinput.SetValue("")
			}
		}

		if m.isGameOn {
			var cmd tea.Cmd
			m.textinput, cmd = m.textinput.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}
