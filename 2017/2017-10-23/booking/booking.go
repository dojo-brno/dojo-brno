package booking

import (
	"fmt"
	"time"
)

type Event struct {
	StartTime, EndTime time.Time
	Room               string
}

func CanInsert(newEvent Event, listOfEvents []Event) error {
	if err := validateEvent(newEvent); err != nil {
		return err
	}
	for _, event := range listOfEvents {
		if err := collidingEvents(event, newEvent); err != nil {
			return err
		}
	}
	return nil
}

func collidingEvents(existingEvent, newEvent Event) error {
	if existingEvent.Room != newEvent.Room {
		return nil
	}
	if newEvent.StartTime.After(existingEvent.EndTime) || newEvent.StartTime.Equal(existingEvent.EndTime) {
		return nil
	}
	if newEvent.EndTime.Before(existingEvent.StartTime) || newEvent.EndTime.Equal(existingEvent.StartTime) {
		return nil
	}
	return collidingEvent
}

func validateEvent(event Event) error {
	if event.Room == "" {
		return invalidRoom
	}
	if event.StartTime.Equal(event.EndTime) {
		return equalStartAndEndTime
	}
	if event.EndTime.Before(event.StartTime) {
		return endsBeforeStarts
	}
	return nil
}

var (
	collidingEvent       = fmt.Errorf("colliding event")
	equalStartAndEndTime = fmt.Errorf("start and end times must not be equal")
	endsBeforeStarts     = fmt.Errorf("event finishes before its start")
	invalidRoom          = fmt.Errorf("invalid room")
)
