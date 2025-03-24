package main

import "fmt"

type person_t struct {
	name   string
	height float64
}

func test(info person_t) string {
	return fmt.Sprintf("name: %s, height: %.1f", info.name, info.height)
}

func main() {
	var msg string = test(person_t{name: "raj", height: 4.5})
	println(msg)
}
