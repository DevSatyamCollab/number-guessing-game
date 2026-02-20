package tui

import "fmt"

func (m model) View() string {
	header := `
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Please select the difficulty level:
`
	mainView := "\n"
	footer := "\nPress q for Quit\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		mainView += fmt.Sprintf("%s %s", cursor, choice)
	}

	if m.isGameOn {
		header = ""
		levelstr := ""
		attempts := 0
		switch m.cursor {
		case 0:
			levelstr = "Easy"
			attempts = 10
		case 1:
			levelstr = "Medium"
			attempts = 5
		case 2:
			levelstr = "Hard"
			attempts = 3
		}
		mainView = fmt.Sprintf("Great! You have selected the %s difficulty level.\nYou have %d chances to guess the correct number.\n\nLet's start the game!\n", levelstr, attempts)
	}

	return fmt.Sprintf("%s\n%s\n%s\n", header, mainView, footer)
}
