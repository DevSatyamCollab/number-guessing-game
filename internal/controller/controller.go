package controller

import (
	"errors"
	"number-guessing-game/internal/engine"
	"strconv"
)

type GameState int

const (
	Playing GameState = iota
	Won
	Lost
)

type Controller struct {
	engine      engine.GameEngine
	state       GameState
	level       int
	attempts    int
	MaxAttempts int
}

func NewController(level int) *Controller {
	return &Controller{
		engine:      *engine.NewGame(),
		state:       Playing,
		attempts:    0,
		level:       level,
		MaxAttempts: getMaxAttempts(level),
	}
}

// check if have the valid attempts or not
func (c *Controller) ValidateGuess(input string) (int, error) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}

	return n, nil
}

func (c *Controller) HandleGuess(input string) (string, error) {
	if c.state != Playing {
		return "", errors.New("game over")
	}

	// validate guess
	guess, err := c.ValidateGuess(input)
	if err != nil {
		return "", errors.New("please enter a valid number")
	}

	// process game logic
	c.attempts++
	res := c.engine.Check(guess)

	// update the state
	if res == engine.Correct {
		c.state = Won
	}

	if c.attempts >= c.MaxAttempts {
		c.state = Lost
	}

	return res.String(), nil
}

func (c *Controller) GetProgess() (int, GameState) {
	return c.attempts, c.state
}

func (c *Controller) GetLevel() string {
	switch c.level {
	case 0:
		return "Easy"
	case 1:
		return "Medium"
	default:
		return "Hard"
	}
}

func (c *Controller) GetSecretNum() int {
	return c.engine.GetMysteryNum()
}

func getMaxAttempts(level int) int {
	switch level {
	case 0:
		return 10
	case 1:
		return 5
	default:
		return 3
	}
}
