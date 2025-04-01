package main

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"sync"
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

func createMatrix(rows, cols int) (matrix [][]int) {
	matrix = make([][]int, rows)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			matrix[row] = append(matrix[row], row*col)
		}
	}
	return matrix
}

/*

A variadic function in Go is a function that accepts a variable number of arguments
of a specific type
It is defined using ... before the type in the function parameter list

*/

func sums(nums ...int) int {
	// nums can now just be accessed as []int which contain all the arguments passed to the function
	total := 0
	for index := 0; index < len(nums); index++ {
		total += nums[index]
	}

	return total
}

/*

Go does not have a direct equivalent to the "spread operator" like JS or Python (*args)
However, the ... notation can be used in two ways:

1. Defining variadic functions; as shown before
2. Expanding a slice into arguments: When calling a variadic function, a slice can be
expanded into individual arguments (in the order in which they are present in the slice)
using ...

For example:

args := []int{1, 2, 3, 4, 5}
sum := sums(args...)

The call is equivalent to sums(1, 2, 3, 4, 5)

*/

func changeMap(personAgeMap map[string]int) {
	personAgeMap["yug"] = 18
	personAgeMap["aditya"] = 19
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

		A slice type is denoted [](type)
		For eg, [](int), [](float)

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

	/*
		CREATING A SLICE USING make()

		- The `make()` function is used to create a slice with a specific length and capacity.
		- The syntax is: `make([]Type, length, capacity)`
		- The `length` defines how many elements are initialized.
		- The `capacity` defines the total size of the underlying array backing the slice.
		- If the `capacity` argument is omitted, it defaults to the `length`, meaning the slice has
		  no extra room for appending beyond its initial size.

		Below, we create a slice with:
		- Length: 5
		- Capacity: 10
		- The slice overlays only on the first 5 indexes of the underlying array.
		- The underlying array is allocated with size 10 and zero-initialized.
		- The type of the underlying array is `[10]int`, while the slice itself is `[]int`.

		NOTE:
		- The underlying array exists separately, and the slice is just a view into it.
		- The slice's capacity dictates how much can be appended before a new array needs to be allocated.
	*/

	newSlice := make([]int, 5, 10)

	// The initial slice state:
	fmt.Println("newSlice:", newSlice)      // Output: [0 0 0 0 0]
	fmt.Println("Length:", len(newSlice))   // Output: 5
	fmt.Println("Capacity:", cap(newSlice)) // Output: 10

	/*
		APPENDING ELEMENTS TO A SLICE

		- When elements are appended using `append(slice, element)`, Go determines whether there is
		  enough capacity in the existing underlying array.
		- If there is available capacity:
		  - The new element is added to the underlying array without changing its reference.
		- If capacity is exceeded:
		  - A **new, larger underlying array** is allocated.
		  - All elements from the old array are copied into the new array.
		  - The new element is added to the new array.
		  - The `append()` function **returns a new slice** that references this new array.

		IMPORTANCE OF STORING THE RETURNED SLICE:

		- Since `append()` can return a slice with a **different** underlying array, we must always
		  assign the result of `append()` back to our original slice variable.
		- If we fail to do so, we might continue using the old slice, which still references the old array.
	*/

	newSlice = append(newSlice, 6)        // Appending within capacity
	newSlice = append(newSlice, 7)        // Still within capacity
	newSlice = append(newSlice, 8, 9, 10) // Still within capacity
	newSlice = append(newSlice, 11)       // This exceeds capacity! A new array is allocated.

	/*
		At this point:
		- The first few `append()` calls used the original underlying array.
		- The final `append(11)` exceeded the capacity of 10.
		- Go created a **new** underlying array (size typically doubled, but depends on the runtime).
		- The old underlying array remains in memory but is no longer referenced by `newSlice`.
		- The new slice `newSlice` references the newly allocated array.
	*/

	fmt.Println("newSlice after appends:", newSlice)
	fmt.Println("Length after appends:", len(newSlice))
	fmt.Println("Capacity after appends:", cap(newSlice)) // Likely increased (usually doubled)

	/*
		COPYING SLICES

		- If we need to **copy** elements from one slice to another while ensuring they
		  do not share the same underlying array, we use `copy()`.
		- The syntax is: `copy(destination, source)`
		- This copies the **minimum of** `len(destination)` and `len(source)` elements.
		- This is useful when we need to create a completely **independent** slice.

		Below:
		- We create a new slice (`destSlice`) with the same length as `newSlice`.
		- `copy(destSlice, newSlice)` copies the values into `destSlice`.
		- `destSlice` is now independent and modifications to it do not affect `newSlice`.
	*/

	destSlice := make([]int, len(newSlice)) // New slice with the same length
	copy(destSlice, newSlice)               // Copy elements

	fmt.Println("Copied slice:", destSlice) // Independent copy of `newSlice`

	// a nil slice is a slice that has not been initialized. it's value is nil, and both
	// it's length and capacity are 0
	// it doesn't have an underlying array
	// it behaves like an empty slice in loops and append operations
	// appending an element to a nil slice works fine; Go automatically allocates an underlying array
	// once appended, it's value is no longer nil

	var nilSlice [](int) // no initialized, so it's nil
	_ = nilSlice

	// an empty slice is a valid slice that references an existing (but empty) underlying array
	// the slice is not nil
	// it has length 0 and capacity 0, but it does have an underlying array (even though it's empty)
	// it behaves like a normal slice and supoorts appends
	// unlike a nil slice, it explicitly exists in memory

	emptySlice := []int{}
	_ = emptySlice

	srcSlice := []int{1, 2, 3}
	fmt.Println(len(srcSlice), cap(srcSlice))

	matrix := createMatrix(10, 10)
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Printf("\n")
	}

	/*

		In Go, the range keyword is used to iterate over slices, arrays, maps, and channels.
		When iterating over a slice with range, it returns both the index and value of each element

	*/

	mySlice := []int{10, 20, 30, 40, 50, 60}

	// Iterate over mySlice using range.
	// 'index' represents the current position in the slice.
	// 'val' holds the corresponding value at that index.
	// The loop automatically iterates from index 0 to len(mySlice) - 1.
	for index, val := range mySlice {
		fmt.Printf("Element at index %d: Value = %d\n", index, val)
	}

	// to ignore the index
	for _, val := range mySlice {
		_ = val
	}

	// to ignore the value
	for index, _ := range mySlice {
		_ = index
	}

	// we can't ignore both

	/*

		A map in Go is a built-in data type that represents a hash table or dictionary. It
		stores key-value pairs, where each key is unique and maps to a specific value. Maps
		provide fast lookups, additions, and deletions due to their underlying hash table
		implementation.

		Maps are declared using the map keyword

		var myMap map[keyType]valueType

		keyType defines the type of keys (must be comparable types like string, int, etc)
		valueType defines the type of values

	*/

	// there are multiple ways to declare and initialize a map

	myMap := make(map[string]int) // create an empty map with string keys and int values
	_ = myMap

	// declaring and initializing a map with predefined key-value pairs
	personAge := map[string]int{
		"raj":     19,
		"rishika": 17,
	}
	_ = personAge

	// declaring a map without initializing it leads to a nil map
	var mySecondMap map[string]int // nil map (cannot be used until initialized)
	// to use it, initialize it, with make for ex,
	mySecondMap = make(map[string]int)
	_ = mySecondMap

	// we can add key-value pairs (to non nil maps) and access value using their keys (again in non-nil maps)
	personAge["pranav"] = 20
	personAge["prince"] = 18
	fmt.Println(personAge["raj"]) // prints 19

	// to check if a key exists, use the comma ok idiom
	// exists is true if the key exists, false otherwise
	// if key is not in the map, then elem is the zero value for the map's value type
	elem, exists := personAge["ritesh"]
	if exists {
		fmt.Println("Ritesh's age:", elem)
	} else {
		fmt.Println("We don't know Ritesh's age")
	}
	// if the key doesn't exist, age will be the zero value of the value type (0 for int)

	// we can remove a key-value pair using delete() function
	// deleting a non-existent key does nothing
	delete(personAge, "raj")
	fmt.Println(personAge)

	// we can use a for range loop to iterate over key-value pairs
	for key, value := range personAge {
		fmt.Println("key:", key, "value:", value)
	}

	// we can also iterate over keys only
	for key := range personAge {
		fmt.Println(key)
	}

	// maps are reference types, meaning they share underlying data when assigned
	copyPersonAge := personAge // both point to the same underlying map
	_ = copyPersonAge

	// to create a true copy, manually copy elements
	copyPersonAge = make(map[string]int)
	for key, value := range personAge {
		copyPersonAge[key] = value
	}
	// or better
	maps.Copy(copyPersonAge, personAge)

	// the len function on a map returns the number of key-value pairs the map contains; 0 if the
	// map is nil
	fmt.Println(len(personAge))

	/*

		Maps in Go are unordered, iteration order is random
		Avoid modifying a map while iterating over it

	*/

	/*

		Like slices, maps are also passed by reference into functions. This means
		that when a map is passed into a function we write, we can make changes
		to the original, we don't have a copyu

	*/

	changeMap(personAge)
	fmt.Println(personAge)

	/*

		In Go, keys in a map must be of a comparable type. This means types that
		support the ==, != operations.

		Go does not allow types that are not comparable, including slices, maps, functions, and
		structs the contain non-comparable fields

		Structs containing only comparable fields can be used as keys into a map

	*/

	type name_t struct {
		firstName string
		lastName  string
	}

	nameWorthMap := make(map[name_t]int)
	nameWorthMap[name_t{firstName: "raj", lastName: "shukla"}] = 10
	nameWorthMap[name_t{"rishika", "rajoriya"}] = 100

	fmt.Println(nameWorthMap)

	/*

		Like slices, maps holds references to the underlying data structure.
		If you pass a map to a function that changes the contents of the map, the
		changes will be visible in the caller.

	*/

	// maps can be constructed using the usual composite literal syntax
	// with colon-separated key-value pairs
	var nameWeight = map[string]int{
		"raj":     80,
		"rishika": 60,
	}
	_ = nameWeight

	// an attempt to fetch a map value with a key that is not present in the map
	// will return the zero value for the type of entries in the map
	attended := map[string]bool{
		"raj":     true,
		"rishika": true,
		"aryaman": true,
	}

	if !attended["zinnia"] { // will be true if person is not in the map
		fmt.Println("zinnia was not at the meeting")
	}

	fmt.Println(aggregate(10, 9, 8, add))  // will print 27
	fmt.Println(aggregate(10, 9, 8, mult)) // will print 720

	// Go supports first-class functions, meaning functions can be:
	// assigned to variables
	// passed as arguments
	// returned from other functions

	// a higher order function is a function that
	// takes another function as an argument, or
	// returns a function

	squareFunc := selfMath(mult)
	fmt.Println(squareFunc(12))

	doubleFunc := selfMath(add)
	fmt.Println(doubleFunc(12))

	processFile("example.txt")

	firstIncrement := counter() // returns a closure
	fmt.Println(firstIncrement())
	fmt.Println(firstIncrement())
	fmt.Println(firstIncrement())

	secondIncrement := counter() // returns another closure capturing a different copy of count
	fmt.Println(secondIncrement())
	fmt.Println(secondIncrement())
	fmt.Println(secondIncrement())

	/*
		A pointer in Go is a variable that stores the memory address of another variable.
		Pointers allow functions and methods to modify the original value instead of working with a copy,
		making them useful for efficiency and state modification.

		### Pointer Receivers in Methods:
		- A method with a pointer receiver (`*Type`) operates on the original struct instance.
		- This enables modifications to struct fields and avoids unnecessary copying, which is
		  especially beneficial when dealing with large structs.

		### Implicit Dereferencing:
		- Go automatically dereferences pointers when accessing their fields or calling methods.
		- For example, if `teddy` is a pointer, both `(*teddy).color` and `teddy.color` are valid ways to access the `color` field.
		- This shorthand (implicit dereferencing) helps avoid verbose code.

		### Benefits of Using Pointers in Go:
		✅ Modify values indirectly by passing references instead of copying data.
		✅ Efficient memory usage, reducing overhead when dealing with large structs.
		✅ Enable optional values by allowing `nil` pointers.
		✅ Support shared state, where multiple functions can operate on the same data.

		### Limitations of Pointers in Go:
		❌ Dereferencing a `nil` pointer causes a runtime panic.
		❌ No pointer arithmetic, unlike C/C++, preventing manual memory manipulation.
		❌ No manual memory allocation (`malloc/free`); Go uses garbage collection.
		❌ Must explicitly dereference (`*ptr`) to access a pointer's stored value when working with the pointer itself.

		### Summary of Pointer Behavior:
		- Go automatically dereferences pointers for field access and method calls (implicit dereferencing).
		- This simplifies working with pointers but still requires explicit dereferencing when performing pointer-based operations.
		- The safety and simplicity of Go’s pointer system help prevent common memory-related bugs, while providing efficient memory usage and the ability to modify data.

		Go's pointer system provides safety and simplicity while ensuring memory efficiency, but it enforces strict rules to avoid common memory-related bugs.
	*/

	var teddy teddy_t = teddy_t{"white", 100}
	teddy.setColor("blue")
	fmt.Println(teddy.color)

	fmt.Println(getSquare(16))
}

