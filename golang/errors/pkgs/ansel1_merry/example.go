package main

import (
	"errors"
	"fmt"

	"github.com/ansel1/merry"
)

func main() {
	var err error

	// uncomment and try different examples
	//err = noError()
	//err = regularError()
	//err = merryUserError("this is user generated error")
	//err = merryWrapError()
	//err = merryAppendToRegularError("here is the error")
	err = merryAppendToRegularError("here is the error")

	if err != nil {
		fmt.Println(err)
	}
}

func noError() error {
	return nil
}

func regularError() error {
	return errors.New("some regular error")
}

func merryUserError(msg string) merry.Error {
	return merry.New(msg)
}

func merryWrapError() merry.Error {
	var err error

	err = noError()
	if err != nil {
		return merry.Wrap(err) // wraping nil error
	}

	err = regularError()
	if err != nil {
		return merry.Wrap(err) // wraping a regular error
	}

	err = merryUserError("this is user generated error")
	if err != nil {
		return merry.Wrap(err) // wraping a merry error
	}

	return nil
}

func merryAppendToMerryError(msg string) merry.Error {
	var err error

	err = merryUserError("user defined error")
	if err != nil {
		return merry.WithMessagef(err, "%v: %v", msg, err.Error())
	}

	return nil
}

func merryAppendToRegularError(msg string) merry.Error {
	var err error

	err = regularError()
	if err != nil {
		return merry.WithMessagef(err, "%v: %v", msg, err.Error())
	}

	return nil
}
