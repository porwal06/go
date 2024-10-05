package main

import (
	"example.com/note/note"
	"example.com/note/todo"
)

// Create interface to print and save data
type Printer interface {
	Print()
	Save() error
}

func PrintAndSave(data Printer) {
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
}
