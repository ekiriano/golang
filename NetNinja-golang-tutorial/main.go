package main

import "fmt" // fmt is a package from the package library.
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

}

func main() { // main function
	// lessonOneVariables()
	// lessonTwoPrintingAndFormattingStrings()
	// lessonThreeArraysAndSlices()
	lessonFourTheStandardLibray()
}
