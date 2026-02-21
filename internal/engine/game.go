package engine

import "math/rand"

type IGame interface {
	Check(guess int) GuessResult
	generateRandomNumber() int
}

type GuessResult int

const (
	High GuessResult = iota
	Low
	Correct
)

type GameEngine struct {
	mysteryNum int
}

func NewGame() *GameEngine {
	return &GameEngine{
		mysteryNum: generateMysteryNumber(),
	}
}

func (g *GameEngine) Check(guess int) GuessResult {
	if g.mysteryNum == guess {
		return Correct
	}

	if g.mysteryNum < guess {
		return Low
	}

	return High
}

func (g *GameEngine) GetMysteryNum() int {
	return g.mysteryNum
}

func (r GuessResult) String() string {
	return [...]string{"less than", "greater than", "correct"}[r]
}

func generateMysteryNumber() int {
	max, min := 101, 1
	return rand.Intn(max-min) + min
}
