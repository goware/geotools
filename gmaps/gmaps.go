package gmaps

import (
	"time"

	"googlemaps.github.io/maps"
)

var queryTimeout = time.Second * 10

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
