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
}
