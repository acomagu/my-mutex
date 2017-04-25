package main

import (
	"fmt"
)

var k = 0

func kanesawa() {
	k++
}

func ito() {
	k++
}

func main() {
	wait := make(chan bool)

	go func() {
		for i := 0; i < 10000; i++ {
			kanesawa()
		}
		wait <- true
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			ito()
		}
		wait <- true
	}()

	<-wait
	<-wait

	fmt.Println(k)
}
