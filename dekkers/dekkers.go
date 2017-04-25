package main

import (
	"fmt"
)

var flags = make([]bool, 2)
var turn = 0

func lock(id int) {
	revid := id ^ 1
	flags[id] = true
	for flags[revid] == true {
		if turn != id {
			flags[id] = false
			for turn != id {
			}
			flags[id] = true
		}
	}
}

func unlock(id int) {
	revid := id ^ 1
	turn = revid
	flags[id] = false
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
			lock(0)
			kanesawa()
			unlock(0)
		}
		wait <- true
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			lock(1)
			ito()
			unlock(1)
		}
		wait <- true
	}()

	<-wait
	<-wait

	fmt.Println(k)
}
