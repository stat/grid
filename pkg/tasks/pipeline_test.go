package tasks_test

import (
	"testing"
	"time"

	"grid/pkg/models"
	"grid/pkg/tasks/consumer"
	"grid/pkg/tasks/producer"
	"grid/pkg/utils"

	"github.com/stretchr/testify/assert"

	_ "grid/testing"
)

func TestPipeline(t *testing.T) {
	aircraftID := "aircraft-01"
	stationID := "station-01"

	start := time.Now()

	tests := []struct {
		Event    *models.ADSB
		Expected error
	}{
		{
			Event: &models.ADSB{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start),
			},
			Expected: nil,
		},
		{
			Event: &models.ADSB{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start),
			},
			Expected: consumer.ProcessEventTimestampEqualError,
		},
		{
			Event: &models.ADSB{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start.Add(time.Second)),
			},
			Expected: consumer.ProcessEventTimestampAfterError,
		},
		{
			Event: &models.ADSB{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start.Add(time.Second * -1)),
			},
			Expected: consumer.ProcessEventTimestampBeforeError,
		},
		{
			Event: &models.ADSB{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  nil,
			},
			Expected: models.ADSBValidateTimestampError,
		},
	}

	// process

	for _, test := range tests {
		event := test.Event

		// consumer

		err := consumer.Process(event)

		// assert

		assert.Equal(t, test.Expected, err)

		if test.Expected != nil {
			continue
		}

		// producer

		err = producer.Process(event)

		assert.Equal(t, nil, err)
	}
}
