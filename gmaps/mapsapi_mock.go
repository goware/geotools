// +build mock

package gmaps

import (
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) Autocomplete(input string) ([]maps.QueryAutocompletePrediction, error) {
	key := mockKey("Autocomplete", input)

	if v, err := readMock(key); err == nil {
		res, _ := v[0].([]maps.QueryAutocompletePrediction)
		err, _ := v[1].(error)
		return res, err
	}

	res, err := c.doAutocomplete(input)
	writeMock(key, res, err)
	return res, err
}

func (c *MapsApiClient) Details(placeID string) (*maps.PlaceDetailsResult, error) {
	key := mockKey("Details", placeID)

	if v, err := readMock(key); err == nil {
		res, _ := v[0].(*maps.PlaceDetailsResult)
		err, _ := v[1].(error)
		return res, err
	}

	res, err := c.doDetails(placeID)
	writeMock(key, res, err)
	return res, err
}

func (c *MapsApiClient) ReverseGeocode(lat, lng float64) ([]maps.GeocodingResult, error) {
	key := mockKey("ReverseGeocode", lat, lng)

	if v, err := readMock(key); err == nil {
		res, _ := v[0].([]maps.GeocodingResult)
		err, _ := v[1].(error)
		return res, err
	}

	res, err := c.doReverseGeocode(lat, lng)
	writeMock(key, res, err)
	return res, err
}
