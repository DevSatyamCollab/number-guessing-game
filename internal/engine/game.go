package engine

import "math/rand"

type IGame interface {
	Check(guess int) GuessResult
	generateRandomNumber() int
}

// hint
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
	if guess > g.mysteryNum {
		return High
	}

	if guess < g.mysteryNum {
		return Low
	}

	return Correct
}

func (g *GameEngine) GetMysteryNum() int {
	return g.mysteryNum
}

func (r GuessResult) String() string {
	return [...]string{"less than", "greater than", "correct"}[r]
}

func generateMysteryNumber() int {
	return rand.Intn(100) + 1
}
