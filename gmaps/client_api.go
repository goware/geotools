// +build !mock

package gmaps

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

func (c *MapsApiClient) Autocomplete(ctx context.Context, input string) ([]maps.QueryAutocompletePrediction, error) {
	return c.doAutocomplete(ctx, input)
}

func (c *MapsApiClient) TextSearch(ctx context.Context, input string) ([]maps.PlacesSearchResult, error) {
	return c.doTextSearch(ctx, input)
}

func (c *MapsApiClient) Details(ctx context.Context, placeID string) (*maps.PlaceDetailsResult, error) {
	return c.doDetails(ctx, placeID)
}

func (c *MapsApiClient) ReverseGeocode(ctx context.Context, lat, lng float64) ([]maps.GeocodingResult, error) {
	return c.doReverseGeocode(ctx, lat, lng)
}
