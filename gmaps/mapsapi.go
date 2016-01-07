// +build !mock

package gmaps

import (
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) Autocomplete(input string) ([]maps.QueryAutocompletePrediction, error) {
	return c.doAutocomplete(input)
}

func (c *MapsApiClient) Details(placeID string) (*maps.PlaceDetailsResult, error) {
	return c.doDetails(placeID)
}

func (c *MapsApiClient) ReverseGeocode(lat, lng float64) ([]maps.GeocodingResult, error) {
	return c.doReverseGeocode(lat, lng)
}
