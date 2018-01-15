package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

	bkoff := BackoffConfig{}
	setDefaults(&bkoff)
	bkoff.MaxRetry = 5         // overwrite the default values
	bkoff.RetryForever = false // overwrite the default values for testing

	err := bkoff.retry(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("output is ", res)
}

func myFunction(a, b int) (int, error) {
	fmt.Println("executing myFunction")
	//return 5, nil // enable to test the positive case
	return 0, fmt.Errorf("simulated error") // enable this to test the negative case
}

func (bc BackoffConfig) retry(r retryable) error {
	i := 0
	var err error

	for i = 1; ; i++ {
		if err = r(); err != nil {
			fmt.Println("received error. retrying... attempt", i)
			delay := bc.backoff(i)
			if delay < 0 {
				return errors.New("max retry duration exceeded, still failing. latest error: " + err.Error())
			}
			time.Sleep(delay)
			continue
		}
		return nil
	}

}

///////////////////////////////////
// below code is copied from https://github.com/grpc/grpc-go/blob/master/backoff.go

// DefaultBackoffConfig uses values specified for backoff
var DefaultBackoffConfig = BackoffConfig{
	MaxRetry:     10,
	MaxDelay:     120 * time.Second,
	RetryForever: false,
	baseDelay:    1 * time.Second,
	factor:       1.6,
	jitter:       0.2,
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
	// MaxRetry is the upper bound of number of retries
	MaxRetry int

	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration

	// RetryForever indicate to retry forever or stop after max retry count is reached.
	RetryForever bool

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
	if retries >= bc.MaxRetry && !bc.RetryForever {
		return -1 // return negative time to indicate to stop retrying
	}
	backoff, max := float64(bc.baseDelay), float64(bc.MaxDelay)
	for backoff < max && retries > 0 {
		backoff *= bc.factor
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + bc.jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
