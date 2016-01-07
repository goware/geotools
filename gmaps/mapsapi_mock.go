// +build mock

package gmaps

import (
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) Autocomplete(input string) ([]maps.QueryAutocompletePrediction, error) {
	if v, err := readMock(mockKey("Autocomplete", input)); err == nil {
		res, _ := v[0].([]maps.QueryAutocompletePrediction)
		err, _ := v[1].(error)
		return res, err
	}
	return nil, errNoMockData
}

func (c *MapsApiClient) Details(placeID string) (*maps.PlaceDetailsResult, error) {
	if v, err := readMock(mockKey("Details", placeID)); err == nil {
		res, _ := v[0].(*maps.PlaceDetailsResult)
		err, _ := v[1].(error)
		return res, err
	}
	return nil, errNoMockData
}

func (c *MapsApiClient) ReverseGeocode(lat, lng float64) ([]maps.GeocodingResult, error) {
	if v, err := readMock(mockKey("ReverseGeocode", lat, lng)); err == nil {
		res, _ := v[0].([]maps.GeocodingResult)
		err, _ := v[1].(error)
		return res, err
	}
	return nil, errNoMockData
}
