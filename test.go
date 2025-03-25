package main

import (
	"errors"
	"fmt"
)

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
	return fmt.Sprintf("make: %s, model: %s, height: %d m", car.Make, car.Model, car.Height)
}

// Embedded Structs
type dimension_t struct {
	height int
	widht  int
	length int
}

type car_t struct {
	make        string
	model       string
	dimension_t // this embeds the dimension_t structure fields in the car_t struct
	// so now, to access the height through a variable car of type car_t, we would
	// simply do car.height
}

// Composite structs
type human_t struct {
	name   string
	weight int
	dimen  dimension_t // this is a composite struct
	// to access height through a variable human of type human_t, we would
	// do human.dimen.height
}

// method on the type human_t
func (human human_t) returnVolume() int {
	// human contains the human_t structure's copy on which
	// the method returnVolume() was called
	return human.dimen.height * human.dimen.length * human.dimen.widht
}

// method on the type car_t
func (car car_t) returnVolume() int {
	return car.height * car.widht * car.length
}

// both car_t and human_t implement this interface(implicitly)
// since they both implement the functions required by the volume interface
type volume_t interface {
	returnVolume() int
}

// the empty interface, interface{}, is always implemented by every type
// because it has no requirements

// In Go, interface{} is an empty interface, meaning it can hold values of any type.
// Since all types implement at least the empty interface, it acts as a universal container.

// a single type can implement multiple interfaces since implementing an
// interface is an implicit thing

// we can pass the structs of any type that implements
// the volume interface in this function
func getVolume(volume volume_t) int {
	return volume.returnVolume()
}

func getSomeInfo(volume volume_t) int {

	// type assertions
	// since we don't know what types are being passed in the interface
	// we use type assertions

	/*

		Type assertions in Go allow you to retrieve the concrete value from an interface{} type.
		This is useful when you have a value stored as an interface{} and need to
		convert it back to its original type.

		We can also use some custom interface type other than interface{} like volume_t

	*/

	var retVal int
	car, ok := volume.(car_t)
	if ok {
		retVal = car.height
	}

	human, ok := volume.(human_t)
	if ok {
		retVal = human.dimen.height
	}

	return retVal
}

// implements the same functionality as above but uses type switches
func getSome(volume volume_t) int {
	var retVal int

	switch v := volume.(type) {
	case car_t:
		retVal = v.height
	case human_t:
		retVal = v.dimen.height
	}

	return retVal
}

// an interface with methods that have names for their values and return values (for clarity only)
/*

When we implement the copy function for a type
that satisfies the copier_t interface, we
can use different names for the parameters and return values as
long as the function signature matches (signature includes the function name too)

*/
type copier_t interface {
	copy(destinationFile, sourceFile string) (bytesCopied int)
}

/*

In Go, error is a built-in interface that represents an error condition.
The error interface is defined as:

type error interface {
	Error() string
}

This means that any type that implements the Error() method (which returns a string)
satisfies an error interface and can be used as an error

*/

/*

In Go, when you pass an error interface to the fmt.Print functions, it automatically
calls the Error() method of the error interface behind the scenes

*/

type user_t struct {
	username   string
	password   string
	userActive bool
	activeTime int
}

type error_t struct {
	code int
	msg  string
}

// make error_t implement the error interface by implementing the Error() function on it
func (err error_t) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", err.code, err.msg)
}

func getActiveTime(user user_t) (int, error) {
	if user.userActive {
		return user.activeTime, nil
	} else {
		return 0, error_t{code: 404, msg: fmt.Sprintf("user %s not active", user.username)}
	}
}

// test performs integer division and returns the quotient, remainder, and an error if the dividend is zero.
func test(divisor, dividend int) (quotient, remainder int, err error) {
	if dividend == 0 {
		// errors.New("Error: Cannot divide by zero!") internally does the following:
		// - It creates a new instance of an unexported struct called `errorString` (defined in the `errors` package).
		// - The struct `errorString` has a single field: `s string`, which stores the error message.
		// - The struct implements the `Error() string` method, making it compatible with the `error` interface.
		// - Finally, errors.New() returns a pointer to this struct, allowing it to be used as an error.
		return 0, 0, errors.New("Error: Cannot divide by zero!")
	}

	return divisor / dividend, divisor % dividend, nil
}

func fizzbuzz() {
	for index := 1; index <= 100; index++ {
		if index%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if index%3 == 0 {
			fmt.Println("fizz")
		} else if index%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(index)
		}
	}
}

func getMessages() [3]string {
	return [3]string{"hello", "world", "kaise ho"}
}

