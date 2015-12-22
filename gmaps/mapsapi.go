package gmaps

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"time"
)

var (
	defaultTimeout = time.Second * 15
)

/*
type autocompleteResponse struct {
	Status      string
	Predictions []prediction
}

type prediction struct {
	Description string
	PlaceId     string `json:"place_id"`
}

type geocodeResponse struct {
	Status  string
	Results []maps.PlaceDetailsResult
}

type geometry struct {
	Location latlng
	Viewport *viewport
}

type viewport struct {
	Northeast latlng
	Southwest latlng
}

type latlng struct {
	Lat float64
	Lng float64
}

func (l latlng) LatLng() []float64 {
	return []float64{l.Lat, l.Lng}
}

func (l *latlng) String() string {
	return fmt.Sprintf("%0.6f,%0.6f", l.Lat, l.Lng)
}
*/

type MapsApiClient struct {
	key    string
	client *maps.Client
}

func NewMapsClient(key string) *MapsApiClient {
	c := &MapsApiClient{key: key}
	c.client, _ = maps.NewClient(maps.WithAPIKey(c.key))
	return c
}

func (c *MapsApiClient) Autocomplete(input string) ([]maps.QueryAutocompletePrediction, error) {
	if c.client == nil {
		return nil, errMissingClient
	}

	req := maps.QueryAutocompleteRequest{
		Input: input,
		// Missing lat,lng, missing radius.
	}

	ctx, fn := context.WithTimeout(context.Background(), defaultTimeout)
	defer fn()

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

	ctx, fn := context.WithTimeout(context.Background(), defaultTimeout)
	defer fn()

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

	ctx, fn := context.WithTimeout(context.Background(), defaultTimeout)
	defer fn()

	res, err := c.client.Geocode(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, err
}
