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

	lctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	res, err := c.client.QueryAutocomplete(lctx, &req)
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

	lctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	res, err := c.client.TextSearch(lctx, &req)
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

	lctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	res, err := c.client.PlaceDetails(lctx, &req)
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

	lctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	res, err := c.client.Geocode(lctx, &req)
	if err != nil {
		return nil, err
	}

	return res, err
}
