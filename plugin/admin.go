package main

import "fmt"

type adminString string

func (u adminString) Welcome() {
	fmt.Printf("Welcome admin %v!\n", u)
}

var User adminString
