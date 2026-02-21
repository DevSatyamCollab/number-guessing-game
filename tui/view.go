package tui

import (
	"fmt"
	"strconv"
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
			lineContent := fmt.Sprintf("%s %s", cursor, choice)
			if m.cursor == i {
				cursor = ">"
				lineContent = fmt.Sprintf(m.style.Cursor.Render("%s %s"), cursor, choice)
			}

			fmt.Fprintf(&mainView, "%s", lineContent)
		}
	}

	if m.isGameOn {
		header = ""
		mainView.Reset()
		fmt.Fprintf(
			&mainView,
			"Great! You have selected the %s difficulty level.\nYou have %s chances to guess the correct number.\n\nLet's start the game!\n\n",
			m.style.Cursor.Render(m.controller.GetLevel()),
			m.style.Cursor.Render(strconv.Itoa(m.controller.MaxAttempts)),
		)

		fmt.Fprintf(
			&mainView,
			"%s",
			lipgloss.JoinVertical(lipgloss.Left,
				lipgloss.JoinHorizontal(lipgloss.Left,
					m.style.Cursor.Render("Enter your choice: "),
					m.textinput.View(),
				),
				"\n"+m.result,
			),
		)
	}

	if !m.isChoiceOn && !m.isGameOn {
		return "\n" + m.result + footer
	}

	if m.quit {
		return m.style.Quit.Render("\nThanks for playing! See you next time.\n\n")
	}

	return fmt.Sprintf("%s\n%s\n%s\n", header, mainView.String(), footer)
}
