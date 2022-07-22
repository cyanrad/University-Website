package db

import (
	"context"
	"testing"
	"time"

	"github.com/cyanrad/university/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEvent(t *testing.T) Event {
	// >> events random parameters
	arg := CreateEventParams{
		Name:        util.RandomString(20),
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour * 24),
		Link:        util.RandomLink(),
		Description: util.RandomString(300),
	}

	// >> doing the query
	event, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	// >> checking data correctness
	require.Equal(t, arg.Name, event.Name)
	require.WithinDuration(t, arg.StartDate, event.StartDate, time.Hour*24)
	require.WithinDuration(t, arg.EndDate, arg.EndDate, time.Hour*24)
	require.Equal(t, arg.Link, event.Link)
	require.Equal(t, arg.Description, event.Description)

	return event
}

// >> testing getting of the count of current events
// TODO: only retrieve the events that are not expired
func TestGetEventCount(t *testing.T) {
	// >> getting current count
	previous, err := testQueries.GetEventCount(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, previous, int64(0))

	// >> creating 3 new events
	for i := 0; i < 3; i++ {
		CreateRandomEvent(t)
	}

	// >> checking if the new count is increased by 3
	current, err := testQueries.GetEventCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, previous, current-3)
}

// >> testing the creation of an event
func TestCreateEvent(t *testing.T) {
	CreateRandomEvent(t)
}

// >> testing the getting of a single event
func TestGetEvent(t *testing.T) {
	// >> creating the row, and getting it's location
	created := CreateRandomEvent(t)
	count, err := testQueries.GetEventCount(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(0))

	// >> retrieving the row
	retrieved, err := testQueries.GetEvent(context.Background(), int32(count-1))
	require.NoError(t, err)
	require.NotEmpty(t, retrieved)

	// >> checking data correctness
	require.Equal(t, created.Name, retrieved.Name)
	require.WithinDuration(t, created.StartDate, retrieved.StartDate, time.Hour*24)
	require.WithinDuration(t, created.EndDate, created.EndDate, time.Hour*24)
	require.Equal(t, created.Link, retrieved.Link)
	require.Equal(t, created.Description, retrieved.Description)
}
