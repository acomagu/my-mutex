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
		fmt.Println("f0")
		f0 = true
		for f1 == true {
			if turn != 0 {
				f0 = false
				for turn != 0 {
				}
				f0 = true
			}
		}

		fmt.Println("a")
		time.Sleep(1 * time.Second)

		turn = 1
		f0 = false
	}()

	go func() {
		fmt.Println("f1")
		f1 = true
		for f0 == true {
			if turn != 1 {
				f1 = false
				for turn != 1 {
				}
				f1 = true
			}
		}

		fmt.Println("b")
		time.Sleep(1 * time.Second)

		turn = 0
		f1 = false
	}()

	time.Sleep(3 * time.Second)

}
