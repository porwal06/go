package modules

import (
	"time"

	"example.com/rest-api/db"
)

// Create event struct
// `binding:"required"` meta tag of struct make it mandatory fields
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// Create empty event slice for storing events
var events = []Event{}

// Add save event method to Event struct
func (e Event) Save() error {
	//Save event in sqllite database
	query := `INSERT INTO events (title, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	//Error handling
	if err != nil {
		return err
	}
	defer stmt.Close() // defer will make sure to close only after exection is done. It's position doesn't matter.
	// Use Exec() method for Insert, update, delete
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	//events = append(events, e)
	return err
}

// Get all events
func GetAllEvents() ([]Event, error) {
	//Use Query method to get all events from database. Exec() can also be used but it for fetching Query is recommened
	var singleEvent Event
	result, err := db.DB.Query("SELECT * FROM events")
	//Error handling
	if err != nil {
		return nil, err
	}
	for result.Next() {

		err = result.Scan(&singleEvent.ID, &singleEvent.Name, &singleEvent.Description, &singleEvent.Location, &singleEvent.DateTime, &singleEvent.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, singleEvent)
	}
	defer result.Close()
	return events, err
}
