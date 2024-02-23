package main

import "fmt"

func main() {
	dataChan := make(chan int, 2)

	dataChan <- 789
	dataChan <- 123

	n := <-dataChan
	fmt.Printf("n = %d\n", n)
	n = <-dataChan
	fmt.Printf("n = %d\n", n)
}
