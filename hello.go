package main

import "fmt"

func main() {
	messageLen := 10

	if maxMessageLen := 20; messageLen > maxMessageLen {
		fmt.Println("Message not sent")
	} else {
		fmt.Println("Message sent")
	}
}
