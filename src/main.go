package main

import (
	"blind_snake/src/engine"
	"blind_snake/src/utils"
	"fmt"
	"sync"
)

func main(){
	test_count := 1000
	randomTorusSizes := utils.GenerateRandomTorusPairs(test_count)
	ch := make(chan [2]int, len(randomTorusSizes))
	for _, pair := range randomTorusSizes {
		ch <- pair
	}
	close(ch)
	var wg sync.WaitGroup
	engineCount := 15
	
	faild_count := 0
	passed_count := 0
	worst_move_usage := 0.0

	for i := 0; i < engineCount; i++ {
		wg.Add(1)
		go func(engineID int) {
			defer wg.Done()
			
			engine := engine.NewEngine(1,10,false) // just a holder
			for pair := range ch {
				engine.SetNewGame(pair[0], pair[1], false)
				status, move_count := engine.Play()
				if status {
					S := pair[0] * pair[1]
					passed_count++
					fmt.Printf("Pass for %d x %d, with %.2f x S moves \n", pair[0], pair[1], float64(move_count)/float64(S))
					if float64(move_count)/float64(S) > float64(worst_move_usage){
						worst_move_usage = float64(move_count)/float64(S)
					}
				}else{
					faild_count++
					fmt.Printf("Failed for %d x %d\n", pair[0], pair[1])
				}
			}
		}(i)
	}
	wg.Wait()

	fmt.Printf("\nPassed: %d, Failed: %d\n", passed_count, faild_count)
	fmt.Printf("Sucess rate: %.2f\n", float64(passed_count)/float64(passed_count+faild_count))
	fmt.Printf("Worst move usage: %.2f x S\n", worst_move_usage)	
}