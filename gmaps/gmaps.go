package gmaps

import (
	"googlemaps.github.io/maps"
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