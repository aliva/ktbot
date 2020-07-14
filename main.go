package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Booting up ktbot...")

	for {
		select {
		case <-time.After(time.Minute * 1):
			if canSend() {
				send()
			} else {
				fmt.Println("Skip,", time.Now())
			}
		}
	}
}
