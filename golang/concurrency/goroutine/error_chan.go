// goroutine sending a data as well as error as response back to calling routine using channels

package main

import "fmt"

func main() {

	resCh := make(chan float64)
	errCh := make(chan error)

	go divide(2, 1, resCh, errCh)

	r := <-resCh
	err := <-errCh

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)
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
