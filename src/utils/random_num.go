package utils

import (
	"fmt"
	"math/rand"
)

func GetRandomNum(min int, max int) int {
	if min >= max {
		fmt.Printf("min: %d, max: %d\n", min, max)
		panic("min should be less than max")
	}
	return rand.Intn(max-min) + min
}

func GenerateRandomTorusPairs(pairCount int) [][2]int {
	const maxProduct = 1000000
	pairs := make([][2]int, 0, pairCount)
	change_direction := false
	for len(pairs) < pairCount {
		A := rand.Intn(1000) + 1 
		B := rand.Intn(maxProduct/A) + 1 
		if change_direction {
			A, B = B, A
		}
		pairs = append(pairs, [2]int{A, B})
		change_direction = !change_direction
	}

	return pairs
}
