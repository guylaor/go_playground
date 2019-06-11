package main

import (
	"log"
)

func divide(a int, b int) {

	// this is the "try catch"
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovering inside function", r)
		}
	}()

	c := a / b
	log.Println(c)
}

func main() {

	// this defer is useless, becasue we are in main
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovering", r)
		}
	}()

	// all good here
	divide(4, 2)

	// this will panic
	divide(4, 0)

	// this should work
	divide(8, 2)

}