type teddy_t struct {
	color  string
	volume int
}

func (teddy *teddy_t) setColor(color string) {
	teddy.color = color // same as (*teddy).color due to implicit pointer dereferencing
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

// a function that takes another function as its parameter
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	val := arithmetic(a, b)
	return arithmetic(val, c)
}

/*
Currying is a technique where a function that takes multiple arguments
is transformed into a sequence of functions, each taking a single argument
*/
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

/*
A closure is an anonymous function that captures and retains access to variables
from it's surrounding scope, even after the surrounding function has returned
*/
func counter() func() int {
	// counter returns a function func() int
	// the returned function captures count, even though counter has returned
	count := 0 // captured variable
	return func() int {
		count++ // modifies the captured variable
		return count
	}
}

var mu sync.Mutex

func processFile(filename string) (success bool) {
	// the defer keyword in Go is used to postpone the execution of a function
	// call just before the surrounding function returns.

	// when a function call is deferred:

	// the function itself does not execute immediately

	// it's arguments are evaluated immediately, but the execution of the
	// function body is delayed

	// deferred functions execute in LIFO (Last-In-First-Out) order, meaning the
	// last deferred function runs first

	// arguments of the deferred function are evaluated at the time of defer, not
	// when the function actually executes

	// when multiple defer statements exist, they execute in reverse order of their
	// declaration

	// when a function panics, deferred calls still execute before the program crashes
	// if a function returns named return value, deferred functions can modify it

	// if a function returns in between, only the defers registered till that point
	// will execute for that return

	success = false
	defer func() {
		if !success {
			fmt.Println("Processing Failed, Cleaning Up")
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Opening File: ", err)
		return success
	}

	defer func() {
		fmt.Println("Closing File:", filename)
		file.Close()
	}()

	// lock a mutex for safe access to shared resources

	mu.Lock()
	defer func() {
		fmt.Println("Releasing mutex lock")
		mu.Unlock()
	}()

	fmt.Println("Processing file")

	// some processing
	success = true
	return success
}
