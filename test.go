package main

import "fmt"

type person_t struct {
	name   string
	height float64
}

func returnInfoString(info person_t) string {
	return fmt.Sprintf("name: %s, height: %.1f", info.name, info.height)
}

func main() {
	raj := person_t{}
	raj.name = "raj"
	raj.height = 4.5

	var msg string = returnInfoString(raj)
	println(msg)
}
