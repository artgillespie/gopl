package main

import "fmt"

type EntityNotFoundErr struct{}

func (e EntityNotFoundErr) Error() string { return "EntityNotFoundErr " }

func getError() error {
	err := EntityNotFoundErr{}
	fmt.Printf("Address in getError: %p\n", &err)
	return err
}

func getError2() error {
	err := &EntityNotFoundErr{}
	fmt.Printf("Address in getError2: %p\n", err)
	return err
}

func main() {
	err := getError()
	fmt.Printf("Address in main: %p\n", &err)
	err = getError2()
	fmt.Printf("Address in main: %p\n", err)
}
