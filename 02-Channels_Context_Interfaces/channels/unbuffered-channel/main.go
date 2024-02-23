package main

import "fmt"

func main() {
	dataChan := make(chan int)

	go func() {
		dataChan <- 789
	}()

	n := <-dataChan

	fmt.Printf("n = %d\n", n)
}
