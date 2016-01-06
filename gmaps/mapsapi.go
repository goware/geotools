package gmaps

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"time"
)

var (
	defaultTimeout = time.Second * 15
)

type MapsApiClient struct {
	key    string
	client *maps.Client
}

func NewMapsClient(key string) (*MapsApiClient, error) {
	c := &MapsApiClient{key: key}
	mapsClient, err := maps.NewClient(maps.WithAPIKey(c.key))
	if err != nil {
		return nil, err
	}
	c.client = mapsClient
	return c, nil
}

func (c *MapsApiClient) Autocomplete(input string) ([]maps.QueryAutocompletePrediction, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.QueryAutocompleteRequest{
		Input: input,
		// Missing lat,lng, missing radius.
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	res, err := c.client.QueryAutocomplete(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res.Predictions, err
}

func (c *MapsApiClient) Details(placeID string) (*maps.PlaceDetailsResult, error) {
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

func (c *MapsApiClient) ReverseGeocode(lat, lng float64) ([]maps.GeocodingResult, error) {
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
