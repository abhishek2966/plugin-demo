package main

import "fmt"

type userString string

func (u userString) Welcome() {
	fmt.Printf("Welcome user %v!\n", u)
}

var User userString
