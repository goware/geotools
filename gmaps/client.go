package gmaps

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) doAutocomplete(ctx context.Context, input string) ([]maps.QueryAutocompletePrediction, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.QueryAutocompleteRequest{
		Input: input,
	}

	res, err := c.client.QueryAutocomplete(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res.Predictions, err
}

func (c *MapsApiClient) doTextSearch(ctx context.Context, input string) ([]maps.PlacesSearchResult, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.TextSearchRequest{
		Query: input,
	}

	res, err := c.client.TextSearch(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res.Results, err
}

func (c *MapsApiClient) doDetails(ctx context.Context, placeID string) (*maps.PlaceDetailsResult, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.PlaceDetailsRequest{
		PlaceID: placeID,
	}

	res, err := c.client.PlaceDetails(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *MapsApiClient) doReverseGeocode(ctx context.Context, lat, lng float64) ([]maps.GeocodingResult, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
	}

	res, err := c.client.Geocode(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, err
}
