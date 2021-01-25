package main

import (
	"errors"
	"fmt"
	"time"

	"k8s.io/client-go/util/retry"
)

var errDBConnectionFailure = errors.New("DB connection failed")

func main() {
	a, b := 2, 1

	retry.OnError(retry.DefaultRetry,
		func(err error) bool {
			// check the error type and decide to retry or not
			if err == errDBConnectionFailure {
				return true
			}
			return false
		},
		func() error {
			if err := myFunc(a, b); err != nil {
				fmt.Printf("myFunc() encountered error: %v\n", err)
				return err
			}
			return nil
		})
}

func myFunc(a, b int) error {
	if a == 0 || b == 0 {
		return fmt.Errorf("invalid data: %v, %v", a, b)
	}
	fmt.Printf("executing myFunc(%v, %v)\n", a, b)
	time.Sleep(time.Second)
	// simulating db connection failure
	return errDBConnectionFailure
}
