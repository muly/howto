package main

// based on the below blogpost
// https://varunksaini.com/blog/backoff-retry-in-golang/
// or the below cached copy
// https://webcache.googleusercontent.com/search?q=cache:NJUoeIN_KV8J:https://varunksaini.com/blog/backoff-retry-in-golang/+&cd=3&hl=en&ct=clnk&gl=us

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff"
	"log"
	"time"
)

func main() {
	a, b := 1, 2
	operation := func() error {
		err := doSomething(a, b)
		return err
	}

	//constantBackOff(operation)
	exponentialBackOff(operation)
}

func doSomething(a, b int) error {
	//process a, b; and return error
	fmt.Println("running doSomething()")

	//return nil
	return errors.New("simulated error")
}

func exponentialBackOff(operation backoff.Operation) {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 30 * time.Second

	err := backoff.Retry(operation, b)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
}

func constantBackOff(operation backoff.Operation) {
	b := backoff.NewConstantBackOff(1 * time.Second)
	//Note: runs forever as it is not possible to set the max elapsed time

	err := backoff.Retry(operation, b)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}
}
