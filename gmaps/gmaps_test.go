package gmaps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	apiKey = "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"
)

func TestDetails(t *testing.T) {
	api, err := NewMapsClient(apiKey)
	assert.NoError(t, err)
	place, err := api.Details("ChIJL6wn6oAOZ0gRoHExl6nHAAo")
	assert.NoError(t, err)
	assert.NotNil(t, place)
	t.Logf("place: %v ", place)
}

func TestAutocomplete(t *testing.T) {
	api, err := NewMapsClient(apiKey)
	assert.NoError(t, err)
	predictions, err := api.Autocomplete("dublin")
	assert.NoError(t, err)
	assert.True(t, len(predictions) > 0)
	for _, p := range predictions {
		t.Logf("%v", p)
	}
}

func TestReverseGeocode(t *testing.T) {
	api, err := NewMapsClient(apiKey)
	assert.NoError(t, err)
	places, err := api.ReverseGeocode(53.339897, -6.538458899999999)
	assert.NoError(t, err)
	assert.True(t, len(places) > 0)
	for _, p := range places {
		t.Logf("%v", p)
	}
}
