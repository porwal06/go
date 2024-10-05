package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Content string
}

func New() Todo {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter TODO content: ")
	newContent, _ := reader.ReadString('\n')
	newContent = newContent[:len(newContent)-1] // Remove the newline character
	return Todo{Content: newContent}
}

func (t Todo) Print() {
	fmt.Printf("TODO content: %v\n\n", t.Content)
}

func (t Todo) Save() error {
	// Create json file
	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	os.WriteFile("todo.json", json, 0644)
	return nil
}
