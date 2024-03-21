package grid

import (
	"fmt"
	"time"
)

const (
	ReconsilationWindow = 2
)

var (
	Index                      map[string][]*Payload = make(map[string][]*Payload)
	AppendEventAircraftIDError                       = fmt.Errorf("must include aircraft id!")
)

type Payload struct {
	AircraftID string
	Latitude   float64
	Longitude  float64
	StationID  string
	Timestamp  *time.Time
}

func main() {
	fmt.Println("hello, world")

	// append events
}

func AppendEvent(payload *Payload, reconsileEvents bool) error {
	aircraftID := payload.AircraftID

	// sanity check

	fmt.Println(payload)
	if aircraftID == "" {
		return AppendEventAircraftIDError
	}

	if payload.StationID == "" {
		return fmt.Errorf("must include station id!")
	}

	if payload.Timestamp == nil {
		return fmt.Errorf("must include timestamp!")
	}

	// append

	data := append(Index[payload.AircraftID], payload)

	// check to see if we need to reconsile

	if reconsileEvents == true && len(data) > ReconsilationWindow {
		events, err := ReconsileEvents(aircraftID)

		if err != nil {
			return err
		}

		// flush
		data = []*Payload{}

		err = storeEvents(events)

		if err != nil {
			return err
		}

	}

	// ensure index update

	Index[payload.AircraftID] = data

	// success

	return nil
}

// returns list of reconsiled events
func ReconsileEvents(aircraftID string) ([]*Payload, error) {
	// get all data for aircraftid window

	data := Index[aircraftID]

	if data == nil {
		return nil, fmt.Errorf("no data available for %s", aircraftID)
	}

	events := []*Payload{}

	var last *Payload = nil
	for _, item := range data {
		if isEventEqual(last, item) {
			continue
		}

		events = append(events, item)
		last = item
	}

	return events, nil
}

func isEventEqual(e1, e2 *Payload) bool {
	if e1 == nil || e2 == nil {
		return false
	}

	return isTimeEqual(e1.Timestamp, e2.Timestamp)
}
func isTimeEqual(t1, t2 *time.Time) bool {
	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Sub(*t2) == 0
}
func storeEvents(events []*Payload) error {
	return nil
}
