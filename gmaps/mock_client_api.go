// +build mock

package gmaps

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) Autocomplete(ctx context.Context, input string) ([]maps.QueryAutocompletePrediction, error) {
	key := mockKey("Autocomplete", input)

	if v, err := readMock(key); err == nil {
		res, _ := v[0].([]maps.QueryAutocompletePrediction)
		err, _ := v[1].(error)
		return res, err
	}

	res, err := c.doAutocomplete(ctx, input)
	writeMock(key, res, err)
	return res, err
}

func (c *MapsApiClient) Details(ctx context.Context, placeID string) (*maps.PlaceDetailsResult, error) {
	key := mockKey("Details", placeID)

	if v, err := readMock(key); err == nil {
		res, _ := v[0].(*maps.PlaceDetailsResult)
		err, _ := v[1].(error)
		return res, err
	}

	res, err := c.doDetails(ctx, placeID)
	writeMock(key, res, err)
	return res, err
}

func (c *MapsApiClient) ReverseGeocode(ctx context.Context, lat, lng float64) ([]maps.GeocodingResult, error) {
	key := mockKey("ReverseGeocode", lat, lng)

	if v, err := readMock(key); err == nil {
		res, _ := v[0].([]maps.GeocodingResult)
		err, _ := v[1].(error)
		return res, err
	}

	res, err := c.doReverseGeocode(ctx, lat, lng)
	writeMock(key, res, err)
	return res, err
}
