package main

import (
	"fmt"
)

func _lock(flag *int32)

var locked = int32(0)

func lock() {
	_lock(&locked)
}

func unlock() {
	locked = 0
}

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
			lock()
			kanesawa()
			unlock()
		}
		wait <- true
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			lock()
			ito()
			unlock()
		}
		wait <- true
	}()

	<-wait
	<-wait

	fmt.Println(k)
}
