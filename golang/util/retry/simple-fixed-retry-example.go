package main

import (
	"fmt"
)

const backoffAttempts = 5

type retryable func() error

func main() {

	a, b := 0, 0

	res := 0

	f := func() error {
		var err error
		res, err = myFunction(a, b)
		if err != nil {
			return err
		}
		return nil
	}

	err := withSimpleRetry(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("output is ", res)
}

func withSimpleRetry(r retryable) error {
	i := 0
	var err error
	for i = 1; i <= backoffAttempts; i++ {
		if err = r(); err != nil {
			fmt.Println("received error. retrying... attempt", i)
			// TODO: sleep for rand interval between 0 and 5 sec
			continue
		}
		return nil
	}
	return fmt.Errorf("failed after %v errors. last error is %v", i, err)
}

func myFunction(a, b int) (int, error) {
	//return 5, nil // enable to test the positive case
	return 0, fmt.Errorf("simulated error") // enable this to test the negative case
}
