package gmaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

const (
	apiKey = "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"
)

func testMapsClient(apiKey string) (MapsApiClient, error) {
	//return NewMapsClient(apiKey)
	return newMockMapsClient(apiKey)
}

func TestDetails(t *testing.T) {
	api, err := testMapsClient(apiKey)
	assert.NoError(t, err)
	place, err := api.Details(context.Background(), "ChIJL6wn6oAOZ0gRoHExl6nHAAo")
	assert.NoError(t, err)
	assert.NotNil(t, place)
	t.Logf("place: %v ", place)
}

func TestAutocomplete(t *testing.T) {
	api, err := testMapsClient(apiKey)
	assert.NoError(t, err)
	predictions, err := api.Autocomplete(context.Background(), "dublin")
	assert.NoError(t, err)
	assert.True(t, len(predictions) > 0)
	for _, p := range predictions {
		t.Logf("%v", p)
	}
}

func TestTextSearch(t *testing.T) {
	api, err := testMapsClient(apiKey)
	assert.NoError(t, err)
	places, err := api.TextSearch(context.Background(), "Toronto")
	assert.NoError(t, err)
	assert.True(t, len(places) > 0)
	for _, p := range places {
		t.Logf("%v", p)
	}
}

func TestReverseGeocode(t *testing.T) {
	api, err := testMapsClient(apiKey)
	assert.NoError(t, err)
	places, err := api.ReverseGeocode(context.Background(), 53.339897, -6.538458899999999)
	assert.NoError(t, err)
	assert.True(t, len(places) > 0)
	for _, p := range places {
		t.Logf("%v", p)
	}
}
