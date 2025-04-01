package mathutil

import "fmt"

func init() {
	fmt.Println("initialized mathUtil package")
}

// exported function (starts with a capital letter)
func GetSquare(num int) (square int) {
	return num * num
}

func GetDouble(num int) (double int) {
	return 2 * num
}

// this function is not exported (starts with a small letter)
func getSquare(num int) (square int) {
	return num * num
}
