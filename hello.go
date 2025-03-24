package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := "Hi %s, your open rate is %.1f percent\n"
	fmt.Printf(msg, name, openRate)
	fmt.Println("Hi", name+",", "your open rate is", openRate, "percent")
}
