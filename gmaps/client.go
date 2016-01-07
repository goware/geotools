package gmaps

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) doAutocomplete(input string) ([]maps.QueryAutocompletePrediction, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.QueryAutocompleteRequest{
		Input: input,
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.client.QueryAutocomplete(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res.Predictions, err
}

func (c *MapsApiClient) doDetails(placeID string) (*maps.PlaceDetailsResult, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.PlaceDetailsRequest{
		PlaceID: placeID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.client.PlaceDetails(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *MapsApiClient) doReverseGeocode(lat, lng float64) ([]maps.GeocodingResult, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.client.Geocode(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, err
}
