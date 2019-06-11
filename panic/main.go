package main

import (
	"log"
	"time"
)

func asyncJob() {
	i := 1
	ticker := time.NewTicker(500 * time.Millisecond)
	for t := range ticker.C {
		log.Printf("Doing this %d - %v \n", i, t)
		i++
	}
}

func divide(a int, b int) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovering inside function", r)
		}
	}()

	c := a / b
	log.Println(c)
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovering", r)
		}
	}()

	//	go asyncJob()

	// all good here
	divide(4, 3)

	time.Sleep(1 * time.Second)

	// this will panic
	divide(4, 0)

	// sleep again, so we see if we returned to normal functioning
	time.Sleep(2 * time.Second)

	// this should work
	divide(8, 1)

}
