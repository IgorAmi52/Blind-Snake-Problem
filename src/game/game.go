package game

import (
	"blind_snake/src/utils"
	"errors"
)

type Game struct{
	torus [][]int
	torus_width int
	torus_height int
	snake_x int
	snake_y int
	apples_left int
	moves_left int
}

func NewGame(width int, height int, one_apple bool) *Game {
	torus := make([][]int, width)
	apple_x, apple_y := -1, -1
	apples_left := height * width - 1
	moves_left := height * width * 35
	if one_apple {
		apple_x = utils.GetRandomNum(0, width)
		apple_y = utils.GetRandomNum(0, height)
		apples_left = 1
	}
	for i := 0; i < width; i++ {
		torus[i] = make([]int, height)
		for j := 0; j < height; j++ {
			if one_apple && i == apple_x && j == apple_y {
				torus[i][j] = 1
				continue
			}
			if !one_apple {
				torus[i][j] = 1
				continue
			}
			torus[i][j] = 0
		}
	}
	snake_x := utils.GetRandomNum(0, width)
	snake_y := utils.GetRandomNum(0, height)
	// ensuring that the snake does not spawn on the apple
	for snake_x == apple_x && snake_y == apple_y {
		snake_x = utils.GetRandomNum(0, width)
		snake_y = utils.GetRandomNum(0, height)
	}
	torus[snake_x][snake_y] = 0
	return &Game{torus: torus, torus_width: width, torus_height: height, snake_x: snake_x, snake_y: snake_y, apples_left: apples_left, moves_left: moves_left}
}

func (g *Game) Move(dir rune) (bool, error) {
	// move the snake in the given direction
	// if the snake eats all the apples, return true
	// if the snake runs out of moves, return false
	// if the snake runs out of moves, return an error
	switch dir {
		case 'u', 'U':
			g.snake_y = (g.snake_y - 1 + g.torus_height) % g.torus_height
		case 'd', 'D':
			g.snake_y = (g.snake_y + 1) % g.torus_height
		case 'l', 'L':
			g.snake_x = (g.snake_x - 1 + g.torus_width) % g.torus_width
		case 'r', 'R':
			g.snake_x = (g.snake_x + 1) % g.torus_width
	}
	if g.torus[g.snake_x][g.snake_y] == 1 {
		g.torus[g.snake_x][g.snake_y] = 0
		g.apples_left--
		if g.apples_left == 0 {
			return true, nil
		}
	}
	g.moves_left--
	if g.moves_left == 0 {
		return false, errors.New("Game Over: No moves left")
	}
	return false, nil
}