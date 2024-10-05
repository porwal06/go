package note

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note struct with json metadata
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New() Note {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note title: ")
	noteTitle, _ := reader.ReadString('\n')

	fmt.Print("Enter note content: ")
	noteContent, _ := reader.ReadString('\n')

	return Note{
		Title:     strings.ReplaceAll(noteTitle, "\n", ""),
		Content:   strings.ReplaceAll(noteContent, "\n", ""),
		CreatedAt: time.Now(),
	}
}

func (n Note) Print() {
	fmt.Printf("Note title: %v\nNote Content :\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	if n.Title == "" || n.Content == "" {
		return fmt.Errorf("please provide valid note data")
	}
	// Create json file
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	// Create file from note title after replacing spaces with underscores & lower letter
	fileName := strings.ReplaceAll(strings.ToLower(n.Title), " ", "_") + ".json"

	os.WriteFile(fileName, json, 0644)
	return nil
}
