package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// sample prgram for creating custom error (error handling in Golang)
func main() {
	defer catchError() // to put an catchError() function at the end of execution

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}
}

// function to check input name is empty or not
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	} else if _, err := strconv.Atoi(input); err == nil { // a condition to check whether the input is a number or character
		return false, errors.New("input a character not a number")
	}
	return true, nil
}

//function to catch an Error in input
func catchError() {
	if r := recover(); r != nil {
		fmt.Println("Error ecountered", r)
	} else {
		fmt.Println("App running perfectly")
	}
}
