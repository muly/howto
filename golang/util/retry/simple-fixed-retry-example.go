package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

	bkoff := BackoffConfig{}
	setDefaults(&bkoff)

	for i = 1; i <= backoffAttempts; i++ {
		if err = r(); err != nil {
			fmt.Println("received error. retrying... attempt", i)
			delay := bkoff.backoff(i)
			if delay < 0 {
				return errors.New("max retry duration exceeded, still failing. latest error: " + err.Error())
			}
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

///////////////////////////////////
// below code is copied from https://github.com/grpc/grpc-go/blob/master/backoff.go

// DefaultBackoffConfig uses values specified for backoff
var DefaultBackoffConfig = BackoffConfig{
	MaxDelay:  120 * time.Second,
	baseDelay: 1.0 * time.Second,
	factor:    1.6,
	jitter:    0.2,
}

// backoffStrategy defines the methodology for backing off after a grpc
// connection failure.
type backoffStrategy interface {
	// backoff returns the amount of time to wait before the next retry given
	// the number of consecutive failures.
	backoff(retries int) time.Duration
}

// BackoffConfig defines the parameters for the default gRPC backoff strategy.
type BackoffConfig struct {
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration

	// baseDelay is the amount of time to wait before retrying after the first
	// failure.
	baseDelay time.Duration

	// factor is applied to the backoff after each retry.
	factor float64

	// jitter provides a range to randomize backoff delays.
	jitter float64
}

func setDefaults(bc *BackoffConfig) {
	md := bc.MaxDelay
	*bc = DefaultBackoffConfig

	if md > 0 {
		bc.MaxDelay = md
	}
}

func (bc BackoffConfig) backoff(retries int) time.Duration {
	if retries == 0 {
		return bc.baseDelay
	}
	backoff, max := float64(bc.baseDelay), float64(bc.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= bc.factor
		retries--
	}
	if backoff > max {
		return -1
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + bc.jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
