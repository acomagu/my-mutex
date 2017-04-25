package main

import (
	"fmt"
)

func increment(k *int32)

var k = int32(0)

func kanesawa() {
	increment(&k)
}

func ito() {
	increment(&k)
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
