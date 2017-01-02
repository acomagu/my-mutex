package main

import (
	"fmt"
	"time"
)

func main() {
	f0 := false
	f1 := false
	turn := 0

	go func() {
		f0 = true
		for f1 == true {
			if turn != 0 {
				f0 = false
				for turn != 0 {
				}
				f0 = true
			}
		}

		fmt.Println("I have a Pen.")
		fmt.Println("I have an Apple.")
		fmt.Println("Ah!")
		fmt.Println("Apple-Pen!")
		fmt.Println()

		turn = 1
		f0 = false
	}()

	go func() {
		f1 = true
		for f0 == true {
			if turn != 1 {
				f1 = false
				for turn != 1 {
				}
				f1 = true
			}
		}

		fmt.Println("I have a Pen")
		fmt.Println("I have a Pineapple")
		fmt.Println("Ah!")
		fmt.Println("Pineapple-Pen!")
		fmt.Println()

		turn = 0
		f1 = false
	}()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Apple-Pen,")
	fmt.Println("Pineapple-Pen,")
	fmt.Println("Ah!")
	fmt.Println("Pen-Pineapple-Apple-Pen!")

}
