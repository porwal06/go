package modules

import "time"

// Create event struct
// `binding:"required"` meta tag of struct make it mandatory fields
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// Create empty event slice for storing events
var events = []Event{}

// Add save event method to Event struct
func (e Event) Save() {
	events = append(events, e)
}

// Get all events
func GetAllEvents() []Event {
	return events
}
