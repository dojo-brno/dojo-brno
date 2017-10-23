package booking

import (
	"testing"
	"time"
)

func TestCanInsert(t *testing.T) {
	tests := []struct {
		testName     string
		newEvent     Event
		listOfEvents []Event
		want         error
	}{
		{
			testName:     "Can add an event into an empty list of events",
			newEvent:     Event{StartTime: noon, EndTime: onePM, Room: "A"},
			listOfEvents: []Event{},
			want:         nil,
		},
		{
			testName:     "Added event collides with the first existing event",
			newEvent:     Event{StartTime: noon, EndTime: onePM, Room: "A"},
			listOfEvents: []Event{Event{StartTime: noon, EndTime: onePM, Room: "A"}},
			want:         collidingEvent,
		},
		{
			testName:     "Can add new event after an existing one",
			newEvent:     Event{StartTime: onePM, EndTime: twoPM, Room: "A"},
			listOfEvents: []Event{Event{StartTime: noon, EndTime: onePM, Room: "A"}},
			want:         nil,
		},
		{
			testName:     "Can add new event before an existing one",
			newEvent:     Event{StartTime: elevenAM, EndTime: noon, Room: "A"},
			listOfEvents: []Event{Event{StartTime: noon, EndTime: onePM, Room: "A"}},
			want:         nil,
		},
		{
			testName:     "Start and end time of the new event equal",
			newEvent:     Event{StartTime: elevenAM, EndTime: elevenAM, Room: "A"},
			listOfEvents: []Event{Event{StartTime: noon, EndTime: onePM, Room: "A"}},
			want:         equalStartAndEndTime,
		},
		{
			testName:     "End time of the new event is before its start time",
			newEvent:     Event{StartTime: noon, EndTime: elevenAM, Room: "A"},
			listOfEvents: []Event{Event{StartTime: noon, EndTime: onePM, Room: "A"}},
			want:         endsBeforeStarts,
		},
		{
			testName:     "Invalid room in the new event",
			newEvent:     Event{StartTime: noon, EndTime: onePM, Room: ""},
			listOfEvents: []Event{},
			want:         invalidRoom,
		},
		{
			testName: "New added event collides with the second existing one",
			newEvent: Event{StartTime: noon, EndTime: onePM, Room: "A"},
			listOfEvents: []Event{
				Event{StartTime: elevenAM, EndTime: noon, Room: "A"},
				Event{StartTime: noon, EndTime: onePM, Room: "A"},
			},
			want: collidingEvent,
		},
		{
			testName: "Can add new event to another room",
			newEvent: Event{StartTime: noon, EndTime: onePM, Room: "B"},
			listOfEvents: []Event{
				Event{StartTime: elevenAM, EndTime: noon, Room: "A"},
				Event{StartTime: noon, EndTime: onePM, Room: "A"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		if got := CanInsert(tt.newEvent, tt.listOfEvents); got != tt.want {
			t.Errorf("%v: CanInsert(%v, %v) = %v WANT %v", tt.testName, tt.newEvent, tt.listOfEvents, got, tt.want)
		}
	}
}

var (
	elevenAM = time.Date(2017, 10, 23, 11, 0, 0, 0, time.FixedZone("CET", 0))
	noon     = time.Date(2017, 10, 23, 12, 0, 0, 0, time.FixedZone("CET", 0))
	onePM    = time.Date(2017, 10, 23, 13, 0, 0, 0, time.FixedZone("CET", 0))
	twoPM    = time.Date(2017, 10, 23, 14, 0, 0, 0, time.FixedZone("CET", 0))
)
