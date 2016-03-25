package geotools

import (
	"github.com/goware/geotools/gmaps"
	"golang.org/x/net/context"
)

// MapsClient defines methods for geoy client.
type MapsClient interface {
	PlaceDetails(context.Context, string) (*Place, error)
	LookupCoordinates(context.Context, LatLnger) ([]*Place, error)
	LookupName(context.Context, string) ([]*Place, error)
}

// Client represents a geoy maps client.
type Client struct {
	m gmaps.MapsApiClient
}

var (
	defaultMapsClient MapsClient
)

// SetAPIKey sets the Google Maps API key.
func SetAPIKey(key string) error {
	c, err := NewClient(key)
	if err != nil {
		return err
	}
	SetDefaultClient(c)
	return nil
}

// SetDefaultClient sets the default client (package-wide).
func SetDefaultClient(c MapsClient) {
	defaultMapsClient = c
}

// NewClient creates a maps client given a Google Maps API key.
func NewClient(key string) (MapsClient, error) {
	c, err := gmaps.NewMapsClient(key)
	if err != nil {
		return nil, err
	}
	return NewClientFromMap(c), nil
}

// NewClientFromMap creates a client with the given gmaps.MapsApiClient.
func NewClientFromMap(c gmaps.MapsApiClient) MapsClient {
	return &Client{c}
}

func (c *Client) mapsAPI() gmaps.MapsApiClient {
	if c == nil || c.m == nil {
		panic("Maps client was not initialized. Missing call to SetAPIKey()?")
	}
	return c.m
}

// PlaceDetails returns the details of a place given its placeID.
func (c *Client) PlaceDetails(ctx context.Context, placeID string) (*Place, error) {
	place, err := c.mapsAPI().Details(ctx, placeID)
	if err != nil {
		return nil, err
	}
	res := result{
		PlaceID:           place.PlaceID,
		AddressComponents: place.AddressComponents,
		Geometry:          place.Geometry,
		FormattedAddress:  place.FormattedAddress,
	}
	return toPlace(&res), nil
}

// LookupCoordinates lookups a coordinate and returns all its associated
// places.
func (c *Client) LookupCoordinates(ctx context.Context, p LatLnger) ([]*Place, error) {
	l := p.LatLng()
	places, err := c.mapsAPI().ReverseGeocode(ctx, l[0], l[1])
	if err != nil {
		return nil, err
	}

	res := make([]*Place, len(places))
	for i := range places {
		res[i] = toPlace(&result{
			PlaceID:           places[i].PlaceID,
			AddressComponents: places[i].AddressComponents,
			Geometry:          places[i].Geometry,
			FormattedAddress:  places[i].FormattedAddress,
		})
	}

	return res, nil
}

// LookupName converts a string place name/address into
func (c *Client) LookupName(ctx context.Context, s string) ([]*Place, error) {
	places, err := c.mapsAPI().TextSearch(ctx, s)
	if err != nil {
		return nil, err
	}
	res := make([]*Place, len(places))
	for i := range places {
		res[i] = toPlace(&result{
			PlaceID: places[i].PlaceID,
			//AddressComponents: places[i].AddressComponents,
			Geometry:         places[i].Geometry,
			FormattedAddress: places[i].FormattedAddress,
		})
	}
	return res, nil
}

// PlaceDetails returns PlaceDetails using the default client.
func PlaceDetails(ctx context.Context, s string) (*Place, error) {
	return defaultMapsClient.PlaceDetails(ctx, s)
}

// LookupCoordinates returns LookupCoordinates using the default client.
func LookupCoordinates(ctx context.Context, ll LatLnger) ([]*Place, error) {
	return defaultMapsClient.LookupCoordinates(ctx, ll)
}

// LookupName returns LookupName using the default client.
func LookupName(ctx context.Context, s string) ([]*Place, error) {
	return defaultMapsClient.LookupName(ctx, s)
}
