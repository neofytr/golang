package main

import "fmt"

type person_t struct {
	name   string
	height float64
}

func main() {
	p := person_t{name: "Alice", height: 5.7}
	fmt.Printf("name: %s, height: %.1f\n", p.name, p.height)
}
