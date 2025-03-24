package main

func getInfo() (age int, name string) {
	age = 18
	name = "raj shukla"

	return age, name
}

func main() {
	var age int
	var name string

	age, name = getInfo()
	println(age, name)
}
