package main

import (
	"fmt"

	"example.com/note/note"
	"example.com/note/todo"
)

// Create interface to print and save data
// We can add multiple methods or Interface also for creating combine interface. It's called embeded interface
type outputable interface {
	Print()
	Save() error
}

func PrintAndSave(data outputable) {
	data.Print()
	data.Save()
}

// Main function without implementing Interface

/*func main() {
	todoDetail := todo.New()
	todoDetail.Print()
	todoDetail.Save()

	noteDetail := note.New()
	noteDetail.Print()
	noteDetail.Save()
}*/

// Main function with Interface implementation
func main() {
	todoDetail := todo.New()
	PrintAndSave(todoDetail)

	noteDetail := note.New()
	PrintAndSave(noteDetail)

	// Testing of generic type function
	result := addAnthing(10, 20)
	fmt.Println(result)

	resultString := addAnthing("Add", "Anything")
	fmt.Println(resultString)
}

// Generic type function
func addAnthing[T int | float64 | string](a T, b T) T {
	return a + b
}
