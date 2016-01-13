package gmaps

import (
	"time"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var queryTimeout = time.Second * 10

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
