package gmaps

import (
	"testing"
)

const (
	apiKey = "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"
)

func TestDetails(t *testing.T) {
	api := NewMapsClient(apiKey)
	place, err := api.Details("ChIJL6wn6oAOZ0gRoHExl6nHAAo")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("place: %v ", place)
}

/*
func TestAutocomplete(t *testing.T) {
	api := &MapsApiClient{Key: "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"}
	predictions, err := api.Autocomplete("dublin")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for _, p := range predictions {
		t.Logf("%v", p)
	}
}

func TestReverseGeocode(t *testing.T) {
	api := &MapsApiClient{Key: "AIzaSyAwoYYcg8R4K91Sc8fim3hw7OPe48wX2RI"}
	places, err := api.ReverseGeocode(53.339897, -6.538458899999999)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	for _, p := range places {
		t.Logf("%v", p)
	}
}
*/
