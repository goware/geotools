package gmaps

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

const (
	apiKey = "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"
)

func TestDetails(t *testing.T) {
	api, err := NewMapsClient(apiKey)
	assert.NoError(t, err)
	ctx := context.Background()
	place, err := api.Details(ctx, "ChIJL6wn6oAOZ0gRoHExl6nHAAo")
	assert.NoError(t, err)
	assert.NotNil(t, place)
	t.Logf("place: %v ", place)
}

func TestAutocomplete(t *testing.T) {
	api, err := NewMapsClient(apiKey)
	assert.NoError(t, err)
	ctx := context.Background()
	predictions, err := api.Autocomplete(ctx, "dublin")
	assert.NoError(t, err)
	assert.True(t, len(predictions) > 0)
	for _, p := range predictions {
		t.Logf("%v", p)
	}
}

func TestReverseGeocode(t *testing.T) {
	api, err := NewMapsClient(apiKey)
	assert.NoError(t, err)
	ctx := context.Background()
	places, err := api.ReverseGeocode(ctx, 53.339897, -6.538458899999999)
	assert.NoError(t, err)
	assert.True(t, len(places) > 0)
	for _, p := range places {
		t.Logf("%v", p)
	}
}
