package engine

import (
	"blind_snake/src/game"
)
type Engine struct {
	game *game.Game
}

func NewEngine(width int, height int, one_apple bool) *Engine {
	return &Engine{game.NewGame(width, height, one_apple)}
}
func (e *Engine) SetNewGame(width int, height int, one_apple bool) {
	e.game = game.NewGame(width, height, one_apple)
}
func (e *Engine) Play() (bool, int){ //returns if he won the game and steps count
	move_count := 0
	switch_on := 2
	switch_row_count := 0
	main_moves := []rune{'r', 'd'}
	main_move_index := 0
	for{		
		move_count++

		move := main_moves[main_move_index]
		main_move_index = (main_move_index + 1) % 2
		
		if switch_row_count == switch_on{
			move = 'r'
			switch_on+=2
			switch_row_count = -1

		}
		switch_row_count++
		status, err := e.game.Move(move)
		if err != nil {
			return false, move_count
		}
		if status {
			return true, move_count
		}	
		
	}
}