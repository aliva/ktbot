package main

import "fmt"

func send() {
	client := getClient()
	msg := generateMessage()
	fmt.Printf("Sending: '%s'\n", msg)

	_, _, err := client.Statuses.Update(msg, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}
