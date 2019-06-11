package main

import (
	"log"
	"time"
)

func asyncJob(done chan bool) {
	i := 1
	ticker := time.NewTicker(500 * time.Millisecond)
	for t := range ticker.C {
		go divide(i, 3)
		if i == 4 {
			go divide(i, 0)
		}
		log.Printf("Doing this %d - %v \n", i, t)
		i++
		if i == 10 {
			done <- true
		}
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

	// making some async job
	done := make(chan bool, 1)
	go asyncJob(done)
	<-done

	log.Println("finished program")

}
