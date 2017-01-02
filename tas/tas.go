package main

import (
	"fmt"
	"time"
)

func testAndSet(lock *int32, v int32) (initial int32)

var locked = int32(0)

func lock() {
	for testAndSet(&locked, 1) == 1 {
	}
}

func unlock() {
	locked = 0
}

func main() {
	go func() {
		lock()

		fmt.Println("I have a Pen.")
		fmt.Println("I have an Apple.")
		fmt.Println("Ah!")
		fmt.Println("Apple-Pen!")
		fmt.Println()

		unlock()
	}()

	go func() {
		lock()

		fmt.Println("I have a Pen")
		fmt.Println("I have a Pineapple")
		fmt.Println("Ah!")
		fmt.Println("Pineapple-Pen!")
		fmt.Println()

		unlock()
	}()

	time.Sleep(100 * time.Millisecond)
}
