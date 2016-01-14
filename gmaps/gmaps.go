package gmaps

import (
	"errors"
	"time"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var queryTimeout = time.Second * 10

var (
	errMissingClient = errors.New("Missing client")
)

type MapsApiClient interface {
	Autocomplete(context.Context, string) ([]maps.QueryAutocompletePrediction, error)
	TextSearch(context.Context, string) ([]maps.PlacesSearchResult, error)
	Details(context.Context, string) (*maps.PlaceDetailsResult, error)
	ReverseGeocode(context.Context, float64, float64) ([]maps.GeocodingResult, error)
}

type mapsApiClient struct {
	key    string
	client *maps.Client
}

func NewMapsClient(key string) (MapsApiClient, error) {
	c := &mapsApiClient{key: key}
	mapsClient, err := maps.NewClient(maps.WithAPIKey(c.key))
	if err != nil {
		return nil, err
	}
	c.client = mapsClient
	return c, nil
}

func (c *mapsApiClient) Autocomplete(ctx context.Context, input string) ([]maps.QueryAutocompletePrediction, error) {
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

func (c *mapsApiClient) TextSearch(ctx context.Context, input string) ([]maps.PlacesSearchResult, error) {
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

func (c *mapsApiClient) Details(ctx context.Context, placeID string) (*maps.PlaceDetailsResult, error) {
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

func (c *mapsApiClient) ReverseGeocode(ctx context.Context, lat, lng float64) ([]maps.GeocodingResult, error) {
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
