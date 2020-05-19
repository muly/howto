// goroutine sending a data as well as error as response back to calling routine using channels
// with 2 goroutines

package main

import "fmt"

func main() {

	resCh1 := make(chan float64)
	resCh2 := make(chan float64)
	errCh1 := make(chan error)
	errCh2 := make(chan error)

	go divide(2, 1, resCh1, errCh1)
	go divide(10, 0, resCh2, errCh2)

	r1 := <-resCh1
	r2 := <-resCh2

	if err := <-errCh1; err != nil {
		fmt.Println(err)
		return
	}
	if err := <-errCh2; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r1, r2)
}

func divide(a, b float64, resCh chan float64, errCh chan error) {
	if b == 0 {
		resCh <- 0
		errCh <- fmt.Errorf("divide by zero error")
		return
	}
	resCh <- a / b
	errCh <- nil
	return
}
