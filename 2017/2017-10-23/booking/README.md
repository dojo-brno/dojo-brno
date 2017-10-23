# Exercise

Booking System

Implement:
// CanInsert returns:
// nil - in case the newEvent can be inserted into the ListOfEvents without collision with another event
// error - in case the newEvent collides with an event in the ListOfEvents
func CanInsert(newEvent Event, ListOfEvents []Event) error

type Event {
  StartTime, EndTime time
  Room string
}
