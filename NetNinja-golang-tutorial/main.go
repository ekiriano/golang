package main

import (
	"fmt" // fmt is a package from the package library.
	"strings"
)

// fmt is for formating strings and printing to the console.

func lessonOneVariables() {
	// strings
	var nameOne string = "Mario"
	var nameTwo = "Luigi"
	var nameThree string // Empty string is the default value
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowser"

	fmt.Println(nameOne, nameTwo, nameThree)
	fmt.Println("Hello ninjas!")

	nameFour := "Yoshi" // shorthand version , cannot use this syntax outside of a function.

	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory

	var numOne int8 = 25 // int8 -128 to 127.
	var numTwo int16 = 128
	var numThree uint8 = 25 // only positive number 0 to 255;

	var scoreOne float32 = 2.444 // for the most part we'll use float64.

	fmt.Println(numOne, numTwo, numThree, scoreOne)
}

func lessonTwoPrintingAndFormattingStrings() {
	age := 27
	name := "Amine"
	//  Print
	fmt.Print("hello, ")  // No new line
	fmt.Print("world!\n") // \n create new line.

	// Println
	fmt.Println("Hello world!")   // creates new line
	fmt.Println("Goodbye world!") // creates new line
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (Formatted strings)
	fmt.Printf("my age is %v and my name is %q \n", age, name) // %v => variable , %q => places double quotes

	// https://pkg.go.dev/fmt
	//	%v	the value in a default format
	//	when printing structs, the plus flag (%+v) adds field names
	//
	// %#v	a Go-syntax representation of the value
	// %T	a Go-syntax representation of the type of the value
	// %%	a literal percent sign; consumes no value
	// %t	the word true or false

	// Sprintf (save formatted strings)
	str := fmt.Sprintf("my age is %v and my name is %q \n", age, name)
	fmt.Printf(str)
}

func lessonThreeArraysAndSlices() {
	var ages [3]int = [3]int{20, 25, 30} // it has to be fixed length.
	var ages2 = [2]int{20, 25}           // type is inferred.

	fmt.Println(ages, len(ages))
	fmt.Println(ages2, len(ages2))

	// slices ( use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores = append(scores, 3) // append return a new slice for us , so we have to do an assignment.
	fmt.Println(scores, len(scores))

	// slice ranges (get a range of elements to store in a new slice)
	rangeOne := scores[1:3] // inclusive of the first number but exclusive of the last number.
	fmt.Println(rangeOne)

	rangeTwo := scores[1:] // goes up to the end of the slice
	fmt.Print(rangeTwo)

	rangeThree := scores[:3] //from the star to 3 not including
	fmt.Println(rangeThree)
}

func lessonFourTheStandardLibray() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "friends", "ninjas"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "h"))
	fmt.Println(strings.Split(greeting, " "))
}

func lessonFiveLoops() {
	x := 0

	for x < 5 {
		fmt.Println("the value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("the value of i is:", i)
	}

	names := []string{"mario", "luigi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("Index: %v , name: %v \n", index, value)
		value = "test" // doesnt update the real value inside the original slice
	}
}

func lessonSixBooleansAndConditionals() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("I'm under 30")
	} else if age < 40 {
		fmt.Println("age is elss than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {

		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		fmt.Println("index,val", index, value)

		if index == 3 {
			fmt.Println("stopping at pos", index)
			break
		}
	}

}

func sayGreeting(name string) {
	fmt.Println("Hello, ", name)
}

func sayBye(name string) {
	fmt.Println("Goodbye, ", name)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func lessonSevenUsingFunctions() {
	sayGreeting("Amine")
	sayBye("Amine")
	cycleNames([]string{"one", "two", "three"}, sayGreeting)
}

func LessonEightMultipleReturnValue() (int, int) {
	return 1, 2
}

func lessonNineMaps() {
	// a bit like objects in javascript , all the keys in the single map must have the same type , same for the values

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967:   "mario",
		2675823237:  "luigi",
		26758223327: "bowser",
	}

	fmt.Println(phonebook)
}

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 5.99
}

func LessonTenPassByValue() {
	// Go is known as a pass by value language,
	// go makes "copies" of values when passed into functions.
	//
	// Group A: strings, ints , floats , booleans , arrays , string (non pointer values)
	// Group B : slices , maps , functions (pointer Wrapper values)

	name := "tifa"
	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	updateMenu(menu)
	fmt.Println(menu)
}

func updateNameMem(x *string) {
	*x = "wedge"
}

func lessonElevenPointers() {
	// pointers point to a memory location , it is also a type.

	name := "tifa"
	m := &name // gets the memory addresss with &
	fmt.Println("Memory address of name is: ", m)
	fmt.Println("value at memory address", *m)

	updateNameMem(&name)

	fmt.Println(name)
}

func main() { // main function
	// lessonOneVariables()
	// lessonTwoPrintingAndFormattingStrings()
	// lessonThreeArraysAndSlices()
	// lessonFourTheStandardLibray()
	// lessonFiveLoops()
	// lessonSixBooleansAndConditionals()
	// lessonSevenUsingFunctions()
	// LessonNineMaps()
	// LessonTenPassByValue()
	lessonElevenPointers()
}
