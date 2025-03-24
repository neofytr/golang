package main

import "fmt"

type person_t struct {
	name   string
	height float64
}

func returnInfoString(info person_t) string {
	return fmt.Sprintf("name: %s, height: %.1f", info.name, info.height)
}

func returnCarInfoToString(car struct {
	Make   string
	Model  string
	Height int
}) string {
	return fmt.Sprintf("make: %s, model: %s, height: %dm\n", car.Make, car.Model, car.Height)
}

func main() {
	raj := person_t{}
	raj.name = "raj"
	raj.height = 4.5
	var msg string = returnInfoString(raj)
	println(msg)

	// anonymous structs; avoid them
	myCar := struct {
		Make   string
		Model  string
		Height int
	}{Make: "tesla", Model: "model B", Height: 3}

	var carMsg string = returnCarInfoToString(myCar)
	println(carMsg)
}