// the array passed to this function is passed by value and not by reference
func val(arr [3]int) {
	arr[0] = 10
	arr[1] = 11
	fmt.Println(arr[0], arr[1])
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

	// composite structs
	human := human_t{name: "raj", weight: 80, dimen: dimension_t{length: 20, height: 67, widht: 4}}

	// embedded struct
	// we still do a similar thing as composite structs but we initialize with the
	// type name instead of the name of the variable of the type as in composite structs
	car := car_t{make: "tesla", model: "model b", dimension_t: dimension_t{length: 20, height: 40, widht: 4}}

	println(human.name)

	// the real difference comes will accessing the fields
	println(human.dimen.height)
	println(car.height)

	println(car.returnVolume())
	println(human.returnVolume())

	println(getVolume(car))
	println(getVolume(human))

	println(getSome(car))
	println(getSome(human))

	println(getSomeInfo(car))
	println(getSomeInfo(human))

	user := user_t{username: "raj", password: "rishika", userActive: false, activeTime: 0}

	activeTime, err := getActiveTime(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("user %s active time: %d\n", user.username, activeTime)
	}

	quotient, remainder, err := test(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("quotient: %d, remainder: %d\n", quotient, remainder)
	}

	// a for loop in GO
	for index := 0; index < 10; index++ {
		fmt.Printf("cost for index %d is %0.3f\n", index, 0.01*float64(index))
	}

	// there is no explicit while loop
	// a while loop is just a for loop with just the condition
	index := 0
	for index < 10 {
		fmt.Printf("cost for index %d is %0.3f\n", index, 0.01*float64(index))
		index++
	}

	fizzbuzz()

	// the continue keyword stops the current iteration of a loop and continues
	// to the next iteration (if there is any)

	// the break keyword stops the current iteration of the loop and immediately
	// exits the loop

	// whenever and wherever in go we don't initialize a variable, it takes on it's
	// default zero value
	// for example, for int, the zero value is 0, for string, it is ""

	// this is an array of three integers
	var arr [3]int
	arr = [3]int{1, 2, 3}

	// the array is passed by value
	val(arr)
	fmt.Println(arr[0], arr[1])

	/*

		An array in Go is a fixed-size collection of elements of the same type
		Arrays are value types: assigning an array to another copies all elements
		The size of an array is part of it's type

	*/

	var array [5]int              // declaring a variable that will hold an array of type [5]int
	array = [5]int{1, 2, 3, 4, 5} // assigning the array variable an array of type [5]int

	var arrayCopy [5]int
	arrayCopy = array // this will create a copy of the array stored inside the variable array and assign the copy
	// to the variable arrayCopy

	_ = arrayCopy // to avoid compiler warning

	/*

		A slice is a dynamically-sized, flexible view into an array
		A slice consists of:
			1. A pointer to an array
			2. A length (number of elements)
			3. A capacity (number of elements in the underlying array starting from the pointer)

		Slices reference an underlying array, so modifying a slice modifies the original array
		Slices have length (num of accessible elements) and capacity (max elements available)
		Slices can grow dynamically using append()
	*/

	slice := array[1:4] // slice from index 1 to index 3 of the array variable array (both inclusive)

	fmt.Println(slice)      // [2, 3, 4]
	fmt.Println(len(slice)) // 3 (length of slice)
	fmt.Println(cap(slice)) // 4 (capacity from array[1] (start of the slice) to array[4] (end of the underlying array))

	slice[0] = 20 // modifies array[1] too
	fmt.Println(array)

	/*

			append(slice, elements...) adds elements to a slice and returns the updated slice
			If capacity allows, the elements are added to the same underlying array
			If capacity is exceeded, Go allocates a new array, copies elements, and uses that. So, its
			important to store the returned slice somewhere, often in the previous variable itself

			Since `append()` may allocate a **new array**, always assign its result back: `slice = append(slice, newElement)`
		    If we don't, the new slice will be lost, and modifications may not persist.

	*/

	// Append an element to the slice
	slice = append(slice, 12)
	// Since the slice's current length (4) is within the underlying array's capacity (5),
	// the value `12` is placed in array[4], modifying the original array.
	// The slice now looks like: [1, 2, 3, 4, 12]
	// and the underlying `array` is updated: [1, 2, 3, 4, 12]

	fmt.Println("Array after first append:", array) // [1 2 3 4 12]
	fmt.Println("Slice after first append:", slice) // [1 2 3 4 12]

	// Append another element to the slice
	slice = append(slice, 13)
	// This time, the slice's length would increase beyond the array's capacity (5),
	// meaning Go can no longer store this slice in `array`.
	// Go automatically creates a **new underlying array**, copies the elements from the
	// previous slice, and adds `13` at the end.
	// This new array is completely separate from `array`.

	fmt.Println("Slice after second append:", slice) // [1 2 3 4 12 13]

	// The original `array` remains unchanged beyond its capacity:
	fmt.Println("Array after second append:", array) // [1 2 3 4 12] (unchanged)

	// The new slice now has its own memory, independent of `array`.
}
