package engine

import "math/rand"

type IGame interface {
	Check(guess int) GuessResult
	GenerateRandomNumber() int
}

type GuessResult int

const (
	High GuessResult = iota
	Low
	Correct
)

type GameEngine struct {
	MaxAttempts int
	MysteryNum  int
}

func NewGame(attmepts int) *GameEngine {
	return &GameEngine{
		MysteryNum:  generateMysteryNumber(),
		MaxAttempts: attmepts,
	}
}

func (g *GameEngine) Check(guess int) GuessResult {
	if g.MysteryNum == guess {
		return Correct
	}

	if g.MysteryNum < guess {
		return Low
	}

	return High
}

func generateMysteryNumber() int {
	max, min := 101, 1
	return rand.Intn(max-min) + min
}
