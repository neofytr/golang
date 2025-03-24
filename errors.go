package main

import "fmt"

/*

An interface in Go is just a collection of method signatures. If a type implements
those methods, it automatically satisfies the interface

The error interface is:

type error interface {
	Error() string
}

Any type that has a method called Error() (that returns a string) automatically
becomes an error type

*/

// creating a custom type (a structure)
type customError struct {
	code    int
	message string
}

// creating a function on the custom type
func (err customError) Error() string {
	return fmt.Sprintf("Error %d: %s", err.code, err.message)
}

// function that returns a custom error
func doSomething(flag bool) error {
	if flag {
		return customError{code: 404, message: "Resource not found"}
	}

	return nil
}

func main() {
	err := doSomething(true)
	if err != nil {
		fmt.Println(err)
	}
}
