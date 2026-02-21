package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	header := `
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Please select the difficulty level:
`
	footer := "\nPress q for Quit . r new round\n\n"
	var mainView strings.Builder

	if m.isChoiceOn {
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			fmt.Fprintf(&mainView, "%s %s", cursor, choice)
		}
	}

	if m.isGameOn {
		header = ""
		mainView.Reset()
		fmt.Fprintf(
			&mainView,
			"Great! You have selected the %s difficulty level.\nYou have %d chances to guess the correct number.\n\nLet's start the game!\n\n",
			m.controller.GetLevel(), m.controller.MaxAttempts)

		fmt.Fprintf(
			&mainView,
			"%s",
			lipgloss.JoinVertical(lipgloss.Left,
				lipgloss.JoinHorizontal(lipgloss.Left, "Enter your choice: ", m.textinput.View()),
				"\n"+m.result,
			),
		)
	}

	if !m.isChoiceOn && !m.isGameOn {
		return "\n" + m.result + footer
	}

	if m.quit {
		return "\nThanks for playing! See you next time.\n\n"
	}

	return fmt.Sprintf("%s\n%s\n%s\n", header, mainView.String(), footer)
}
